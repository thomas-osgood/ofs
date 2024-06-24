package postgresauthenticator

import "database/sql"

type PostGresAuthenticator struct {
	db *sql.DB
}

type PostGresAuthOption struct {
	Dbname   string
	Host     string
	Password string
	Port     int
	User     string
	Schema   string
	SSL      bool
}
