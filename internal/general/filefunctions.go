package general

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

	destination = strings.TrimSpace(filepath.Clean(destination))

	// attempt to open the destination file for writing.
	fptr, err = os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		// if an ErrNotExist error is detected, it is assumed
		// that the directory structure specified for the destination
		// file does not exist. this block will create the filepath
		// for the destination file and re-attempt file creation once
		// the directory structure has been created.
		//
		// if the directory structure creation fails, an error will
		// be returned.
		if errors.Is(err, os.ErrNotExist) {

			// create the destination file directory structure.
			err = CreateFileDirs(destination)
			if err != nil {
				return err
			}

			// re-attempt file creation. if this is not
			// done, an error will occur later on in this
			// function's logic because the destination
			// file pointer will be null.
			fptr, err = os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0644))
			if err != nil {
				return err
			}

		} else {
			return err
		}
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

// overload of the CopyFile function that takes in both
// arguments as filepaths.
func CopyFileSS(source string, destination string) (err error) {
	var fptr *os.File

	fptr, err = os.Open(source)
	if err != nil {
		return err
	}
	defer fptr.Close()

	err = CopyFile(fptr, destination)
	if err != nil {
		return err
	}

	return nil
}

// function designed to create all directories necessary for
// a file.
//
// example: if "test/directory/struct/file.txt" is passed in,
// "test/directory/struct" will be created so "file.txt" can
// be created.
func CreateFileDirs(destination string) (err error) {
	var basepath string
	var sepstr string = fmt.Sprintf("%c", os.PathSeparator)
	var splitpath []string

	destination = strings.TrimSpace(filepath.Clean(destination))
	splitpath = strings.Split(destination, sepstr)

	// if the split path is only one element, only the filename
	// has been provided, so there is nothing to do.
	if len(splitpath) < 2 {
		return nil
	}

	basepath = strings.Join(splitpath[:len(splitpath)-1], sepstr)

	err = os.MkdirAll(basepath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// function designed to move a given source file to the
// designated destination. this will attempt to delete the
// source file after the data has been copied over
// to the destination.
func MoveFile(source string, destination string) (err error) {

	// move data from the source to the destination.
	err = CopyFileSS(source, destination)
	if err != nil {
		return err
	}

	// remove the source file.
	//
	// if this fails an error will be returned.
	err = os.Remove(source)
	if err != nil {
		return err
	}

	return nil
}

// function designed to read and return the content of
// a given filename.
func ReadFileBytes(filename string) (content []byte, err error) {
	var fptr *os.File

	fptr, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fptr.Close()

	content, err = io.ReadAll(fptr)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// function designed to create (or overwrite) a given file
// and write the provided content to it.
func WriteFileBytes(filename string, content []byte) (err error) {
	var fptr *os.File

	fptr, err = os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fptr.Close()

	_, err = fptr.Write(content)
	if err != nil {
		return err
	}

	return nil
}
