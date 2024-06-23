package ofsauthenticators

// interface defining an object that can be used to
// authenticate client requests to the server.
type OFSAuthenticator interface {
	ValidateUser(interface{}) error
}
