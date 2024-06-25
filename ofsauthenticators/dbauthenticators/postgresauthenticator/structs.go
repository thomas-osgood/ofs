package postgresauthenticator

import (
	"database/sql"
	"time"

	"github.com/thomas-osgood/ofs/ofsauthenticators"
	"github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators"
)

type PostGresAuthenticator struct {
	db        *sql.DB
	hasher    ofsauthenticators.Hasher
	headers   ofsauthenticators.MetadataInfo
	tableinfo dbauthenticators.AuthTableInfo
	timeout   time.Duration
}

type PostGresAuthOption struct {
	Dbname    string
	Hasher    ofsauthenticators.Hasher
	Headers   ofsauthenticators.MetadataInfo
	Host      string
	Password  string
	Port      int
	User      string
	Schema    string
	SSL       bool
	TableInfo dbauthenticators.AuthTableInfo
	Timeout   time.Duration
}
