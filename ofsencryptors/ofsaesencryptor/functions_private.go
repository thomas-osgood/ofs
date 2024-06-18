package ofsaesencryptor

import (
	"fmt"

	encmsg "github.com/thomas-osgood/ofs/ofsencryptors/ofsaesencryptor/internal/messages"
)

// function designed to auto-generate a key based on the key
// length passed in.
//
// the key type is determined via the keytype enum.
func autogenKey(keytype int) (key []byte, err error) {
	switch keytype {
	case AES_128:
	case AES_192:
	case AES_256:
	default:
		return nil, fmt.Errorf(encmsg.ERR_KEYTYP)
	}

	return key, nil
}
