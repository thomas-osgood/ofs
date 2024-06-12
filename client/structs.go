package client

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type FClient struct {
	srvaddr         string
	srvopts         []grpc.DialOption
	timeout         time.Duration
	transferTimeout time.Duration
}

type FClientOption struct {
	Creds   credentials.TransportCredentials
	Srvaddr string
	Srvopts []grpc.DialOption
	// timeout value for making connections to the server. this will
	// be applied to both upload and download functions.
	Timeout         time.Duration
	TransferTimeout time.Duration
}
