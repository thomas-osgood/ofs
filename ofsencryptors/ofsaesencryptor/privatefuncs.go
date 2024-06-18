package ofsaesencryptor

import (
	"fmt"

	encryptormessages "github.com/thomas-osgood/ofs/ofsencryptors/internal/messages"
)

// function designed to validate an AES encryptor's key.
func validateKey(key []byte) (err error) {
	var keylen int = len(key)

	if (keylen != 32) && (keylen != 24) && (keylen != 16) {
		return fmt.Errorf(encryptormessages.ERR_AES_KEYLEN)
	}

	return nil
}
