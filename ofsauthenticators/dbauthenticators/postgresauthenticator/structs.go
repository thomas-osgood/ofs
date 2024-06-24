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
