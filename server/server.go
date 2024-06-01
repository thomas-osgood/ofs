package server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	"github.com/thomas-osgood/ofs/protobufs/common"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	ofsdefaults "github.com/thomas-osgood/ofs/server/internal/defaults"
	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
	ofsutils "github.com/thomas-osgood/ofs/server/internal/utils"
	"google.golang.org/grpc/status"
)

// function designed to download a file from a client to the server.
//
// this will save the file uploaded by the client to the root directory
// or downloads directory (if one has been set).
func (fsrv *FServer) DownloadFile(srv filehandler.Fileservice_DownloadFileServer) (err error) {
	var filename string
	var tmpname string

	fsrv.debugMessage(ofsmessages.DBG_IN_DOWNLOAD)

	filename, err = ofsutils.ReadFilenameMD(srv.Context())
	if err != nil {
		fsrv.debugMessageErr(fmt.Sprintf(ofsmessages.ERR_MD, err.Error()))
		return err
	}
	filename = fsrv.cleanFilename(filename, ofsdefaults.FTYPE_DOWNLOAD)

	fsrv.debugMessageSuc(fmt.Sprintf(ofsmessages.DBG_FILENAME, filename))

	// read the data stream and save it to a temporary file.
	tmpname, err = fsrv.readIncomingFile(srv)
	if err != nil {
		return err
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
	var subdir string = fsrv.buildUploadFilename(filepath.Clean(dirreq.GetDirname()))

	// initialize the successful StatusMessage. if everything
	// works as expected, this will not be modified.
	retstatus = &common.StatusMessage{
		Code:    http.StatusCreated,
		Message: "directory created",
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

// function designed to rename a file in the uploads directory. this
// will move the source file to the destination.
func (fsrv *FServer) RenameFile(ctx context.Context, rnreq *filehandler.RenameFileRequest) (resp *common.StatusMessage, err error) {
	var absdest string = fsrv.buildUploadFilename(rnreq.GetNewfilename())
	var abssrc string = fsrv.buildUploadFilename(rnreq.GetOldfilename())

	// check for the existence of the destination file. if the
	// destination file already exists, an error will be returned
	// saying as much.
	if err = fsrv.fileExists(absdest); err != nil {
		if errors.Is(err, os.ErrPermission) {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf(ofsmessages.ERR_FILE_EXISTS)
	}

	// move the source file to the destination.
	err = fsrv.moveTempfile(abssrc, absdest)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// function designed to upload a requested file from the server to the client.
//
// the specified file must exist in the root directory (or the uploads directory
// if one is specified) for a successful (nil error) return.
func (fsrv *FServer) UploadFile(req *filehandler.FileRequest, srv filehandler.Fileservice_UploadFileServer) (err error) {
	var fptr *os.File
	var targetfile string = fsrv.cleanFilename(req.GetFilename(), ofsdefaults.FTYPE_UPLOAD)

	fsrv.debugMessage(fmt.Sprintf(ofsmessages.DBG_FILE_REQUEST, targetfile))

	fptr, err = os.Open(targetfile)
	if err != nil {
		return err
	}
	defer fptr.Close()

	err = ofscommon.TransmitFileBytes(srv, bufio.NewReader(fptr))
	if err != nil {
		return err
	}

	fsrv.debugMessageSuc("file successfully transmitted")

	return nil
}
