package server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"time"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	"github.com/thomas-osgood/ofs/protobufs/common"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	"github.com/thomas-osgood/ofs/protobufs/pingpong"
	ofsdefaults "github.com/thomas-osgood/ofs/server/internal/defaults"
	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
	ofsutils "github.com/thomas-osgood/ofs/server/internal/utils"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// function designed to delete a file in the uploads directory
// as requested by the client.
func (fsrv *FServer) DeleteFile(ctx context.Context, req *filehandler.FileRequest) (resp *common.StatusMessage, err error) {
	var targetpath string = strings.TrimSpace(req.GetFilename())

	resp = &common.StatusMessage{
		Code:    http.StatusOK,
		Message: ofsmessages.FILE_DELETED,
	}

	if len(targetpath) < 1 {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		return resp, nil
	}

	targetpath = fsrv.buildUploadFilename(targetpath)

	fsrv.debugMessage(fmt.Sprintf(ofsmessages.DBG_DELETING_FILE, targetpath))
	err = os.Remove(targetpath)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
	}
	fsrv.debugMessageSuc(ofsmessages.DBG_DELETE_SUCCESS)

	return resp, nil
}

// function designed to download a file from a client to the server.
//
// this will save the file uploaded by the client to the root directory
// or downloads directory (if one has been set).
func (fsrv *FServer) DownloadFile(srv filehandler.Fileservice_DownloadFileServer) (err error) {
	var cancel context.CancelFunc
	var ctx context.Context
	var filename string
	var tmpname string

	fsrv.increaseActiveDownloads()
	defer fsrv.decreaseActiveDownloads()

	fsrv.debugMessage(ofsmessages.DBG_IN_DOWNLOAD)

	filename, err = ofsutils.ReadFilenameMD(srv.Context())
	if err != nil {
		fsrv.debugMessageErr(fmt.Sprintf(ofsmessages.ERR_MD, err.Error()))
		return err
	}
	filename = fsrv.cleanFilename(filename, ofsdefaults.FTYPE_DOWNLOAD)

	fsrv.debugMessage(fmt.Sprintf(ofsmessages.DBG_FILENAME, filename))

	ctx, cancel = context.WithTimeout(context.Background(), fsrv.transferCfg.TransferTimeout)
	defer cancel()

	// attempt to download the file being uploaded by the client.
	// if the timeout takes too long, the function will return an
	// error.
	select {
	case <-ctx.Done():
		return fmt.Errorf(ofsmessages.ERR_TRANSFER_TIMEOUT)
	default:
		// read the data stream and save it to a temporary file.
		tmpname, err = fsrv.readIncomingFile(srv)
		if err != nil {
			return err
		}
	}

	// move over the tmpfile contents to the destination file.
	err = fsrv.moveTempfile(tmpname, filename)
	if err != nil {
		return err
	}

	fsrv.debugMessageSuc(ofsmessages.TEMP_REMOVED)

	return nil
}

// function designed to list out the files in the directory the client has
// the ability to download files from. if there are separate upload and
// download directories, only the "uploads" (files that can be uploaded
// from the server to client) directory will be listed.
func (fsrv *FServer) ListFiles(mpty *common.Empty, srv filehandler.Fileservice_ListFilesServer) (err error) {
	var curfile *filehandler.FileInfo
	var files []*filehandler.FileInfo

	// gather all files in the uploads directory.
	files, err = fsrv.listUploadsDir()
	if err != nil {
		return err
	}

	// transmit the files to the client.
	for _, curfile = range files {
		err = srv.Send(curfile)
		if err != nil {
			return err
		}
	}

	return nil
}

// function designed to create a new sub-directory within the uploads directory.
// if the sub-directory has sub-directories, this function will attempt to
// create all directories.
//
// successful code: 201 Created
//
// failure code: 500 Internal Server Error
func (fsrv *FServer) MakeDirectory(ctx context.Context, dirreq *filehandler.MakeDirectoryRequest) (retstatus *common.StatusMessage, err error) {
	var subdir string = fsrv.buildUploadFilename(dirreq.GetDirname())

	// initialize the successful StatusMessage. if everything
	// works as expected, this will not be modified.
	retstatus = &common.StatusMessage{
		Code:    http.StatusCreated,
		Message: ofsmessages.DIR_CREATED,
	}

	// attempt to create sub-directory specified by the client.
	// if this fails, the StatusMessage returned to the client
	// will indicate what the error was.
	err = os.MkdirAll(subdir, os.ModePerm)
	if err != nil {

		retstatus.Code = http.StatusInternalServerError
		retstatus.Message = status.Convert(err).Message()

	}

	return retstatus, nil
}

