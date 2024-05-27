package common

import (
	"bufio"
	"io"
	"os"
)

// function designed to copy the contents of a source file
// to a specified destination file.
//
// note: the "destination" file will be opened in WRITE mode,
// meaning if a file with that name already exists it will
// be overwritten.
func CopyFile(source *os.File, destination string) (err error) {
	var chunk []byte = make([]byte, DEFAULT_CHUNKSIZE)
	var fptr *os.File
	var readcount int
	var reader *bufio.Reader

	// attempt to open the destination file for writing.
	fptr, err = os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer fptr.Close()

	reader = bufio.NewReader(source)

	for {
		// read the current chunk from the source file.
		//
		// if and EOF error occurs, there is no more data to
		// read, so break the loop. if a different error is
		// encountered, return.
		readcount, err = reader.Read(chunk)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		// write the bytes read above to the destination file.
		//
		// if there is an error writing to the destination, an
		// error will be returned.
		_, err = fptr.Write(chunk[:readcount])
		if err != nil {
			return err
		}
	}

	return nil
}
