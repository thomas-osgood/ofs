// module defining an RSA Encryptor object. this encryptor
// is able to be used with OFS (or as a stand-alone object)
// to encrypt and decrypt files.
//
// note: this encryptor should only be used if the files that
// will be saved, served and encrypted are small. the max file
// size that can be encrypted is ((keysize / 8) - headersize).
// for a key of 4096 bits using PKCS1 that means ((4096 / 8) - 11)
// which comes out to a max file size of 501 bytes. for files with
// sizes larger than this max size, AES (or another encryption type)
// is recommended.
package ofsrsaencryptor

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	consts "github.com/thomas-osgood/ofs/ofsencryptors/internal/constants"
)

// function designed to decrypt provided ciphertext bytes.
func (rsae *RSAEncryptor) DecryptBytes(ciphertext []byte) (plaintext []byte, err error) {
	var privkey *rsa.PrivateKey

	if privkey, err = rsae.constructPrivKey(); err != nil {
		return nil, err
	}

	plaintext, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, privkey, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// function designed to read a given file's data and
// decrypt the file.
func (rsae *RSAEncryptor) DecryptFile(filename string) (err error) {
	return rsae.manipulateFileData(filename, consts.ACT_DECRYPT)
}

// function designed to encrypt provided plaintext bytes.
func (rsae *RSAEncryptor) EncryptBytes(plaintext []byte) (ciphertext []byte, err error) {
	var pubkey *rsa.PublicKey

	if pubkey, err = rsae.constructPublicKey(); err != nil {
		return nil, err
	}

	ciphertext, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, pubkey, plaintext, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

// function designed to read a given file's data and
// encrypt the file.
func (rsae *RSAEncryptor) EncryptFile(filename string) (err error) {
	return rsae.manipulateFileData(filename, consts.ACT_ENCRYPT)
}