// function designed to let the client know that the server is up
// and able to be contacted.
func (fsrv *FServer) Ping(ctx context.Context, ping *pingpong.Ping) (pong *pingpong.Pong, err error) {
	fsrv.debugMessage(fmt.Sprintf("ping received: %s", ping.GetStamp().AsTime().Format(time.RFC3339)))
	return &pingpong.Pong{Reqtime: ping.GetStamp(), Resptime: timestamppb.Now()}, nil
}

// function designed to rename a file in the uploads directory. this
// will move the source file to the destination.
//
// if the file does not already exist, the success or failure status
// will be transmitted to the client using a StatusMessage object. the
// code should be used to determine whether the copy was successful or not.
//
// success status code: 200 OK
//
// failure status code: 500 Internal Server Error
//
// note: if the destination file already exists, or if there is an issue
// related to permissions, an error will be returned.
func (fsrv *FServer) RenameFile(ctx context.Context, rnreq *filehandler.RenameFileRequest) (resp *common.StatusMessage, err error) {
	var absdest string = fsrv.buildUploadFilename(rnreq.GetNewfilename())
	var abssrc string = fsrv.buildUploadFilename(rnreq.GetOldfilename())

	resp = new(common.StatusMessage)

	// check for the existence of the destination file. if the
	// destination file already exists, an error will be returned
	// saying as much.
	if err = fsrv.fileExists(absdest); err == nil {
		return nil, fmt.Errorf(ofsmessages.ERR_FILE_EXISTS)
	} else if errors.Is(err, os.ErrPermission) {
		return nil, err
	}

	// move the source file to the destination.
	fsrv.debugMessage(fmt.Sprintf(ofsmessages.DBG_RENAME_START, abssrc, absdest))
	err = fsrv.moveTempfile(abssrc, absdest)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = status.Convert(err).Message()
	} else {
		resp.Code = http.StatusOK
		resp.Message = ofsmessages.COPY_COMPLETE
	}

	return resp, nil
}

// function designed to calculate and return the amount of storage (in bytes)
// consumed by the uploads and downloads directories.
func (fsrv *FServer) StorageBreakdown(ctx context.Context, mpty *common.Empty) (consumption *filehandler.StorageInfo, err error) {
	consumption = new(filehandler.StorageInfo)

	consumption.Downloads, err = fsrv.calculateDirectoryConsumption(fsrv.downloadsdir)
	if err != nil {
		return nil, err
	}

	consumption.Uploads, err = fsrv.calculateDirectoryConsumption(fsrv.uploadsdir)
	if err != nil {
		return nil, err
	}

	consumption.Total = consumption.Downloads + consumption.Uploads

	return consumption, nil
}

// function designed to upload a requested file from the server to the client.
//
// the specified file must exist in the root directory (or the uploads directory
// if one is specified) for a successful (nil error) return.
func (fsrv *FServer) UploadFile(req *filehandler.FileRequest, srv filehandler.Fileservice_UploadFileServer) (err error) {
	var cancel context.CancelFunc
	var ctx context.Context
	var finfo fs.FileInfo
	var fptr *os.File
	var md metadata.MD = make(metadata.MD)
	var targetfile string = fsrv.cleanFilename(req.GetFilename(), ofsdefaults.FTYPE_UPLOAD)

	fsrv.increaseActiveUploads()
	defer fsrv.decreaseActiveUploads()

	fsrv.debugMessage(fmt.Sprintf(ofsmessages.DBG_FILE_REQUEST, targetfile))

	// get the information for the requested file.
	finfo, err = os.Stat(targetfile)
	if err != nil {
		return err
	}

	// add the file's size to the metadata.
	md.Set(ofscommon.HEADER_FILESIZE, fmt.Sprintf("%d", finfo.Size()))
	// add file's last modified timestamp to metadata.
	md.Set(ofscommon.HEADER_LASTMODIFIED, finfo.ModTime().Format(time.RFC3339))

	// set the metadata that will be transmitted when the server communicates
	// with the client.
	srv.SetHeader(md)

	fptr, err = os.Open(targetfile)
	if err != nil {
		return err
	}
	defer fptr.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fsrv.transferCfg.TransferTimeout)
	defer cancel()

	// attempt to upload the file to the client. if the transfer takes too long,
	// the function will return an error.
	select {
	case <-ctx.Done():
		return fmt.Errorf(ofsmessages.ERR_TRANSFER_TIMEOUT)
	default:
		err = ofscommon.TransmitFileBytes(srv, bufio.NewReader(fptr))
		if err != nil {
			return err
		}
	}

	fsrv.debugMessageSuc(ofsmessages.FILE_SENT)

	return nil
}
