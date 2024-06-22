package client

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/thomas-osgood/ofs/client/internal/ofcconstants"
	genmessages "github.com/thomas-osgood/ofs/internal/messages"
	protocommon "github.com/thomas-osgood/ofs/protobufs/common"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	"google.golang.org/grpc"
)

// function designed to create a filepath inside subdirectories.
//
// reference:
//
// https://stackoverflow.com/questions/59961510/golang-os-create-path-with-nested-directories
func (fc *FClient) createFilepath(fpath string) (fptr *os.File, err error) {
	err = os.MkdirAll(filepath.Dir(fpath), os.FileMode(0770))
	if err != nil {
		return nil, err
	}
	return os.Create(fpath)
}

// function designed to request the encryption or decryption of
// a target file.
func (fc *FClient) cryptoAction(filename string, action int) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var status *protocommon.StatusMessage

	conn, client, err = fc.initConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	switch action {
	case ofcconstants.CRYPTO_DECRYPT:
		status, err = client.DecryptFile(ctx, &filehandler.FileRequest{Filename: filename})
	case ofcconstants.CRYPTO_ENCRYPT:
		status, err = client.EncryptFile(ctx, &filehandler.FileRequest{Filename: filename})
	default:
		err = fmt.Errorf(genmessages.ERR_ACTION_INVALID)
	}

	if err != nil {
		return err
	} else if status.Code >= http.StatusBadRequest {
		return fmt.Errorf(status.GetMessage())
	}

	return nil
}

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
func (fc *FClient) mfdWorker(target string, errs *map[string]error, wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()

	var err error

	// if the current filename is an empty string or only
	// contains non-printable characters, ignore it and
	// move to the next filename.
	target, err = validateFilename(target)
	if err != nil {
		(*errs)[target] = err
		return
	}

	// attempt to download the file. if an error occurs,
	// attach it to the filename via the map.
	err = fc.DownloadFile(&filehandler.FileRequest{Filename: target})
	if err != nil {
		mut.Lock()
		(*errs)[target] = err
		mut.Unlock()
	}
}

// worker function designed to be used by MultifileUpload
// to spawn multiple go routines to upload files in a
// concurrent manner.
func (fc *FClient) mfuWorker(target string, errs *map[string]error, wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()

	var err error

	// make sure the target is not an empty string. if
	// this validation fails, add the error to the errors
	// map and return.
	target, err = validateFilename(target)
	if err != nil {
		(*errs)[target] = err
		return
	}

	err = fc.UploadFile(target)
	if err != nil {
		mut.Lock()
		(*errs)[target] = err
		mut.Unlock()
	}
}
