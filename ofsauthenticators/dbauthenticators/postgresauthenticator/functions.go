package postgresauthenticator

import (
	"database/sql"
	"fmt"
	"strings"

	dbamessages "github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators/internal/messages"
	pgaconsts "github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators/postgresauthenticator/internal/constants"
)

// function designed to create, initialize and return a
// new PostGresAuthenticator object.
func NewPostGresAuthenticator(opts ...PostGresAuthOptFunc) (pga *PostGresAuthenticator, err error) {
	var connstr string
	var curopt PostGresAuthOptFunc
	var defaults PostGresAuthOption = PostGresAuthOption{}
	var sslstr string

	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	switch defaults.SSL {
	case true:
		sslstr = "enable"
	default:
		sslstr = "disable"
	}

	connstr = fmt.Sprintf(pgaconsts.CONNECTION_STRING, defaults.Host, defaults.Port, defaults.User, defaults.Password, defaults.Dbname, sslstr)

	pga = new(PostGresAuthenticator)

	pga.db, err = sql.Open(dbamessages.DRIVER_POSTGRES, connstr)
	if err != nil {
		return nil, err
	}
	pga.timeout = defaults.Timeout

	if err = pga.checkConnection(); err != nil {
		return nil, fmt.Errorf(dbamessages.ERR_TIMEOUT)
	}

	return pga, nil
}

// set the database name the authenticator will use.
func WithDBName(dbname string) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		dbname = strings.TrimSpace(dbname)
		if len(dbname) < 1 {
			return fmt.Errorf(dbamessages.ERR_DBNAME_BLANK)
		}
		pgao.Dbname = dbname
		return nil
	}
}

// set the database name the authenticator will use.
func WithSchema(schemaname string) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		schemaname = strings.TrimSpace(schemaname)
		if len(schemaname) < 1 {
			return fmt.Errorf(dbamessages.ERR_SCHEMA_BLANK)
		}
		pgao.Schema = schemaname
		return nil
	}
}

// set SSL usage for connection.
func WithSSL(pgao *PostGresAuthOption) error {
	pgao.SSL = true
	return nil
}
