package ofctypes

import (
	"sync"
	"time"
)

type TransferConfig struct {
	downloadConfig
	uploadConfig

	Timeout time.Duration
}

type downloadConfig struct {
	ActiveDownloads int
	MaxDownloads    int
	DownMut         *sync.Mutex
	DownSem         chan struct{}
}

type uploadConfig struct {
	ActiveUploads int
	MaxUploads    int
	UpMut         *sync.Mutex
	UpSem         chan struct{}
}
