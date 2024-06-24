package postgresauthenticator

type PostGresAuthenticator struct {
	dbname string
	schema string
	ssl    bool
}

type PostGresAuthOption struct {
	Dbname string
	Schema string
	SSL    bool
}
