package general

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	ofscommon "github.com/thomas-osgood/ofs/common"
	protocommon "github.com/thomas-osgood/ofs/protobufs/common"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	"google.golang.org/grpc/status"
)

// function designed to transmit bytes over the wire using
// ByteString objects.
func TransmitByteStrings(transmitter ByteTransmitter, byteReader *bytes.Reader) (err error) {
	var bytesread int
	var currentBlock []byte = make([]byte, ofscommon.DEFAULT_CHUNKSIZE)

	// iterate through the slice of bytes and transmit
	// each buffer to the receiver. if an error is encountered,
	// return it unless it is an EOF error.
	for {
		bytesread, err = byteReader.Read(currentBlock)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		err = transmitter.Send(&protocommon.ByteString{Chunk: currentBlock[:bytesread]})
		if err != nil {
			return fmt.Errorf("error sending bytes: %s", status.Convert(err).Message())
		}
	}

	return nil
}

// function designed to read a file and transmit its contents
// to another machine via a Transmitter object. this takes in
// a pointer to a bufio.Reader object which should already be
// pointing to a target file.
func TransmitFileBytes(transmitter Transmitter, scanner *bufio.Reader) (err error) {
	var currentChunk []byte = make([]byte, ofscommon.DEFAULT_CHUNKSIZE)
	var bytesread int

	// loop through file, reading it chunk by chunk and
	// uploading it to the requester.
	for {
		// read the current chunk from the file.
		bytesread, err = scanner.Read(currentChunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error scanning file during upload: %s", err.Error())
		}

		// transmit the current chunk to the requester.
		err = transmitter.Send(&filehandler.FileChunk{Chunk: currentChunk[:bytesread]})
		if err != nil {
			return fmt.Errorf("error transmitting chunk: %s", status.Convert(err).Message())
		}
	}

	return nil
}
