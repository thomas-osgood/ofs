package postgresauthenticator

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
