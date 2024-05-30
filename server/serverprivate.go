package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	protocommon "github.com/thomas-osgood/ofs/protobufs/common"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	ofsdefaults "github.com/thomas-osgood/ofs/server/internal/defaults"
	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
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
// information and give it an absolute path within the directory
// associated with the type of file specified (download, upload, etc).
func (fs *FServer) cleanFilename(filename string, ftype string) (cleaned string) {
	var fnsplit []string
	var subdir string

	switch strings.ToLower(ftype) {
	case ofsdefaults.FTYPE_DOWNLOAD:
		subdir = fs.downloadsdir
	case ofsdefaults.FTYPE_UPLOAD:
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

// helper function for outputting an error message to STDOUT. this
// will only print output if the debug flag is set.
func (fs *FServer) debugMessageErr(message string) {
	if fs.debug {
		fs.printer.ErrMsg(message)
	}
}

// helper function for outputting a success message to STDOUT. this
// will only print output if the debug flag is set.
func (fs *FServer) debugMessageSuc(message string) {
	if fs.debug {
		fs.printer.SucMsg(message)
	}
}

// function designed to move the contents of a temporary file
// to a specified destination.
//
// this will open the temporary file and close it when the
// function returns.
func (fs *FServer) moveTempfile(tmpname string, filename string) (err error) {

	fs.debugMessageSuc(ofsmessages.COPY_IN_PROGRESS)

	err = ofscommon.MoveFile(tmpname, filename)
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_COPY_FILE, err.Error()))
		return err
	}

	fs.debugMessageSuc(ofsmessages.COPY_COMPLETE)

	return nil
}

// function designed to read the file data from the incoming
// file byte stream and save it to a temporary file.
//
// this will create a temporary file containing the data that
// gets uploaded, close the file upon return and return the
// temporary file name so the file can be used later on.
func (fs *FServer) readIncomingFile(srv filehandler.Fileservice_DownloadFileServer) (tmpname string, err error) {
	var tmpfile *os.File

	// create a temporary file to hold the uploaded information.
	tmpfile, err = os.CreateTemp("", "download")
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_TEMP, err.Error()))
		return "", err
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
		return "", err
	}

	// send a successful status message and close the stream.
	err = srv.SendAndClose(&protocommon.StatusMessage{Message: ofsmessages.UPLOAD_COMPLETE, Code: http.StatusOK})
	if err != nil {
		fs.debugMessage(fmt.Sprintf(ofsmessages.ERR_ACK, err.Error()))
	}

	if fs.debug {
		fs.printer.SucMsg(ofsmessages.UPLOAD_COMPLETE)
	}

	return tmpname, nil
}
