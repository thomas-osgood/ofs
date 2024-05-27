package interrupts

import (
	"os"
	"sync"

	"google.golang.org/grpc"
)

// function designed to handle when the user enters CTRL+C
// to shutdown the listener. this will gracefull handle the
// shutting down of the listener.
func HandleKeyboardInterrupt(shutdownChan chan os.Signal, gsrv *grpc.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	// wait for the keyboard interrupt signal to be triggered.
	<-shutdownChan

	// immediately shutdown all gRPC connections. do not wait
	// for the user to enter "exit" if they are currently
	// connected to a client.
	gsrv.Stop()
}
