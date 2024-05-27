package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/thomas-osgood/OGOR/protobufs/definitions/common"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	"github.com/thomas-osgood/OGOR/protobufs/general"
	"github.com/thomas-osgood/ofs/client/internal/messages"
	"google.golang.org/grpc"
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
		log.Printf(messages.ERR_OPEN_FILE, err.Error())
		return err
	}
	defer fptr.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	// get object used to download data from the server.
	uploader, err = client.UploadFile(ctx, req)
	if err != nil {
		log.Printf(messages.ERR_UPLOAD_FILE, err.Error())
		return err
	}

	// read content streamed down from the server and save it
	// in the file specified in the request message.
	err = general.ReceiveFileBytes(uploader, fptr)
	if err != nil {
		// close the file pointer and remove the empty file,
		// then return the error that was thrown during the
		// transfer process.
		fptr.Close()
		os.Remove(req.GetFilename())
		return err
	}

	// close the stream.
	err = uploader.CloseSend()
	if err != nil {
		return err
	}

	return nil
}

// function designed to upload a file to the file server.
func (fc *FClient) UploadFile(filename string) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var fptr *os.File
	var srv filehandler.Fileservice_DownloadFileClient
	var status *common.StatusMessage

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
	err = general.TransmitFileBytes(srv, bufio.NewReader(fptr))
	if err != nil {
		return err
	}

	// close the stream and get the server's status response message.
	status, err = srv.CloseAndRecv()
	if err != nil {
		return err
	} else if status.GetCode() != http.StatusOK {
		return fmt.Errorf(messages.ERR_TRANSMIT_FILE, status.GetMessage())
	}

	return nil
}
