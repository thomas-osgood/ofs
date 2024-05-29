package client

import (
	"context"

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
