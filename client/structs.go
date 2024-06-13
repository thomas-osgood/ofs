package client

import (
	"time"

	"github.com/thomas-osgood/ofs/client/internal/ofctypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type FClient struct {
	srvaddr     string
	srvopts     []grpc.DialOption
	timeout     time.Duration
	transferCfg ofctypes.TransferConfig
}

type FClientOption struct {
	Creds       credentials.TransportCredentials
	Srvaddr     string
	Srvopts     []grpc.DialOption
	TransferCfg ofctypes.TransferConfig
	// timeout value for making connections to the server. this will
	// be applied to both upload and download functions.
	Timeout time.Duration
}
