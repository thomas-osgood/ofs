package server

import (
	"crypto/tls"

	"github.com/thomas-osgood/OGOR/output"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	"google.golang.org/grpc"
)

type FServer struct {
	filehandler.UnimplementedFileserviceServer

	chunksize    int
	debug        bool
	downloadsdir string
	printer      *output.Outputter
	rootdir      string
	uploadsdir   string
}

type FServerOption struct {
	Chunksize int
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
