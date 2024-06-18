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
