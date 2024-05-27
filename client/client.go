package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/thomas-osgood/OGOR/protobufs/definitions/common"
	"github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	"github.com/thomas-osgood/OGOR/protobufs/general"
	"google.golang.org/grpc"
)

// ease-of-use function that creates the gRPC connection and a
// filehandler client to communicate with the server.
func (fc *FClient) initConnection() (conn *grpc.ClientConn, client filehandler.FileserviceClient, err error) {

	conn, err = grpc.Dial(fc.srvaddr, fc.srvopts...)
	if err != nil {
		return nil, nil, err
	}

	client = filehandler.NewFileserviceClient(conn)

	return conn, client, nil
}

// function designed to download the file contents from the file server.
//
// this will save the content to the filename specified in the req object.
func (fc *FClient) DownloadFile(req *filehandler.FileRequest) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var fptr *os.File
	var uploader filehandler.Fileservice_UploadFileClient

	conn, client, err = fc.initConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	fptr, err = os.OpenFile(req.GetFilename(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		log.Printf("[OPENFILE] %s\n", err.Error())
		return err
	}
	defer fptr.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	uploader, err = client.UploadFile(ctx, req)
	if err != nil {
		log.Printf("[UPLOAD] %s\n", err.Error())
		return err
	}

	err = general.ReceiveFileBytes(uploader, fptr)
	if err != nil {
		// close the file pointer and remove the empty file,
		// then return the error that was thrown during the
		// transfer process.
		fptr.Close()
		os.Remove(req.GetFilename())
		return err
	}

	err = uploader.CloseSend()
	if err != nil {
		return err
	}

	return nil
}

// function designed to upload a file to the file server.
func (fc *FClient) UploadFile(filename string) (err error) {
	var cancel context.CancelFunc
	var client filehandler.FileserviceClient
	var conn *grpc.ClientConn
	var ctx context.Context
	var fptr *os.File
	var srv filehandler.Fileservice_DownloadFileClient
	var status *common.StatusMessage

	conn, client, err = fc.initConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	fptr, err = os.Open(filename)
	if err != nil {
		return err
	}
	defer fptr.Close()

	ctx, cancel = context.WithTimeout(context.Background(), fc.timeout)
	defer cancel()

	ctx, err = addFilenameMD(ctx, filename)
	if err != nil {
		return err
	}
	srv, err = client.DownloadFile(ctx)
	if err != nil {
		return err
	}

	err = general.TransmitFileBytes(srv, bufio.NewReader(fptr))
	if err != nil {
		return err
	}

	status, err = srv.CloseAndRecv()
	if err != nil {
		return err
	} else if status.GetCode() != http.StatusOK {
		return fmt.Errorf("error transmitting file: %s", status.GetMessage())
	}

	return nil
}
