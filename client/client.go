package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	ofcmessages "github.com/thomas-osgood/ofs/client/internal/messages"
	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	protocommon "github.com/thomas-osgood/ofs/protobufs/common"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	"github.com/thomas-osgood/ofs/protobufs/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ease-of-use function that creates the gRPC connection and a
// filehandler client to communicate with the server.
func (fc *FClient) initConnection() (conn *grpc.ClientConn, client filehandler.FileserviceClient, err error) {

	conn, err = grpc.Dial(fc.srvaddr, fc.srvopts...)
	if err != nil {
		return nil, nil, err
	}

	client = filehandler.NewFileserviceClient(conn)

	return conn, client, nil
}

// function designed to request a file be deleted from the server's
// uploads directory.
func (fc *FClient) DeleteFile(targetfile string) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var resp *protocommon.StatusMessage

	conn, client, err = fc.initConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	resp, err = client.DeleteFile(ctx, &filehandler.FileRequest{Filename: targetfile})
	if err != nil {
		return err
	} else if resp.GetCode() >= http.StatusBadRequest {
		return fmt.Errorf(resp.GetMessage())
	}

	return nil
}

// function designed to download the file contents from the file server.
//
// this will save the content to the filename specified in the req object.
func (fc *FClient) DownloadFile(req *filehandler.FileRequest) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var fptr *os.File
	var uploader filehandler.Fileservice_UploadFileClient

	// keep track of the number of active downloads.
	fc.increaseActiveDownloads()
	defer fc.decreaseActiveDownloads()

	conn, client, err = fc.initConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	// get a file pointer pointing to the destination. this will
	// open the destination file in WRITE mode and TRUNC mode, clearing
	// out any data if the file already exists.
	fptr, err = os.OpenFile(req.GetFilename(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		log.Printf(ofcmessages.ERR_OPEN_FILE, err.Error())
		return err
	}
	defer fptr.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	// get object used to download data from the server.
	uploader, err = client.UploadFile(ctx, req)
	if err != nil {
		log.Printf(ofcmessages.ERR_UPLOAD_FILE, err.Error())
		return err
	}

	// update the timeout value for the context to the transfer
	// timeout. by this point, the connection to the server has
	// already been established.
	ctx, cancel = context.WithTimeout(context.Background(), fc.transferCfg.Timeout)
	defer cancel()

	// this select statement acts as a timeout mechanism for the
	// receive file bytes function. if the context times out before
	// all bytes are received, an error will be returned and the
	// partial download will be deleted.
	select {
	case <-ctx.Done():
		// close the file pointer and remove the empty file, then
		// return an error stating there was a timeout while attempting
		// to transfer the file to the client.
		fptr.Close()
		os.Remove(req.GetFilename())
		return fmt.Errorf(ofcmessages.ERR_TRANSMIT_TIMEOUT)
	default:
		// read content streamed down from the server and save it
		// in the file specified in the request message.
		err = ofscommon.ReceiveFileBytes(uploader, fptr)
		if err != nil {
			// close the file pointer and remove the empty file,
			// then return the error that was thrown during the
			// transfer process.
			fptr.Close()
			os.Remove(req.GetFilename())
			return err
		}
	}

	// close the stream.
	err = uploader.CloseSend()
	if err != nil {
		return err
	}

	return nil
}

// function designed to get the list of files the client
// is able to download from the server.
func (fc *FClient) ListFiles() (files []*filehandler.FileInfo, err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var curfile *filehandler.FileInfo
	var lister filehandler.Fileservice_ListFilesClient

	conn, client, err = fc.initConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	lister, err = client.ListFiles(ctx, &protocommon.Empty{})
	if err != nil {
		return nil, err
	}

	files = make([]*filehandler.FileInfo, 0)

	// read all fileinfo objects the server streams to the client.
	for {
		curfile, err = lister.Recv()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		files = append(
			files,
			&filehandler.FileInfo{
				Name:         curfile.GetName(),
				Sizebytes:    curfile.GetSizebytes(),
				Isdir:        curfile.GetIsdir(),
				Lastmodified: curfile.GetLastmodified(),
			},
		)
	}

	err = lister.CloseSend()
	if err != nil {
		return nil, err
	}

	return files, nil
}

// function designed to make a specified directory within
// the server's uploads directory.
func (fc *FClient) MakeDirectory(dirname string) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var resp *protocommon.StatusMessage

	conn, client, err = fc.initConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	resp, err = client.MakeDirectory(ctx, &filehandler.MakeDirectoryRequest{Dirname: dirname})
	if err != nil {
		return err
	} else if resp.GetCode() >= http.StatusBadRequest {
		return fmt.Errorf(resp.GetMessage())
	}

	return nil
}

