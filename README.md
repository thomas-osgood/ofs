# Osgood File Server (OFS)

## Overview

The Osgood File Server (OFS) is designed to be a simple file server that utilizes protocol buffers (protobufs) and gRPC to transfer files. The server is written in Golang.

The server can be added to an existing gRPC server (by registering the filehandler via `filehandler.RegisterFileserviceServer`) or can be run as a stand-alone server using the `RunServer` function.

## License

This software is distributed under the [General Public License v3.0](LICENSE). Any and all use of this software should adhere to this license.

## Example

The following code showcases the `RunServer` functionality. This will spawn a new instance of the fileserver in "stand-alone" mode and have a client upload a specified file, then request a file for download.

```go
package main

import (
	"log"
	"time"

	"github.com/thomas-osgood/ofs/client"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	"github.com/thomas-osgood/ofs/server"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var cfg *server.GrpcConfig
	var err error
	var srv *server.FServer

	srv, err = server.NewOFS(server.WithDebug(), server.WithDirRoot("uploads"), server.WithDownloadsDir("in"), server.WithUploadsDir("out"))
	if err != nil {
		log.Fatalf("[NEWOFS] %s\n", err.Error())
	}

	cfg, err = server.NewGrpcConfiguration()
	if err != nil {
		log.Fatalf("[NEWCFG] %s\n", err.Error())
	}

	go func() {
		time.Sleep(5 * time.Second)

		clnt, err := client.NewClient(client.WithTimeout(5*time.Minute), client.WithCreds(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("[CLIENTNEW] %s\n", err.Error())
		}

		err = clnt.UploadFile("test.txt")
		if err != nil {
			log.Fatalf("[CLIENTUP] %s\n", err.Error())
		}

		files, err := clnt.ListFiles()
		if err != nil {
			log.Fatalf("[LISTFILES] %s\n", err.Error())
		}

		log.Printf("[LIST] %d files discovered ...", len(files))
		for _, curfile := range files {
			log.Printf("%s -- %d bytes -- Directory: %t\n", curfile.GetName(), curfile.GetSizebytes(), curfile.GetIsdir())
		}

		err = clnt.MakeDirectory("testdir/with/subdirs")
		if err != nil {
			log.Fatalf("[MKDIR] %s\n", err.Error())
		}

		err = clnt.RenameFile("test.txt", "copydir/testcopy.txt")
		if err != nil {
			log.Fatalf("[COPYFILE] %s\n", err.Error())
		}

		err = clnt.DeleteFile("copydir/testcopy.txt")
		if err != nil {
			log.Fatalf("[DELETEFILE] %s\n", err.Error())
		}

		err = clnt.DownloadFile(&filehandler.FileRequest{Filename: "test.txt"})
		if err != nil {
			log.Fatalf("[CLIENTDOWN] %s\n", err.Error())
		}
	}()

	err = server.RunServer(srv, cfg)
	if err != nil {
		log.Fatalf("[RUNSRV] %s\n", err.Error())
	}
}

```
