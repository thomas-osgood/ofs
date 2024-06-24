package postgresauthenticator

import "context"

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
