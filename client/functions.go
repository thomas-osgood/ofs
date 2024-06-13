package client

import (
	"fmt"
	"strings"
	"time"

	ofcmessages "github.com/thomas-osgood/ofs/client/internal/messages"
	"github.com/thomas-osgood/ofs/client/internal/ofctypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// create, intialize and return a new FClient object.
func NewClient(opts ...FClientOptFunc) (client *FClient, err error) {
	var defaults FClientOption = FClientOption{
		Creds:       insecure.NewCredentials(),
		Srvaddr:     DEFAULT_SRVADDR,
		Srvopts:     []grpc.DialOption{},
		Timeout:     DEFAULT_TIMEOUT,
		TransferCfg: ofctypes.TransferConfig{},
	}
	var opt FClientOptFunc

	// initialize transfer config defaults
	defaults.TransferCfg.ActiveDownloads = 0
	defaults.TransferCfg.ActiveUploads = 0
	defaults.TransferCfg.MaxDownloads = DEFAULT_MAX_DOWNLOADS
	defaults.TransferCfg.MaxUploads = DEFAULT_MAX_UPLOADS
	defaults.TransferCfg.Timeout = DEFAULT_TRANSFER_TIMEOUT

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
	client.transferCfg = defaults.TransferCfg

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

// set the maximum number of concurrent downloads allowed
func WithMaxDownloads(maxdownloads int) FClientOptFunc {
	return func(fo *FClientOption) error {
		if maxdownloads < 1 {
			return fmt.Errorf(ofcmessages.ERR_MAXTRANSFER_LOW)
		}

		fo.TransferCfg.MaxDownloads = maxdownloads

		return nil
	}
}

// set the maximum number of concurrent downloads allowed
func WithMaxUploads(maxuploads int) FClientOptFunc {
	return func(fo *FClientOption) error {
		if maxuploads < 1 {
			return fmt.Errorf(ofcmessages.ERR_MAXTRANSFER_LOW)
		}

		fo.TransferCfg.MaxUploads = maxuploads

		return nil
	}
}

// set the server address (<ip/domain>:<port>)
func WithSrvAddr(addr string) FClientOptFunc {
	return func(fo *FClientOption) error {
		addr = strings.TrimSpace(addr)
		if len(addr) < 1 {
			return fmt.Errorf(ofcmessages.ERR_EMPTY_ADDR)
		}

		fo.Srvaddr = addr

		return nil
	}
}

// set the connection timeout value.
func WithTimeout(timeout time.Duration) FClientOptFunc {
	return func(fo *FClientOption) error {
		if timeout < (1 * time.Second) {
			return fmt.Errorf(ofcmessages.ERR_NEGATIVE_TIMEOUT)
		}

		fo.Timeout = timeout

		return nil
	}
}

// set the transfer timeout value. this will be the timeout
// value for a file upload and download.
func WithTransferTimeout(timeout time.Duration) FClientOptFunc {
	return func(fo *FClientOption) error {
		if timeout < (1 * time.Second) {
			return fmt.Errorf(ofcmessages.ERR_NEGATIVE_TIMEOUT)
		}

		fo.TransferCfg.Timeout = timeout

		return nil
	}
}
