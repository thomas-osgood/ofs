package server

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	ofscommon "github.com/thomas-osgood/ofs/general"
	"github.com/thomas-osgood/ofs/protobufs/common"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
	ofsutils "github.com/thomas-osgood/ofs/server/internal/utils"
)

// function designed to build the directory structure the fileserver
// will use to save and server files.
func (fs *FServer) buildDirStructure() (err error) {
	var downloaddir string = filepath.Join(fs.rootdir, fs.downloadsdir)
	var uploaddir string = filepath.Join(fs.rootdir, fs.uploadsdir)

	err = os.MkdirAll(downloaddir, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(uploaddir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// function designed to clean an uploaded filename and return
// only the filename portion of it. this will strip the directory
// information.
func (fs *FServer) cleanFilename(filename string, ftype string) (cleaned string) {
	var fnsplit []string
	var subdir string

	switch strings.ToLower(ftype) {
	case FTYPE_DOWNLOAD:
		subdir = fs.downloadsdir
	case FTYPE_UPLOAD:
		subdir = fs.uploadsdir
	default:
		subdir = ""
	}

	filename = filepath.Clean(filename)
	fnsplit = strings.Split(filename, fmt.Sprintf("%c", os.PathSeparator))
	cleaned = filepath.Join(fs.rootdir, subdir, fnsplit[len(fnsplit)-1])

	return cleaned
}

// helper function for outputting a debug message to STDOUT. this
// will only print output if the debug flag is set.
func (fs *FServer) debugMessage(message string) {
	if fs.debug {
		fs.printer.InfMsg(message)
	}
}

// function designed to download a file from a client to the server.
func (fs *FServer) DownloadFile(srv filehandler.Fileservice_DownloadFileServer) (err error) {
	var filename string
	var tmpfile *os.File
	var tmpname string

	fs.debugMessage(ofsmessages.DBG_IN_DOWNLOAD)

	filename, err = ofsutils.ReadFilenameMD(srv.Context())
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_MD, err.Error()))
		return err
	}
	filename = fs.cleanFilename(filename, FTYPE_DOWNLOAD)

	if fs.debug {
		fs.printer.SucMsg(fmt.Sprintf(ofsmessages.DBG_FILENAME, filename))
	}

	// create a temporary file to hold the uploaded information.
	tmpfile, err = os.CreateTemp("", "download")
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_TEMP, err.Error()))
		return err
	}
	defer tmpfile.Close()

	tmpname = tmpfile.Name()

	if fs.debug {
		fs.printer.SysMsgNB(fmt.Sprintf(ofsmessages.UPLOAD_IN_PROGRESS, tmpname))
	}

	// stream the file contents from the client and write them
	// to the temp file.
	err = ofscommon.ReceiveFileBytes(srv, tmpfile)
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_RECV, err.Error()))
		return err
	}
	err = srv.SendAndClose(&common.StatusMessage{Message: ofsmessages.UPLOAD_COMPLETE, Code: http.StatusOK})
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_ACK, err.Error()))
	}

	if fs.debug {
		fs.printer.SucMsg(ofsmessages.UPLOAD_COMPLETE)
	}

	// close and re-open the temp file in read mode. if this
	// is not done, no bytes will be copied from the temp file
	// to the destination file.
	tmpfile.Close()
	tmpfile, err = os.Open(tmpname)
	if err != nil {
		return err
	}
	defer tmpfile.Close()

	if fs.debug {
		fs.printer.SysMsgNB(ofsmessages.COPY_IN_PROGRESS)
	}

	err = ofscommon.CopyFile(tmpfile, filename)
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_COPY_FILE, err.Error()))
		return err
	}

	if fs.debug {
		fs.printer.SucMsg(ofsmessages.COPY_COMPLETE)
	}

	tmpfile.Close()

	// delete the temporary file used during the upload process.
	err = os.Remove(tmpname)
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_REMOVE_TEMP, err.Error()))
	}

	if fs.debug {
		fs.printer.SucMsg(ofsmessages.TEMP_REMOVED)
	}

	return nil
}

// function designed to upload a requested file from the server to the client.
func (fs *FServer) UploadFile(req *filehandler.FileRequest, srv filehandler.Fileservice_UploadFileServer) (err error) {
	var fptr *os.File
	var targetfile string = fs.cleanFilename(req.GetFilename(), FTYPE_UPLOAD)

	fs.debugMessage(fmt.Sprintf(ofsmessages.DBG_FILE_REQUEST, targetfile))

	fptr, err = os.Open(targetfile)
	if err != nil {
		return err
	}
	defer fptr.Close()

	err = ofscommon.TransmitFileBytes(srv, bufio.NewReader(fptr))
	if err != nil {
		return err
	}

	return nil
}
