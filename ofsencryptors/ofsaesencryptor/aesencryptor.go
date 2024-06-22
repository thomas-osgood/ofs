// module defining an AES Encryptor object. this encryptor
// is able to be used with OFS (or as a stand-alone object)
// to encrypt and decrypt files.
package ofsaesencryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	consts "github.com/thomas-osgood/ofs/ofsencryptors/internal/constants"
	aesmessages "github.com/thomas-osgood/ofs/ofsencryptors/ofsaesencryptor/internal/messages"
)

// function designed to decrypt bytes and return the "plainbytes".
func (ae *AESEncryptor) DecryptBytes(ciphertext []byte) (plaintext []byte, err error) {
	var aesblock cipher.Block
	var gcm cipher.AEAD
	var nonce []byte
	var nonceSize int

	if aesblock, err = aes.NewCipher(ae.key); err != nil {
		return nil, err
	}

	if gcm, err = cipher.NewGCM(aesblock); err != nil {
		return nil, err
	}

	nonceSize = gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf(aesmessages.ERR_CIPHERTEXT_NONCE_SIZE)
	}

	nonce, ciphertext = ciphertext[:nonceSize], ciphertext[nonceSize:]

	return gcm.Open(nil, nonce, ciphertext, nil)
}

// function designed to decrypt a file using AES encryption.
func (ae *AESEncryptor) DecryptFile(filename string) (err error) {
	return ae.manipulateFileData(filename, consts.ACT_DECRYPT)
}

// function designed to encrypt bytes and return the "cipherbytes".
func (ae *AESEncryptor) EncryptBytes(plaintext []byte) (ciphertext []byte, err error) {
	var aesblock cipher.Block
	var gcm cipher.AEAD
	var nonce []byte

	if aesblock, err = aes.NewCipher(ae.key); err != nil {
		return nil, err
	}

	if gcm, err = cipher.NewGCM(aesblock); err != nil {
		return nil, err
	}
	nonce = make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// function designed to encrypt a file using AES encryption.
func (ae *AESEncryptor) EncryptFile(filename string) (err error) {
	return ae.manipulateFileData(filename, consts.ACT_ENCRYPT)
}
