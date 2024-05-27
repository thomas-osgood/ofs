package client

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// create, intialize and return a new FClient object.
func NewClient(opts ...FClientOptFunc) (client *FClient, err error) {
	var defaults FClientOption = FClientOption{
		Creds:   insecure.NewCredentials(),
		Srvaddr: DEFAULT_SRVADDR,
		Srvopts: []grpc.DialOption{},
		Timeout: DEFAULT_TIMEOUT,
	}
	var opt FClientOptFunc

	for _, opt = range opts {
		err = opt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	defaults.Srvopts = append(defaults.Srvopts, grpc.WithTransportCredentials(defaults.Creds))

	client = new(FClient)
	client.srvaddr = defaults.Srvaddr
	client.srvopts = defaults.Srvopts
	client.timeout = defaults.Timeout

	return client, nil
}

// set the transport credentials the client will use to
// connect to the server.
func WithCreds(creds credentials.TransportCredentials) FClientOptFunc {
	return func(fo *FClientOption) error {
		fo.Creds = creds
		return nil
	}
}

// set the server address (<ip/domain>:<port>)
func WithSrvAddr(addr string) FClientOptFunc {
	return func(fo *FClientOption) error {
		addr = strings.TrimSpace(addr)
		if len(addr) < 1 {
			return fmt.Errorf("address must be non-empty string")
		}

		fo.Srvaddr = addr

		return nil
	}
}

// set the connection timeout value.
func WithTimeout(timeout time.Duration) FClientOptFunc {
	return func(fo *FClientOption) error {
		if timeout < (1 * time.Second) {
			return fmt.Errorf("timeout must be a positive second value")
		}

		fo.Timeout = timeout

		return nil
	}
}
