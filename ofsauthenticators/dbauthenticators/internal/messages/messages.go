package messages

const DRIVER_POSTGRES string = "postgres"

const ERR_DBHOST_BLANK string = "database host cannot be an empty string"
const ERR_DBNAME_BLANK string = "database name cannot be an empty string"
const ERR_DBPASS_BLANK string = "database password cannot be an empty string"
const ERR_DBUSER_BLANK string = "database user cannot be an empty string"
const ERR_SCHEMA_BLANK string = "schema name cannto be an empty string"
const ERR_TABLEINFO_EMPTY string = "table information cannot be empty"
const ERR_TIMEOUT string = "db connection timeout"
const ERR_TIMEOUT_SMALL string = "timeout value must be at least 1 second"
