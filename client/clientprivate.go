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
	// enter critical section.
	fc.transferCfg.DownMut.Lock()
	defer fc.transferCfg.DownMut.Unlock()

	// decrease active downloads and semaphore.
	fc.transferCfg.ActiveDownloads--
	<-fc.transferCfg.DownSem
}

// function designed to deccrement the number of "activeuploads"
// for the client.
func (fc *FClient) decreaseActiveUploads() {
	// enter critical section.
	fc.transferCfg.UpMut.Lock()
	defer fc.transferCfg.UpMut.Unlock()

	// decrease active uploads and semaphore.
	fc.transferCfg.ActiveUploads--
	<-fc.transferCfg.UpSem
}

// function designed to increment the number of "activedownloads"
// for the client.
func (fc *FClient) increaseActiveDownloads() {
	// wait for room in semaphore, then increase
	// the semaphore and active uploads.
	fc.transferCfg.DownSem <- struct{}{}

	// enter critical section.
	fc.transferCfg.DownMut.Lock()
	defer fc.transferCfg.DownMut.Unlock()

	fc.transferCfg.ActiveDownloads++
}

// function designed to increment the number of "activeuploads"
// for the client.
func (fc *FClient) increaseActiveUploads() {
	// wait for room in semaphore, then increase
	// the semaphore and active uploads.
	fc.transferCfg.UpSem <- struct{}{}

	// enter critical section.
	fc.transferCfg.UpMut.Lock()
	defer fc.transferCfg.UpMut.Unlock()

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
