package client

import (
	"time"
)

const DEFAULT_MAX_DOWNLOADS int = 10
const DEFAULT_MAX_UPLOADS int = 10
const DEFAULT_SRVADDR string = "127.0.0.1:8999"
const DEFAULT_TIMEOUT time.Duration = 30 * time.Second
const DEFAULT_TRANSFER_TIMEOUT time.Duration = 20 * time.Minute
