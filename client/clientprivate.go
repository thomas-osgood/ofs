package client

import (
	"fmt"
	"strings"
	"sync"

	ofcmessages "github.com/thomas-osgood/ofs/client/internal/messages"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
)

// worker function designed to be used by MultifileDownload
// to spawn multiple go routines to download files in a
// concurrent manner.
func (fc *FClient) mfdWorker(target string, errs *map[string]error, wg *sync.WaitGroup) {
	defer wg.Done()

	var err error

	// if the current filename is an empty string or only
	// contains non-printable characters, ignore it and
	// move to the next filename.
	target = strings.TrimSpace(target)
	if len(target) < 1 {
		(*errs)[target] = fmt.Errorf(ofcmessages.ERR_FILENAME_EMPTY)
		return
	}

	// attempt to download the file. if an error occurs,
	// attach it to the filename via the map.
	err = fc.DownloadFile(&filehandler.FileRequest{Filename: target})
	if err != nil {
		(*errs)[target] = err
	}
}
