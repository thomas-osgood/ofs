package postgresauthenticator

import (
	"database/sql"
	"time"
)

type PostGresAuthenticator struct {
	db      *sql.DB
	timeout time.Duration
}

type PostGresAuthOption struct {
	Dbname   string
	Host     string
	Password string
	Port     int
	User     string
	Schema   string
	SSL      bool
	Timeout  time.Duration
}

type AuthTableInfo struct {
	// name of the table that holds the verification info.
	Tablename string
	// column name of the column that holds the username.
	Usercolumn string
	// column name of the column that holds the password.
	Passcolumn string
}
