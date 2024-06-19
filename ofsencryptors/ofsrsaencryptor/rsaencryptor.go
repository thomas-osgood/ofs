// module defining an RSA Encryptor object. this encryptor
// is able to be used with OFS (or as a stand-alone object)
// to encrypt and decrypt files.
package ofsrsaencryptor

import (
	consts "github.com/thomas-osgood/ofs/ofsencryptors/internal/constants"
)

// function designed to read a given file's data and
// decrypt the file.
func (rsae *RSAEncryptor) DecryptFile(filename string) (err error) {
	return rsae.manipulateFileData(filename, consts.ACT_DECRYPT)
}

// function designed to read a given file's data and
// encrypt the file.
func (rsae *RSAEncryptor) EncryptFile(filename string) (err error) {
	return rsae.manipulateFileData(filename, consts.ACT_ENCRYPT)
}