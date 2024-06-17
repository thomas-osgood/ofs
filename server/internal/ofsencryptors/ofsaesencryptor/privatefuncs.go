package ofsaesencryptor

import (
	"fmt"

	ofsmessages "github.com/thomas-osgood/ofs/server/internal/messages"
)

// function designed to validate an AES encryptor's key.
func validateKey(key []byte) (err error) {
	var keylen int = len(key)

	if (keylen != 32) && (keylen != 24) && (keylen != 16) {
		return fmt.Errorf(ofsmessages.ERR_KEYLEN_AES)
	}

	return nil
}
