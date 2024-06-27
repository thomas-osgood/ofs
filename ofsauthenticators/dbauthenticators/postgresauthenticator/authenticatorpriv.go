package postgresauthenticator

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	dbamessages "github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators/internal/messages"
	"google.golang.org/grpc/metadata"
)

// function designed to ping the database and make sure the
// connection is valid.
//
// if the connection is valid, nil will be returned.
func (pga *PostGresAuthenticator) checkConnection() (err error) {
	var cancel context.CancelFunc
	var ctx context.Context

	ctx, cancel = context.WithTimeout(context.Background(), pga.timeout)
	defer cancel()

	return pga.db.PingContext(ctx)
}

// generic function designed to execute a given SQL query and
// return a single two column string row.
func (pga *PostGresAuthenticator) execStringQuery(query string, sqlargs []any) (value1 string, value2 string, err error) {
	var cancel context.CancelFunc
	var statement *sql.Stmt
	var str1 sql.NullString
	var str2 sql.NullString
	var tx *sql.Tx
	var txctx context.Context

	txctx, cancel = context.WithTimeout(context.Background(), time.Duration(pga.timeout))
	defer cancel()

	if tx, err = pga.db.BeginTx(txctx, nil); err != nil {
		return "", "", fmt.Errorf(dbamessages.ERR_TX_CONN, err.Error())
	}

	if statement, err = tx.Prepare(query); err != nil {
		return "", "", fmt.Errorf(dbamessages.ERR_SQL_PREP, err.Error())
	}
	defer statement.Close()

	err = statement.QueryRow(sqlargs...).Scan(&str1, &str2)
	if err != nil {
		return "", "", err
	}

	if str1.Valid {
		value1 = str1.String
	} else {
		return "", "", fmt.Errorf(dbamessages.ERR_COLUMN_NULL)
	}

	if str2.Valid {
		value2 = str2.String
	} else {
		return "", "", fmt.Errorf(dbamessages.ERR_COLUMN_NULL)
	}

	return value1, value2, nil
}

// function designed to read the username and password information
// from the provided metadata.
func (pga *PostGresAuthenticator) readMetadataInfo(md metadata.MD) (username string, password string, err error) {
	password = strings.TrimSpace(md.Get(pga.headers.HDRPassword)[0])
	username = strings.TrimSpace(md.Get(pga.headers.HDRUsername)[0])

	if len(password) < 1 {
		return "", "", fmt.Errorf(dbamessages.ERR_PASSWORD_BLANK)
	}

	if len(username) < 1 {
		return "", "", fmt.Errorf(dbamessages.ERR_USERNAME_BLANK)
	}

	return username, password, nil
}
