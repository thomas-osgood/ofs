package ofsaesencryptor

import (
	"fmt"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	consts "github.com/thomas-osgood/ofs/ofsencryptors/internal/constants"
	encmessages "github.com/thomas-osgood/ofs/ofsencryptors/internal/messages"
)

// function designed to encrypt or decrypt a file based on the
// action specified by the user.
//
// the actions are defined in an enum within the internal/constants
// module of the OFSAESEncryptor.
func (ae *AESEncryptor) manipulateFileData(filename string, action int) (err error) {
	var original []byte
	var output []byte

	original, err = ofscommon.ReadFileBytes(filename)
	if err != nil {
		return err
	}

	switch action {
	case consts.ACT_DECRYPT:
		output, err = ae.DecryptBytes(original)
	case consts.ACT_ENCRYPT:
		output, err = ae.EncryptBytes(original)
	default:
		err = fmt.Errorf(encmessages.ERR_ACTION_UNKNOWN)
	}

	if err != nil {
		return err
	}

	return ofscommon.WriteFileBytes(filename, output)
}
