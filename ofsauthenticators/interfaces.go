package ofsauthenticators

// interface defining an object that can be used to
// authenticate client requests to the server.
type OFSAuthenticator[T AuthRequest] interface {
	ValidateUser(T) error
}
