package ofssha256

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	hashmessages "github.com/thomas-osgood/ofs/ofsauthenticators/hashers/internal/messages"
)

// function designed to verify whether the given password
// is valid by comparing it to a hash.
func (sh *SHA256Hasher) ComparePassHash(password []byte, hash []byte) (err error) {
	var passhash []byte

	passhash, err = sh.HashPassword(password)
	if err != nil {
		return err
	}

	switch bytes.Equal(passhash, hash) {
	case false:
		return fmt.Errorf(hashmessages.ERR_NO_MATCH)
	default:
		return nil
	}
}

// function designed to hash and return the sha256 sum of
// a given password.
func (sh *SHA256Hasher) HashPassword(password []byte) (hash []byte, err error) {
	var cur byte
	var sum [32]byte = sha256.Sum256(password)
	hash = make([]byte, 0)
	for _, cur = range sum {
		hash = append(hash, cur)
	}
	return hash, nil
}
