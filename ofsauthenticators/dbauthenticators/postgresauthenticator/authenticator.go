package postgresauthenticator

import (
	"fmt"

	pgaqueries "github.com/thomas-osgood/ofs/ofsauthenticators/dbauthenticators/postgresauthenticator/internal/queries"
	"google.golang.org/grpc/metadata"
)

// function designed to validate the user by looking up the
// user's password in the database and comparing the passwords.
func (pga *PostGresAuthenticator) ValidateUser(md metadata.MD) (err error) {
	var hashedpw string
	var password string
	var username string
	var query string = fmt.Sprintf(
		pgaqueries.SELECT_USER_PASSWORD,
		pga.tableinfo.Usercolumn,
		pga.tableinfo.Passcolumn,
		pga.tableinfo.Tablename,
		pga.tableinfo.Usercolumn,
	)

	username, password, err = pga.readMetadataInfo(md)
	if err != nil {
		return err
	}

	_, hashedpw, err = pga.execStringQuery(query, []any{username})
	if err != nil {
		return err
	}

	return pga.hasher.ComparePassHash([]byte(password), []byte(hashedpw))
}
