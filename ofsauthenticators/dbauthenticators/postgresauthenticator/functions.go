package postgresauthenticator

import (
	"fmt"
	"strings"

	dbamessages "github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators/internal/messages"
)

// function designed to create, initialize and return a
// new PostGresAuthenticator object.
func NewPostGresAuthenticator(opts ...PostGresAuthOptFunc) (pga *PostGresAuthenticator, err error) {
	var curopt PostGresAuthOptFunc
	var defaults PostGresAuthOption = PostGresAuthOption{}

	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	pga = new(PostGresAuthenticator)
	pga.dbname = defaults.Dbname
	pga.schema = defaults.Schema
	pga.ssl = defaults.SSL

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
