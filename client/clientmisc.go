package client

import (
	"fmt"
	"time"

	ofcmessages "github.com/thomas-osgood/ofs/client/internal/messages"
)

// function designed to update the address of the server
// the client attempts to connect to when it makes RPC calls.
//
// this will first save the existing address, then set the new
// address, and Ping the new address to make sure it is reachable.
//
// if the new address' Ping fails, the address will be reset to
// the old address.
func (fc *FClient) UpdateServerAddress(newaddress string) (err error) {
	var oldaddress string = fc.srvaddr

	fc.srvaddr = newaddress

	_, err = fc.Ping()
	if err != nil {
		fc.srvaddr = oldaddress
		return err
	}

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
