package ofsrsaencryptor

import (
	"fmt"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	consts "github.com/thomas-osgood/ofs/ofsencryptors/internal/constants"
	ofsmessages "github.com/thomas-osgood/ofs/ofsencryptors/internal/messages"
)

// function designed to decrypt a file's contents.
func (rsae *RSAEncryptor) decryptBytesRSA(ciphertext []byte) (err error) {
	return nil
}

// function designed to encrypt a file's contents.
func (rsae *RSAEncryptor) encryptBytesRSA(plaintext []byte) (err error) {
	return nil
}

// func designed to manipulate the data contained within
// a given file. this will either encrypt or decrypt the
// data and write it back to the file.
func (rsae *RSAEncryptor) manipulateFileData(filename string, action int) (err error) {
	var content []byte

	content, err = ofscommon.ReadFileBytes(filename)
	if err != nil {
		return err
	}

	switch action {
	case consts.ACT_DECRYPT:
		return rsae.decryptBytesRSA(content)
	case consts.ACT_ENCRYPT:
		return rsae.encryptBytesRSA(content)
	default:
		return fmt.Errorf(ofsmessages.ERR_ACTION_UNKNOWN)
	}
}
