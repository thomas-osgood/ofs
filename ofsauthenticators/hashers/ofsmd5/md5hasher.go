package ofsmd5

import (
	"bytes"
	"crypto/md5"
	"fmt"

	hashmessages "github.com/thomas-osgood/ofs/ofsauthenticators/hashers/internal/messages"
)

// function designed to verify whether the given password
// is valid by comparing it to a hash.
func (sh *MD5Hasher) ComparePassHash(password []byte, hash []byte) (err error) {
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

// function designed to hash and return the md5 sum of
// a given password.
func (sh *MD5Hasher) HashPassword(password []byte) (hash []byte, err error) {
	var cur byte
	var sum [16]byte = md5.Sum(password)
	hash = make([]byte, 0)
	for _, cur = range sum {
		hash = append(hash, cur)
	}
	return hash, nil
}
