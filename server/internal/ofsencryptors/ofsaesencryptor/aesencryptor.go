package ofsaesencryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
)

// function designed to decrypt a file using AES encryption.
func (ae *AESEncryptor) DecryptFile(filename string) (err error) {
	return nil
}

// function designed to encrypt a file using AES encryption.
func (ae *AESEncryptor) EncryptFile(filename string) (err error) {
	var aesblock cipher.Block
	var ciphertext []byte
	var gcm cipher.AEAD
	var nonce []byte
	var plaintext []byte

	plaintext, err = ofscommon.ReadFileBytes(filename)
	if err != nil {
		return err
	}

	if aesblock, err = aes.NewCipher(ae.key); err != nil {
		return err
	}

	if gcm, err = cipher.NewGCM(aesblock); err != nil {
		return err
	}
	nonce = make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext = gcm.Seal(nonce, nonce, plaintext, nil)

	return ofscommon.WriteFileBytes(filename, ciphertext)
}

// function designed to encrypt a file using AES encryption.
func (ae *AESEncryptor) GenerateKey() (key []byte, err error) {
	return key, nil
}
