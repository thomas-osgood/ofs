package ofsaesencryptor

import (
	ofscommon "github.com/thomas-osgood/ofs/internal/general"
)

// function designed to decrypt a file using AES encryption.
func (ae *AESEncryptor) DecryptFile(filename string) (err error) {
	var ciphertext []byte
	var plaintext []byte

	ciphertext, err = ofscommon.ReadFileBytes(filename)
	if err != nil {
		return err
	}

	plaintext, err = ae.decryptBytesAES(ciphertext)
	if err != nil {
		return err
	}

	return ofscommon.WriteFileBytes(filename, plaintext)
}

// function designed to encrypt a file using AES encryption.
func (ae *AESEncryptor) EncryptFile(filename string) (err error) {
	var ciphertext []byte
	var plaintext []byte

	plaintext, err = ofscommon.ReadFileBytes(filename)
	if err != nil {
		return err
	}

	ciphertext, err = ae.encryptBytesAES(plaintext)
	if err != nil {
		return err
	}

	return ofscommon.WriteFileBytes(filename, ciphertext)
}
