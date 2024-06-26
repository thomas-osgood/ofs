package ofsnohash

import (
	"bytes"
	"fmt"

	nhmessages "github.com/thomas-osgood/ofs/ofsauthenticators/hashers/ofsnohash/internal/messages"
)

// function designed to compare two plaintext passwords.
func (nh *NoHasher) ComparePassHash(password []byte, hash []byte) error {
	switch bytes.Equal(password, hash) {
	case false:
		return fmt.Errorf(nhmessages.ERR_NO_MATCH)
	default:
		return nil
	}
}

// function designed to mimic hashing a password. this will
// return the password back to the caller.
func (nh *NoHasher) HashPassword(password []byte) (hash []byte, err error) {
	return password, nil
}
