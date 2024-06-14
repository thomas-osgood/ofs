// module designed to hold the internal types specific to
// the Osgood File Server.
package ofstypes

import "time"

type TransferConfig struct {
	downloadConfig
	uploadConfig

	TransferTimeout time.Duration
}

type downloadConfig struct {
	ActiveDownloads int
	DownSem         chan struct{}
}

type uploadConfig struct {
	ActiveUploads int
	UpSem         chan struct{}
}
