package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/thomas-osgood/OGOR/output"
	ofscommon "github.com/thomas-osgood/ofs/common"
	"github.com/thomas-osgood/ofs/server/internal/messages"
	ofsutils "github.com/thomas-osgood/ofs/server/internal/utils"
)

// create, initialize and return a new instance of a file server object.
func NewOFS(opts ...FSrvOptFunc) (srv *FServer, err error) {
	var defaults FServerOption
	var opt FSrvOptFunc
	var rootdir string

	rootdir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	defaults = FServerOption{
		Chunksize:    ofscommon.DEFAULT_CHUNKSIZE,
		Debug:        DEFAULT_DEBUG,
		Downloadsdir: DIR_DOWNLOADS,
		Rootdir:      rootdir,
		Uploadsdir:   DIR_UPLOADS,
	}

	// set the user-defined configuration options.
	for _, opt = range opts {
		err = opt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	err = ofsutils.CheckDirPerms(defaults.Rootdir)
	if err != nil {
		return nil, err
	}

	// assign the server configuration to the server object
	// that will be returned.
	srv = new(FServer)
	srv.chunksize = defaults.Chunksize
	srv.debug = defaults.Debug
	srv.downloadsdir = defaults.Downloadsdir
	srv.rootdir = defaults.Rootdir
	srv.uploadsdir = defaults.Uploadsdir

	srv.printer, err = output.NewOutputter()
	if err != nil {
		return nil, err
	}

	return srv, nil
}

// set the size of each chunk that will be transmitted/received
// during a file transfer.
//
// max chunk size: 65535
func WithChunksize(size int) FSrvOptFunc {
	return func(fo *FServerOption) (err error) {
		if size < MIN_CHUNKSIZE {
			return fmt.Errorf(messages.ERR_CHUNK_SMALL)
		} else if size > MAX_CHUNKSIZE {
			return fmt.Errorf(messages.ERR_CHUNK_BIG, MAX_CHUNKSIZE)
		}

		fo.Chunksize = size

		return nil
	}
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
		dirname = strings.TrimSpace(dirname)
		if len(dirname) < 1 {
			return fmt.Errorf(messages.ERR_DIRSTRING_EMPTY)
		}
		dirname = filepath.Clean(dirname)

		fo.Downloadsdir = dirname

		return nil
	}
}

// set the uploads directory within the root directory.
func WithUploadsDir(dirname string) FSrvOptFunc {
	return func(fo *FServerOption) error {
		dirname = strings.TrimSpace(dirname)
		if len(dirname) < 1 {
			return fmt.Errorf(messages.ERR_DIRSTRING_EMPTY)
		}
		dirname = filepath.Clean(dirname)

		fo.Uploadsdir = dirname

		return nil
	}
}
