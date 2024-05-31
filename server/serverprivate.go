package server

import (
	"fmt"
	"io/fs"
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
func (fsrv *FServer) debugMessage(message string) {
	if fsrv.debug {
		fsrv.printer.InfMsg(message)
	}
}

// helper function for outputting an error message to STDOUT. this
// will only print output if the debug flag is set.
func (fsrv *FServer) debugMessageErr(message string) {
	if fsrv.debug {
		fsrv.printer.ErrMsg(message)
	}
}

// helper function for outputting a success message to STDOUT. this
// will only print output if the debug flag is set.
func (fsrv *FServer) debugMessageSuc(message string) {
	if fsrv.debug {
		fsrv.printer.SucMsg(message)
	}
}

// function designed to list out and return the files contained
// within the uploads directory.
//
// references:
//
// https://stackoverflow.com/questions/14668850/list-directory-in-go
//
// https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html
func (fsrv *FServer) listUploadsDir() (files []*filehandler.FileInfo, err error) {
	var curfile string
	var targetdir string = filepath.Join(fsrv.rootdir, fsrv.uploadsdir)

	// initialize the return slice to prevent a nil reference error.
	files = make([]*filehandler.FileInfo, 0)

	// traverse uploads directory and subdirectories, gathering the
	// filepaths contained within and appending the file data to
	// the slice that will be returned.
	err = filepath.Walk(targetdir, func(path string, info fs.FileInfo, err error) error {

		// beacuse the path is an absolute path, remove the uploads
		// directory that is located at the beginning of the path.
		// this will leave only the relative path within the uploads dir.
		curfile = strings.Replace(path, targetdir, "", 1)

		// replace the leading path separator.
		curfile = strings.TrimSpace(strings.Replace(curfile, fmt.Sprintf("%c", os.PathSeparator), "", 1))

		// if the string is empty (ie: the path was pointing to the
		// uploads root directory) do nothing.
		if len(curfile) < 1 {
			return nil
		}

		// append file data to the return slice.
		files = append(
			files,
			&filehandler.FileInfo{
				Name:      curfile,
				Sizebytes: info.Size(),
				Isdir:     info.IsDir(),
			},
		)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// function designed to move the contents of a temporary file
// to a specified destination.
//
// this will open the temporary file and close it when the
// function returns.
func (fsrv *FServer) moveTempfile(tmpname string, filename string) (err error) {

	fsrv.debugMessageSuc(ofsmessages.COPY_IN_PROGRESS)

	err = ofscommon.MoveFile(tmpname, filename)
	if err != nil {
		fsrv.debugMessageErr(fmt.Sprintf(ofsmessages.ERR_COPY_FILE, err.Error()))
		return err
	}

	fsrv.debugMessageSuc(ofsmessages.COPY_COMPLETE)

	return nil
}

// function designed to read the file data from the incoming
// file byte stream and save it to a temporary file.
//
// this will create a temporary file containing the data that
// gets uploaded, close the file upon return and return the
// temporary file name so the file can be used later on.
func (fsrv *FServer) readIncomingFile(srv filehandler.Fileservice_DownloadFileServer) (tmpname string, err error) {
	var tmpfile *os.File

	// create a temporary file to hold the uploaded information.
	tmpfile, err = os.CreateTemp("", "download")
	if err != nil {
		fsrv.debugMessage(fmt.Sprintf(ofsmessages.ERR_TEMP, err.Error()))
		return "", err
	}
	defer tmpfile.Close()

	tmpname = tmpfile.Name()

	fsrv.debugMessageSuc(fmt.Sprintf(ofsmessages.UPLOAD_IN_PROGRESS, tmpname))

	// stream the file contents from the client and write them
	// to the temp file.
	err = ofscommon.ReceiveFileBytes(srv, tmpfile)
	if err != nil {
		fsrv.debugMessageErr(fmt.Sprintf(ofsmessages.ERR_RECV, err.Error()))
		return "", err
	}

	// send a successful status message and close the stream.
	err = srv.SendAndClose(&protocommon.StatusMessage{Message: ofsmessages.UPLOAD_COMPLETE, Code: http.StatusOK})
	if err != nil {
		fsrv.debugMessageErr(fmt.Sprintf(ofsmessages.ERR_ACK, err.Error()))
	}

	fsrv.debugMessageSuc(ofsmessages.UPLOAD_COMPLETE)

	return tmpname, nil
}
