package ofctypes

import "time"

type TransferConfig struct {
	downloadConfig
	uploadConfig

	Timeout time.Duration
}

type downloadConfig struct {
	ActiveDownloads int
	MaxDownloads    int
}

type uploadConfig struct {
	ActiveUploads int
	MaxUploads    int
}
