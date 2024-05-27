package server

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/thomas-osgood/OGOR/protobufs/definitions/common"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	"github.com/thomas-osgood/OGOR/protobufs/general"
	ofsc "github.com/thomas-osgood/ofs/common"
	"github.com/thomas-osgood/ofs/server/internal/utils"
)

// function designed to clean an uploaded filename and return
// only the filename portion of it. this will strip the directory
// information.
func (fs *FServer) cleanFilename(filename string) (cleaned string) {
	var fnsplit []string

	filename = filepath.Clean(filename)
	fnsplit = strings.Split(filename, fmt.Sprintf("%c", os.PathSeparator))
	cleaned = filepath.Join(fs.rootdir, fnsplit[len(fnsplit)-1])

	return cleaned
}

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

	fs.debugMessage("server in download file function ...")

	filename, err = utils.ReadFilenameMD(srv.Context())
	if err != nil {
		fs.debugMessage(fmt.Sprintf("[MD] %s", err.Error()))
		return err
	}
	filename = fs.cleanFilename(filename)

	if fs.debug {
		fs.printer.SucMsg(fmt.Sprintf("filename: %s", filename))
	}

	// create a temporary file to hold the uploaded information.
	tmpfile, err = os.CreateTemp("", "download")
	if err != nil {
		fs.debugMessage(fmt.Sprintf("[TEMP] %s", err.Error()))
		return err
	}
	defer tmpfile.Close()

	tmpname = tmpfile.Name()

	if fs.debug {
		fs.printer.SysMsgNB(fmt.Sprintf("uploading data to temp file \"%s\"", tmpname))
	}

	// stream the file contents from the client and write them
	// to the temp file.
	err = general.ReceiveFileBytes(srv, tmpfile)
	if err != nil {
		fs.debugMessage(fmt.Sprintf("[RECV] %s", err.Error()))
		return err
	}
	err = srv.SendAndClose(&common.StatusMessage{Message: "upload successful", Code: http.StatusOK})
	if err != nil {
		fs.debugMessage(fmt.Sprintf("[ACK] %s", err.Error()))
	}

	if fs.debug {
		fs.printer.SucMsg("data upload complete.")
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
		fs.printer.SysMsgNB("copying data to destination ...")
	}

	err = ofsc.CopyFile(tmpfile, filename)
	if err != nil {
		fs.debugMessage(fmt.Sprintf("[COPY] %s", err.Error()))
		return err
	}

	if fs.debug {
		fs.printer.SucMsg("copy complete.")
	}

	tmpfile.Close()

	// delete the temporary file used during the upload process.
	err = os.Remove(tmpname)
	if err != nil {
		fs.debugMessage(fmt.Sprintf("[REMOVETMP] %s", err.Error()))
	}

	if fs.debug {
		fs.printer.SucMsg("temp file removed")
	}

	return nil
}

// function designed to upload a requested file from the server to the client.
func (fs *FServer) UploadFile(req *filehandler.FileRequest, srv filehandler.Fileservice_UploadFileServer) (err error) {
	var abspath string
	var fptr *os.File
	var targetfile string = fs.cleanFilename(req.GetFilename())

	abspath = filepath.Join(fs.rootdir, targetfile)

	if fs.debug {
		fs.debugMessage(fmt.Sprintf("client requesting \"%s\"", abspath))
	}

	fptr, err = os.Open(targetfile)
	if err != nil {
		return err
	}
	defer fptr.Close()

	err = general.TransmitFileBytes(srv, bufio.NewReader(fptr))
	if err != nil {
		return err
	}

	return nil
}
