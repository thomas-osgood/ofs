package general

import (
	"fmt"
	"io"
	"os"

	protocommon "github.com/thomas-osgood/ofs/protobufs/common"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	"google.golang.org/grpc/status"
)

// function designed to read ByteStrings from a stream
// and save the data to a file.
func ReceiveByteStrings(receiver ByteReceiver, fptr *os.File) (err error) {
	var currentByteString *protocommon.ByteString

	// read the stream ByteString-by-ByteString and
	// save the data to the file. If End-Of-File is
	// encountered, break the loop and return nil.
	for {
		currentByteString, err = receiver.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf(status.Convert(err).Message())
		}

		_, err = fptr.Write(currentByteString.GetChunk())
		if err != nil {
			return err
		}
	}

	return nil
}

// function designed to read bytes from a stream and
// save them to a file.
func ReceiveFileBytes(receiver Receiver, fptr *os.File) (err error) {
	var currentChunk *filehandler.FileChunk

	// read the stream chunk-by-chunk and save the data
	// to the local file.
	for {
		currentChunk, err = receiver.Recv()
		if err != nil {
			// the stream has been closed. there is no
			// more data to save. break the loop.
			if err == io.EOF {
				break
			}
			return fmt.Errorf(status.Convert(err).Message())
		}

		// write the chunk to the output file. trim the
		// null bytes to insure the file is written correctly
		// and no extraneous bytes are added to the end.
		_, err = fptr.Write(currentChunk.GetChunk())
		if err != nil {
			return err
		}
	}

	return nil
}
