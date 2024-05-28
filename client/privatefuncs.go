package client

import (
	"context"

	ofscommon "github.com/thomas-osgood/ofs/general"
	"google.golang.org/grpc/metadata"
)

func addFilenameMD(ctx context.Context, filename string) (fctx context.Context, err error) {
	var md metadata.MD = make(metadata.MD)
	md.Set(ofscommon.HEADER_FILENAME, filename)
	fctx = metadata.NewOutgoingContext(ctx, md)
	return fctx, nil
}
