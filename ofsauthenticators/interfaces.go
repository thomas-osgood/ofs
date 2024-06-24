package ofsauthenticators

import "google.golang.org/grpc/metadata"

// interface defining an object that can be used to
// authenticate client requests to the server.
type OFSAuthenticator interface {
	// function designed to parse metadata passed into it,
	// pull out the correct credential headers and verify
	// whether the given credentials are valid for a user.
	ValidateUser(metadata.MD) error
}
