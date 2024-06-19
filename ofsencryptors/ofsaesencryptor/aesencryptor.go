// module defining an AES Encryptor object. this encryptor
// is able to be used with OFS (or as a stand-alone object)
// to encrypt and decrypt files.
package ofsaesencryptor

import (
	consts "github.com/thomas-osgood/ofs/ofsencryptors/internal/constants"
)

// function designed to decrypt a file using AES encryption.
func (ae *AESEncryptor) DecryptFile(filename string) (err error) {
	return ae.manipulateFileData(filename, consts.ACT_DECRYPT)
}

// function designed to encrypt a file using AES encryption.
func (ae *AESEncryptor) EncryptFile(filename string) (err error) {
	return ae.manipulateFileData(filename, consts.ACT_ENCRYPT)
}
