package server

import (
	"crypto/tls"

	"github.com/thomas-osgood/OGOR/output"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	"google.golang.org/grpc"
)

// fileserver object. this can be added to an existing
// grpc server by registering it with the filehandler
// service via "filehandler.RegisterFileserviceServer".
type FServer struct {
	filehandler.UnimplementedFileserviceServer

	debug        bool
	downloadsdir string
	printer      *output.Outputter
	rootdir      string
	uploadsdir   string
}

// object used to set the confiruation options for
// a new fileserver.
type FServerOption struct {
	// if this is set to true, output will be printed to
	// STDOUT while the server is running.
	Debug bool
	// subdirectory within the rootdir where uploaded
	// files will be saved.
	Downloadsdir string
	// fileserver root directory.
	Rootdir string
	// subdirectory within the rootdir where files that
	// can be downloaded to a client are stored.
	Uploadsdir string
}

type GrpcConfig struct {
	listenaddr string
	listenport int
	options    []grpc.ServerOption
}

type GrpcConfigOpt struct {
	Listenaddr string
	Listenport int
	Options    []grpc.ServerOption
	TLSCert    tls.Certificate
}
