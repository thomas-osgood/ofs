package server

import (
	"context"
	"fmt"

	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
	"google.golang.org/grpc/metadata"
)

// function designed to validate a request.
//
// if the request is deemed "valid", nil will be returned.
func (fsrv *FServer) validateRequest(ctx context.Context) (err error) {
	var md metadata.MD
	var ok bool

	// extract the metadata information from the incoming context
	// that was passed in.
	//
	// if this is unreadable, return an error saying the metadata
	// could not be read.
	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf(ofsmessages.ERR_HEADER_METADATA)
	}

	// if an authenticator has been set, validate the user's identity.
	// if this validation fails, return the error.
	if fsrv.authenticator != nil {
		if err = fsrv.authenticator.ValidateUser(md); err != nil {
			return err
		}
	}

	return nil
}
