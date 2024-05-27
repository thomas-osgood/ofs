package server

import (
	"github.com/thomas-osgood/OGOR/output"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	"google.golang.org/grpc"
)

type FServer struct {
	filehandler.UnimplementedFileserviceServer

	chunksize int
	debug     bool
	printer   *output.Outputter
	rootdir   string
}

type FServerOption struct {
	Chunksize int
	Debug     bool
	Rootdir   string
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
}
