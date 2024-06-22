package ofsaesencryptor

import (
	"fmt"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	consts "github.com/thomas-osgood/ofs/ofsencryptors/ofsaesencryptor/internal/constants"
	encmsg "github.com/thomas-osgood/ofs/ofsencryptors/ofsaesencryptor/internal/messages"
)

// function designed to auto-generate a key based on the key
// type passed in.
//
// the key type is determined via the keytype enum.
func autogenKey(keytype int) (key []byte, err error) {
	switch keytype {
	case AES_128:
		return ofscommon.GenerateRandomBytes(consts.AES_128_KEYLENGTH, consts.AES_128_KEYLENGTH+1)
	case AES_192:
		return ofscommon.GenerateRandomBytes(consts.AES_192_KEYLENGTH, consts.AES_192_KEYLENGTH+1)
	case AES_256:
		return ofscommon.GenerateRandomBytes(consts.AES_256_KEYLENGTH, consts.AES_256_KEYLENGTH+1)
	default:
		return nil, fmt.Errorf(encmsg.ERR_KEYTYP)
	}
}
