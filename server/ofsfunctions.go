package server

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/thomas-osgood/OGOR/output"
	ofsdefaults "github.com/thomas-osgood/ofs/server/internal/defaults"
	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
	ofsutils "github.com/thomas-osgood/ofs/server/internal/utils"
)

// create, initialize and return a new instance of a file server object.
//
// for Rootdir: if no rootdir is specified, the root directory will be
// set to the current working directory. this can be set using the
// WithDirRoot function. the WithDirRoot function will automatically detect
// if the specified directory is an absolute or relative path. if the path
// is relative it will be prepended with the current working directory.
//
// for Uploadsdir and Downloadsdir: if no value is specified, no subdirectory
// will be created within the server's root directory. these can be set
// using the WithDownloadsDir and WithUploadsDir functions.
func NewOFS(opts ...FSrvOptFunc) (srv *FServer, err error) {
	var defaults FServerOption
	var opt FSrvOptFunc
	var rootdir string

	rootdir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	defaults = FServerOption{
		Debug:           ofsdefaults.DEFAULT_DEBUG,
		Downloadsdir:    ofsdefaults.DIR_DOWNLOADS,
		MaxDownloads:    ofsdefaults.DEFAULT_MAX_DOWNLOADS,
		MaxUploads:      ofsdefaults.DEFAULT_MAX_UPLOADS,
		Rootdir:         rootdir,
		TransferTimeout: ofsdefaults.DEFAULT_TRANSFER_TIMEOUT,
		Uploadsdir:      ofsdefaults.DIR_UPLOADS,
	}

	// set the user-defined configuration options.
	for _, opt = range opts {
		err = opt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	// make sure the specified root directory is able
	// to be written to. if the user running the server
	// does not have write permissions an error will be
	// returned.
	err = ofsutils.CheckDirPerms(defaults.Rootdir)
	if err != nil {
		return nil, err
	}

	// assign the server configuration to the server object
	// that will be returned.
	srv = new(FServer)
	srv.debug = defaults.Debug
	srv.downloadsdir = defaults.Downloadsdir
	srv.rootdir = defaults.Rootdir
	srv.transferCfg.ActiveDownloads = 0
	srv.transferCfg.ActiveUploads = 0
	srv.transferCfg.DownMut = new(sync.Mutex)
	srv.transferCfg.DownSem = make(chan struct{}, defaults.MaxDownloads)
	srv.transferCfg.TransferTimeout = defaults.TransferTimeout
	srv.transferCfg.UpMut = new(sync.Mutex)
	srv.transferCfg.UpSem = make(chan struct{}, defaults.MaxUploads)
	srv.uploadsdir = defaults.Uploadsdir

	srv.printer, err = output.NewOutputter()
	if err != nil {
		return nil, err
	}

	return srv, nil
}

// turn server debug mode on. this will cause messages to be printed
// to STDOUT while the server runs.
func WithDebug() FSrvOptFunc {
	return func(fo *FServerOption) error {
		fo.Debug = true
		return nil
	}
}

// specify root directory for files.
//
// this will automatically detect whether a path is relative or absolute.
func WithDirRoot(dirpath string) FSrvOptFunc {
	return func(fo *FServerOption) (err error) {
		var absdir string
		var curdir string

		if filepath.IsAbs(dirpath) {
			absdir = dirpath
		} else {
			curdir, err = os.Getwd()
			if err != nil {
				return err
			}

			absdir = filepath.Join(curdir, dirpath)
		}

		absdir = filepath.Clean(absdir)

		err = ofsutils.CheckDirPerms(absdir)
		if err != nil {
			return err
		}

		fo.Rootdir = absdir

		return nil
	}
}

// set the downloads directory within the root directory.
func WithDownloadsDir(dirname string) FSrvOptFunc {
	return func(fo *FServerOption) error {
		dirname = ofsutils.CleanDirname(dirname)
		if len(dirname) < 1 {
			return fmt.Errorf(ofsmessages.ERR_DIRSTRING_EMPTY)
		}

		fo.Downloadsdir = dirname

		return nil
	}
}

// set the maximum number of concurrent downloads allowed
// at one time.
func WithMaxDownloads(max int) FSrvOptFunc {
	return func(fo *FServerOption) error {
		if max < 1 {
			return fmt.Errorf(ofsmessages.ERR_TRANSFER_MIN)
		}

		fo.MaxDownloads = max

		return nil
	}
}

// set the maximum number of concurrent uploads allowed
// at one time.
func WithMaxUploads(max int) FSrvOptFunc {
	return func(fo *FServerOption) error {
		if max < 1 {
			return fmt.Errorf(ofsmessages.ERR_TRANSFER_MIN)
		}

		fo.MaxUploads = max

		return nil
	}
}

// set the transfer timeout value for the server.
func WithTransferTimeout(timeout time.Duration) FSrvOptFunc {
	return func(fo *FServerOption) error {
		if timeout < (1 * time.Second) {
			return fmt.Errorf(ofsmessages.ERR_TIMEOUT_VALUE)
		}

		fo.TransferTimeout = timeout

		return nil
	}
}

// set the uploads directory within the root directory.
func WithUploadsDir(dirname string) FSrvOptFunc {
	return func(fo *FServerOption) error {
		dirname = ofsutils.CleanDirname(dirname)
		if len(dirname) < 1 {
			return fmt.Errorf(ofsmessages.ERR_DIRSTRING_EMPTY)
		}

		fo.Uploadsdir = dirname

		return nil
	}
}
