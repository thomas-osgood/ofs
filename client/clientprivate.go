package client

import (
	"fmt"
	"strings"
	"sync"

	ofcmessages "github.com/thomas-osgood/ofs/client/internal/messages"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
)

// function designed to deccrement the number of "activedownloads"
// for the client.
func (fc *FClient) decreaseActiveDownloads() {
	// acquire the mutex lock and enter the critical section.
	fc.transferCfg.DownMut.Lock()
	defer fc.transferCfg.DownMut.Unlock()
	fc.transferCfg.ActiveDownloads--
}

// function designed to deccrement the number of "activeuploads"
// for the client.
func (fc *FClient) decreaseActiveUploads() {
	// acquire the mutex lock and enter the critical section.
	fc.transferCfg.UpMut.Lock()
	defer fc.transferCfg.UpMut.Unlock()
	fc.transferCfg.ActiveUploads--
}

// function designed to increment the number of "activedownloads"
// for the client.
func (fc *FClient) increaseActiveDownloads() {

	// wait until the active downloads are below the max
	// download count allowed.
	//
	// once the download count is at an acceptable number,
	// acquire the mutex lock and continue.
	for {
		if fc.transferCfg.ActiveDownloads < fc.transferCfg.MaxDownloads {
			// acquire the mutex lock and enter the critical section.
			fc.transferCfg.DownMut.Lock()
			defer fc.transferCfg.DownMut.Unlock()
			break
		}
	}

	fc.transferCfg.ActiveDownloads++
}

// function designed to increment the number of "activeuploads"
// for the client.
func (fc *FClient) increaseActiveUploads() {

	// wait until the active downloads are below the max
	// upload count allowed.
	//
	// once the upload count is at an acceptable number,
	// acquire the mutex lock and continue.
	for {
		if fc.transferCfg.ActiveUploads < fc.transferCfg.MaxUploads {
			// acquire the mutex lock and enter the critical section.
			fc.transferCfg.UpMut.Lock()
			defer fc.transferCfg.UpMut.Unlock()
			break
		}
	}

	fc.transferCfg.ActiveUploads++
}

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
