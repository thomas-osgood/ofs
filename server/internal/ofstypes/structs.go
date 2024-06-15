// module designed to hold the internal types specific to
// the Osgood File Server.
package ofstypes

import (
	"sync"
	"time"
)

type TransferConfig struct {
	downloadConfig
	uploadConfig

	TransferTimeout time.Duration
}

type downloadConfig struct {
	ActiveDownloads int
	DownMut         *sync.Mutex
	DownSem         chan struct{}
}

type uploadConfig struct {
	ActiveUploads int
	UpMut         *sync.Mutex
	UpSem         chan struct{}
}
