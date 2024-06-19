package ofsrsaencryptor

import (
	"fmt"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	consts "github.com/thomas-osgood/ofs/ofsencryptors/internal/constants"
	encmessages "github.com/thomas-osgood/ofs/ofsencryptors/internal/messages"
	rsamessages "github.com/thomas-osgood/ofs/ofsencryptors/ofsrsaencryptor/internal/messages"
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

	// if the file content length exceeds the maximum length
	// the current RSA object can encrypt, return an error.
	if len(content) > rsae.maxEncryptionSize() {
		return fmt.Errorf(rsamessages.ERR_CONTENT_TOO_LONG)
	}

	switch action {
	case consts.ACT_DECRYPT:
		return rsae.decryptBytesRSA(content)
	case consts.ACT_ENCRYPT:
		return rsae.encryptBytesRSA(content)
	default:
		return fmt.Errorf(encmessages.ERR_ACTION_UNKNOWN)
	}
}

// function designed to calculate and return the maximum size (in bytes)
// that the RSA object can encrypt.
func (rsae *RSAEncryptor) maxEncryptionSize() (maxsize int) {
	return maxsize
}
