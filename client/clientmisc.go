package client

import (
	"fmt"
	"time"

	ofcmessages "github.com/thomas-osgood/ofs/client/internal/messages"
)

// function designed to update the address of the server
// the client attempts to connect to when it makes RPC calls.
func (fc *FClient) UpdateServerAddress(newaddress string) (err error) {
	fc.srvaddr = newaddress
	return nil
}

// function designed to update the connection timeout for
// the client.
func (fc *FClient) UpdateTimeout(newtimeout time.Duration) (err error) {
	if newtimeout.Seconds() <= 0 {
		return fmt.Errorf(ofcmessages.ERR_NEGATIVE_TIMEOUT)
	}
	fc.timeout = newtimeout
	return nil
}
