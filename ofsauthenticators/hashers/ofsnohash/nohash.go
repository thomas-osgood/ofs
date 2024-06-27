// module designed to define a NoHasher hasher that can be used with
// the Osgood File Server's Authenticators. this will fit the Hasher
// interface defined in ofsauthenticators/interfaces.
//
// this hasher does not actually generate any hashes and is meant to
// be used as a default in the case a user does not specify any hasher
// to use with the authenticator.
package ofsnohash

import (
	"bytes"
	"fmt"

	hashmessages "github.com/thomas-osgood/ofs/ofsauthenticators/hashers/internal/messages"
)

// function designed to compare two plaintext passwords.
func (nh *NoHasher) ComparePassHash(password []byte, hash []byte) error {
	switch bytes.Equal(password, hash) {
	case false:
		return fmt.Errorf(hashmessages.ERR_NO_MATCH)
	default:
		return nil
	}
}

// function designed to mimic hashing a password. this will
// return the password back to the caller.
func (nh *NoHasher) HashPassword(password []byte) (hash []byte, err error) {
	return password, nil
}
