// module defining the fileserver. this holds all the structs and
// configuration functions needed to get the sever up and running.
package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	ofsinterrupts "github.com/thomas-osgood/ofs/server/internal/interrupts"
	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
	"google.golang.org/grpc"
)

// function designed to start the gRPC file server using the specified
// gRPC configuration.
func RunServer(srv *FServer, cfg *GrpcConfig) (err error) {
	var glisten net.Listener
	var gserver *grpc.Server
	var shutdownChan chan os.Signal = make(chan os.Signal, 1)
	var srvaddr string = fmt.Sprintf("%s:%d", cfg.listenaddr, cfg.listenport)
	var wg sync.WaitGroup

	// setup the directory structure the fileserver will use.
	err = srv.buildDirStructure()
	if err != nil {
		return err
	}

	// this will be used to wait for CTRL+C. if this gets triggered,
	// a signal will be sent to shutdownChan and the listener will
	// be shutdown.
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	// create the grpc server using the configuration options sepcified
	// by the user via the GrpcConfig object.
	gserver = grpc.NewServer(cfg.options...)
	defer gserver.GracefulStop()

	// start the tcp listener that will be used to serve the grpc server.
	glisten, err = net.Listen("tcp", srvaddr)
	if err != nil {
		return err
	}
	defer glisten.Close()

	// register the fileserver with the grpc server.
	// this is how the fileserver would be added to
	// an existing grpc server in non-standalone mode.
	filehandler.RegisterFileserviceServer(gserver, srv)

	// spawn go routine used to handle CTRL+C entered
	// by the user. this will cleanly stop the server.
	wg.Add(1)
	go ofsinterrupts.HandleKeyboardInterrupt(shutdownChan, gserver, &wg)

	if srv.debug {
		log.Printf(ofsmessages.SERVER_LISTEN_INFO, srvaddr)
	}

	// start the grpc server and listen for incoming
	// client connections.
	err = gserver.Serve(glisten)
	if err != nil {
		log.Printf(ofsmessages.ERR_SERVE, err.Error())
		return err
	}

	wg.Wait()

	return nil
}