// function designed to download multiple files at one time.
//
// this will return a map containing the filenames and an associated
// error. if no errors occurred, the map will be empty.
func (fc *FClient) MultifileDownload(targets []string) (errs map[string]error) {
	var mut sync.Mutex
	var target string
	var wg sync.WaitGroup

	// initialize the return map to avoid nil reference errors.
	errs = make(map[string]error)

	// loop through each specified file and attempt a download.
	// if successful, no entry will be added to the errs map.
	// if an error occurs, the error will be attached to the
	// filename via the errs map.
	for _, target = range targets {
		wg.Add(1)
		go fc.mfdWorker(target, &errs, &wg, &mut)
	}

	// wait for all go routines to complete before returning.
	wg.Wait()

	return errs
}

// function designed to upload multiple files at one time.
//
// this will return a map containing the filenames and an associated
// error. if no errors occurred, the map will be empty.
func (fc *FClient) MultifileUpload(targets []string) (errs map[string]error) {
	var mut sync.Mutex
	var target string
	var wg sync.WaitGroup

	// initialize the return map to avoid nil reference errors.
	errs = make(map[string]error)

	// loop through each file and attempt to uploda it to the
	// server. if an error occurs add a new entry to the return
	// map connecting the filename and error.
	for _, target = range targets {
		wg.Add(1)
		go fc.mfuWorker(target, &errs, &wg, &mut)
	}

	// wait for all go routines to complete before returning.
	wg.Wait()

	return errs
}

// function designed to check whether the server is up and able
// to be contacted.
func (fc *FClient) Ping() (roundtrip time.Duration, err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var pong *pingpong.Pong
	var reqtime time.Time
	var restime time.Time

	conn, client, err = fc.initConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	pong, err = client.Ping(ctx, &pingpong.Ping{Stamp: timestamppb.Now()})
	if err != nil {
		return 0, err
	}

	reqtime = pong.GetReqtime().AsTime()
	restime = pong.GetResptime().AsTime()

	roundtrip = restime.Sub(reqtime)

	return roundtrip, nil
}

// function designed to request a file in the server's uploads
// directory be renamed.
func (fc *FClient) RenameFile(originalname string, newname string) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var resp *protocommon.StatusMessage

	conn, client, err = fc.initConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	// make sure the filenames passed in are not empty strings.
	originalname = strings.TrimSpace(originalname)
	newname = strings.TrimSpace(newname)
	if (len(originalname) < 1) || (len(newname) < 1) {
		return fmt.Errorf(ofcmessages.ERR_FILENAME_EMPTY)
	}

	// request the server rename the file.
	resp, err = client.RenameFile(
		ctx,
		&filehandler.RenameFileRequest{
			Oldfilename: originalname,
			Newfilename: newname,
		},
	)

	if err != nil {
		return err
	} else if resp.GetCode() >= http.StatusBadRequest {
		return fmt.Errorf(resp.GetMessage())
	}

	return nil
}

// function designed to call the StorageBreakdown rpc and get
// the storage consumption information from the server.
func (fc *FClient) StorageBreakdown() (consumption *filehandler.StorageInfo, err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context

	conn, client, err = fc.initConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	consumption, err = client.StorageBreakdown(ctx, &protocommon.Empty{})
	if err != nil {
		return nil, status.Convert(err).Err()
	}

	return consumption, nil
}

// function designed to upload a file to the file server.
func (fc *FClient) UploadFile(filename string) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var fptr *os.File
	var srv filehandler.Fileservice_DownloadFileClient
	var resp *protocommon.StatusMessage

	// keep track of the number of active uploads.
	fc.increaseActiveUploads()
	defer fc.decreaseActiveUploads()

	conn, client, err = fc.initConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	fptr, err = os.Open(filename)
	if err != nil {
		return err
	}
	defer fptr.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	// add the filename to the header information.
	ctx, err = addFilenameMD(ctx, filename)
	if err != nil {
		return err
	}

	// get object used to stream data up to the server.
	srv, err = client.DownloadFile(ctx)
	if err != nil {
		return err
	}

	// stream the contents of the target file up to the server.
	err = ofscommon.TransmitFileBytes(srv, bufio.NewReader(fptr))
	if err != nil {
		return err
	}

	// close the stream and get the server's status response message.
	resp, err = srv.CloseAndRecv()
	if err != nil {
		return err
	} else if resp.GetCode() != http.StatusOK {
		return fmt.Errorf(ofcmessages.ERR_TRANSMIT_FILE, resp.GetMessage())
	}

	return nil
}
