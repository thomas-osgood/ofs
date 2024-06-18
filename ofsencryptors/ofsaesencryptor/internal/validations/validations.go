package validations

import (
	"fmt"

	encryptormessages "github.com/thomas-osgood/ofs/ofsencryptors/internal/messages"
	consts "github.com/thomas-osgood/ofs/ofsencryptors/ofsaesencryptor/internal/constants"
)

// function designed to validate an AES encryptor's key.
func ValidateKey(key []byte) (err error) {
	var keylen int = len(key)

	if (keylen != consts.AES_128_KEYLENGTH) && (keylen != consts.AES_192_KEYLENGTH) && (keylen != consts.AES_256_KEYLENGTH) {
		return fmt.Errorf(encryptormessages.ERR_AES_KEYLEN)
	}

	return nil
}
