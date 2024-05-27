package utils

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/thomas-osgood/ofs/common"
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
				log.Printf("mkdir error: %s\n", err.Error())
				return err
			}
			return nil
		} else {
			return err
		}
	}

	if !fi.IsDir() {
		return fmt.Errorf("specified path is not a directory")
	} else if (fi.Mode() & os.ModePerm) == os.ModePerm {
		return fmt.Errorf("insufficient permissions to write to directory")
	}

	return nil
}

func ReadFilenameMD(ctx context.Context) (string, error) {
	var md metadata.MD
	var ok bool
	var tmp []string

	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("unable to read metadata")
	}

	tmp = md.Get(common.HEADER_FILENAME)
	if (tmp == nil) || (len(tmp) < 1) {
		return "", fmt.Errorf("filename not found in metadata")
	}

	return tmp[0], nil
}
