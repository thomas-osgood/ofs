package postgresauthenticator

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/thomas-osgood/ofs/ofsauthenticators"
	"github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators"
	dbadefaults "github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators/internal/defaults"
	dbamessages "github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators/internal/messages"
	pgaconsts "github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators/postgresauthenticator/internal/constants"
	"github.com/thomas-osgood/ofs/ofsauthenticators/hashers/ofsnohash"

	_ "github.com/lib/pq"
)

// function designed to create, initialize and return a
// new PostGresAuthenticator object.
func NewPostGresAuthenticator(opts ...PostGresAuthOptFunc) (pga *PostGresAuthenticator, err error) {
	var connstr string
	var curopt PostGresAuthOptFunc
	var defaults PostGresAuthOption = PostGresAuthOption{
		Hasher: nil,
		Headers: ofsauthenticators.MetadataInfo{
			HDRUsername: dbadefaults.DEFAULT_HEADER_USERNAME,
			HDRPassword: dbadefaults.DEFAULT_HEADER_PASSWORD,
		},
		TableInfo: dbauthenticators.AuthTableInfo{
			Tablename:  dbadefaults.DEFAULT_AUTHTABLE,
			Passcolumn: dbadefaults.DEFAULT_AUTHPASSCOL,
			Usercolumn: dbadefaults.DEFAULT_AUTHUSERCOL,
		},
	}
	var sslstr string

	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	if defaults.Hasher == nil {
		defaults.Hasher, err = ofsnohash.NewNoHasher()
		if err != nil {
			return nil, err
		}
	}

	switch defaults.SSL {
	case true:
		sslstr = pgaconsts.SSL_ENABLE
	default:
		sslstr = pgaconsts.SSL_DISABLE
	}

	connstr = fmt.Sprintf(
		pgaconsts.CONNECTION_STRING,
		defaults.Host,
		defaults.Port,
		defaults.User,
		defaults.Password,
		defaults.Dbname,
		sslstr,
		defaults.Schema,
	)

	pga = new(PostGresAuthenticator)
	pga.hasher = defaults.Hasher
	pga.headers = defaults.Headers
	pga.tableinfo = defaults.TableInfo

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

// set the auth table information.
func WithAuthTable(tableinfo dbauthenticators.AuthTableInfo) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		if tableinfo.IsNil() {
			return fmt.Errorf(dbamessages.ERR_TABLEINFO_EMPTY)
		}

		pgao.TableInfo = tableinfo

		return nil
	}
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

// set the hasher the authenticator will use when comparing
// passwords and hashes.
func WithHasher(hasher ofsauthenticators.Hasher) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		pgao.Hasher = hasher
		return nil
	}
}

// set the username/password headers the authenticator should
// read from the metadata.
func WithHeaders(headers ofsauthenticators.MetadataInfo) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		headers.HDRUsername = strings.TrimSpace(headers.HDRUsername)
		headers.HDRPassword = strings.TrimSpace(headers.HDRPassword)

		if (len(headers.HDRPassword) < 1) || (len(headers.HDRUsername) < 1) {
			return fmt.Errorf(dbamessages.ERR_HEADER_BLANK)
		}

		pgao.Headers = headers
		return nil
	}
}

// set the database host.
func WithHost(host string) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		host = strings.TrimSpace(host)
		if len(host) < 1 {
			return fmt.Errorf(dbamessages.ERR_DBHOST_BLANK)
		}
		pgao.Host = host
		return nil
	}
}

// set db user.
func WithPassword(password string) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		password = strings.TrimSpace(password)
		if len(password) < 1 {
			return fmt.Errorf(dbamessages.ERR_DBPASS_BLANK)
		}
		pgao.Password = password
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

// set connection timeout.
func WithTimeout(timeout time.Duration) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		if timeout < 1*time.Second {
			return fmt.Errorf(dbamessages.ERR_TIMEOUT_SMALL)
		}
		pgao.Timeout = timeout
		return nil
	}
}

// set db user.
func WithUser(username string) PostGresAuthOptFunc {
	return func(pgao *PostGresAuthOption) error {
		username = strings.TrimSpace(username)
		if len(username) < 1 {
			return fmt.Errorf(dbamessages.ERR_DBUSER_BLANK)
		}
		pgao.User = username
		return nil
	}
}
