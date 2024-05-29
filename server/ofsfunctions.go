package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/thomas-osgood/OGOR/output"
	ofscommon "github.com/thomas-osgood/ofs/internal/general"
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
		Chunksize:    ofscommon.DEFAULT_CHUNKSIZE,
		Debug:        ofsdefaults.DEFAULT_DEBUG,
		Downloadsdir: ofsdefaults.DIR_DOWNLOADS,
		Rootdir:      rootdir,
		Uploadsdir:   ofsdefaults.DIR_UPLOADS,
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
		if size < ofsdefaults.MIN_CHUNKSIZE {
			return fmt.Errorf(ofsmessages.ERR_CHUNK_SMALL)
		} else if size > ofsdefaults.MAX_CHUNKSIZE {
			return fmt.Errorf(ofsmessages.ERR_CHUNK_BIG, ofsdefaults.MAX_CHUNKSIZE)
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
			return fmt.Errorf(ofsmessages.ERR_DIRSTRING_EMPTY)
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
			return fmt.Errorf(ofsmessages.ERR_DIRSTRING_EMPTY)
		}
		dirname = filepath.Clean(dirname)

		fo.Uploadsdir = dirname

		return nil
	}
}
