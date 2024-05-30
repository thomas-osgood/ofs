package server

import (
	"bufio"
	"fmt"
	"os"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	ofsdefaults "github.com/thomas-osgood/ofs/server/internal/defaults"
	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
	ofsutils "github.com/thomas-osgood/ofs/server/internal/utils"
)

// function designed to download a file from a client to the server.
//
// this will save the file uploaded by the client to the root directory
// or downloads directory (if one has been set).
func (fs *FServer) DownloadFile(srv filehandler.Fileservice_DownloadFileServer) (err error) {
	var filename string
	var tmpname string

	fs.debugMessage(ofsmessages.DBG_IN_DOWNLOAD)

	filename, err = ofsutils.ReadFilenameMD(srv.Context())
	if err != nil {
		fs.debugMessageErr(fmt.Sprintf(ofsmessages.ERR_MD, err.Error()))
		return err
	}
	filename = fs.cleanFilename(filename, ofsdefaults.FTYPE_DOWNLOAD)

	fs.debugMessageSuc(fmt.Sprintf(ofsmessages.DBG_FILENAME, filename))

	// read the data stream and save it to a temporary file.
	tmpname, err = fs.readIncomingFile(srv)
	if err != nil {
		return err
	}

	// move over the tmpfile contents to the destination file.
	err = fs.moveTempfile(tmpname, filename)
	if err != nil {
		return err
	}

	fs.debugMessageSuc(ofsmessages.TEMP_REMOVED)

	return nil
}

// function designed to upload a requested file from the server to the client.
//
// the specified file must exist in the root directory (or the uploads directory
// if one is specified) for a successful (nil error) return.
func (fs *FServer) UploadFile(req *filehandler.FileRequest, srv filehandler.Fileservice_UploadFileServer) (err error) {
	var fptr *os.File
	var targetfile string = fs.cleanFilename(req.GetFilename(), ofsdefaults.FTYPE_UPLOAD)

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

	fs.debugMessageSuc("file successfully transmitted")

	return nil
}
