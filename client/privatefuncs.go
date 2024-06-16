package client

import (
	"context"
	"fmt"
	"strings"

	ofcmessages "github.com/thomas-osgood/ofs/client/internal/messages"
	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	"google.golang.org/grpc/metadata"
)

// function designed to add metadata containing a filename
// header to the given context.
//
// this returns an outgoing context with the generated
// metadata attached to it.
func addFilenameMD(ctx context.Context, filename string) (fctx context.Context, err error) {
	var md metadata.MD = make(metadata.MD)
	md.Set(ofscommon.HEADER_FILENAME, filename)
	fctx = metadata.NewOutgoingContext(ctx, md)
	return fctx, nil
}

// function designed to validate a passed in filename.
//
// this will pass the filename to a cleaning function and
// return the cleaned filename if it passes validation;
// otherwise it will return an error.
func validateFilename(target string) (string, error) {

	target = strings.TrimSpace(target)
	if len(target) < 1 {
		return target, fmt.Errorf(ofcmessages.ERR_FILENAME_EMPTY)
	}

	return target, nil
}
