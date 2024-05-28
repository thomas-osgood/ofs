package server

import (
	"crypto/tls"

	"github.com/thomas-osgood/OGOR/networking/validations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// create, initialize and return a new grpc server configuration.
func NewGrpcConfiguration(opts ...GrpcOptFunc) (config *GrpcConfig, err error) {
	var defaults GrpcConfigOpt = GrpcConfigOpt{
		Listenaddr: DEFAULT_LISTENADDR,
		Listenport: DEFAULT_LISTENPORT,
		Options:    make([]grpc.ServerOption, 0),
		TLSCert:    tls.Certificate{},
	}
	var opt GrpcOptFunc

	for _, opt = range opts {
		err = opt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	// if the user sepecifies a TLS certificate to use, add it
	// to the gRPC options slice.
	if defaults.TLSCert.Certificate != nil {
		defaults.Options = append(defaults.Options, grpc.Creds(credentials.NewServerTLSFromCert(&defaults.TLSCert)))
	}

	config = new(GrpcConfig)
	config.listenaddr = defaults.Listenaddr
	config.listenport = defaults.Listenport
	config.options = defaults.Options

	return config, nil
}

// set the listen address the gRPC server will listen on.
func WithListenaddr(addr string) GrpcOptFunc {
	return func(gco *GrpcConfigOpt) error {
		gco.Listenaddr = addr
		return nil
	}
}

// set the port the gRPC server will listen on.
//
// must be within range 1 - 65535.
func WithListenport(portno int) GrpcOptFunc {
	return func(gco *GrpcConfigOpt) (err error) {
		err = validations.ValidateNetworkPort(portno)
		if err != nil {
			return err
		}

		gco.Listenport = portno

		return nil
	}
}

// set the grpc server options.
func WithGrpcOptions(opts []grpc.ServerOption) GrpcOptFunc {
	return func(gco *GrpcConfigOpt) error {
		gco.Options = opts
		return nil
	}
}

// add a TLS cert to the server configuration.
func WithTLSCert(cert tls.Certificate) GrpcOptFunc {
	return func(gco *GrpcConfigOpt) error {
		gco.TLSCert = cert
		return nil
	}
}
