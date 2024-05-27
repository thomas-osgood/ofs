package client

import (
	"context"

	"github.com/thomas-osgood/ofs/common"
	"google.golang.org/grpc/metadata"
)

func addFilenameMD(ctx context.Context, filename string) (fctx context.Context, err error) {
	var md metadata.MD = make(metadata.MD)
	md.Set(common.HEADER_FILENAME, filename)
	fctx = metadata.NewOutgoingContext(ctx, md)
	return fctx, nil
}
