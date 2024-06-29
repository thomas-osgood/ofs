package defaults

import "time"

const DEFAULT_HEADER_PASSWORD string = "password"
const DEFAULT_HEADER_USERNAME string = "username"

// amount of time to wait until an authentication request times out.
const DEFAULT_TIMEOUT time.Duration = 30 * time.Second
