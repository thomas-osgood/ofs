package defaults

import "time"

const DEFAULT_DEBUG bool = false
const DEFAULT_LISTENADDR string = "0.0.0.0"
const DEFAULT_LISTENPORT int = 8999

const DIR_UPLOADS string = ""
const DIR_DOWNLOADS string = ""

const FTYPE_DOWNLOAD string = "download"
const FTYPE_UPLOAD string = "upload"

const DEFAULT_MAX_DOWNLOADS int = 20
const DEFAULT_MAX_UPLOADS int = 20

const DEFAULT_TRANSFER_TIMEOUT time.Duration = 20 * time.Minute
