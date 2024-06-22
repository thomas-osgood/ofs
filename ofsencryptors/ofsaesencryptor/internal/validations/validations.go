package validations

import (
	"fmt"

	consts "github.com/thomas-osgood/ofs/ofsencryptors/ofsaesencryptor/internal/constants"
	encmsg "github.com/thomas-osgood/ofs/ofsencryptors/ofsaesencryptor/internal/messages"
)

// function designed to validate an AES encryptor's key.
func ValidateKey(key []byte) (err error) {
	var keylen int = len(key)

	if (keylen != consts.AES_128_KEYLENGTH) && (keylen != consts.AES_192_KEYLENGTH) && (keylen != consts.AES_256_KEYLENGTH) {
		return fmt.Errorf(encmsg.ERR_KEYLEN)
	}

	return nil
}
