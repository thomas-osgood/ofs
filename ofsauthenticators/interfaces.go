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

// interface defining an object that can perform hash
// functions to aid in authentication of a user.
type Hasher interface {
	// function designed to take in a password and hash
	// and compare the two.
	//
	// if the password and hash match, nil will be returned.
	ComparePassHash(password []byte, hash []byte) error
	// function designed to calculate and return the hash
	// of a given password.
	HashPassword([]byte) ([]byte, error)
}
