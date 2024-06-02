package messages

const COPY_COMPLETE string = "copy complete."
const COPY_IN_PROGRESS string = "copying data to destination ..."

const DBG_FILENAME string = "filename: %s"
const DBG_FILE_REQUEST string = "client requesting \"%s\""
const DBG_IN_DOWNLOAD string = "server in download file function ..."

const DIR_CREATED string = "directory successfully created"

const ERR_ACK string = "[ACK] %s"
const ERR_CHUNK_BIG string = "chunk size must be less than or equal to %d"
const ERR_CHUNK_SMALL string = "chunksize must be larger than zero"
const ERR_COPY_FILE string = "[COPY] %s"
const ERR_FILE_EXISTS string = "destination file already exists"
const ERR_DIRSTRING_EMPTY string = "directory name must be a non-zero length string"
const ERR_HEADER_FILENAME string = "filename not found in metadata"
const ERR_HEADER_METADATA string = "unable to read metadata"
const ERR_MD string = "[MD] %s"
const ERR_MKDIR string = "mkdir error: %s"
const ERR_PATH_DIR string = "specified path is not a directory"
const ERR_PRIVS_DIR string = "insufficient permissions to write to directory"
const ERR_REMOVE_TEMP string = "[REMOVETMP] %s"
const ERR_RECV string = "[RECV] %s"
const ERR_SERVE string = "[SERVE] %s\n"
const ERR_TEMP string = "[TEMP] %s"

const FILE_DELETED string = "file successfully deleted"
const FILE_SENT string = "file successfully transmitted"

const SERVER_LISTEN_INFO string = "server listening on \"%s\"\n"

const TEMP_REMOVED string = "temp file removed"

const UPLOAD_COMPLETE string = "data upload complete."
const UPLOAD_IN_PROGRESS string = "uploading data to temp file \"%s\""
