package utils

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"

	ofscommon "github.com/thomas-osgood/ofs/general"
	"github.com/thomas-osgood/ofs/server/internal/messages"
	"google.golang.org/grpc/metadata"
)

// function designed to check to make sure the specified directory
// has the correct permission and exists.
//
// if the directory does not exist, the function will attempt to
// create the directory.
func CheckDirPerms(dirpath string) (err error) {
	var fi fs.FileInfo

	fi, err = os.Stat(dirpath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dirpath, os.FileMode(0755))
			if err != nil {
				log.Printf(messages.ERR_MKDIR, err.Error())
				return err
			}
			return nil
		} else {
			return err
		}
	}

	if !fi.IsDir() {
		return fmt.Errorf(messages.ERR_PATH_DIR)
	} else if (fi.Mode() & os.ModePerm) == os.ModePerm {
		return fmt.Errorf(messages.ERR_PRIVS_DIR)
	}

	return nil
}

// function designed to read the filename header value from
// the given context's metadata. if the header cannot be found,
// or is empty, an error will be returned.
func ReadFilenameMD(ctx context.Context) (string, error) {
	var md metadata.MD
	var ok bool
	var tmp []string

	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf(messages.ERR_HEADER_METADATA)
	}

	tmp = md.Get(ofscommon.HEADER_FILENAME)
	if (tmp == nil) || (len(tmp) < 1) {
		return "", fmt.Errorf(messages.ERR_HEADER_FILENAME)
	}

	return tmp[0], nil
}
