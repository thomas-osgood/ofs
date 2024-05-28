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

		err = clnt.UploadFile("/home/kali/test.txt")
		if err != nil {
			log.Fatalf("[CLIENTUP] %s\n", err.Error())
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
