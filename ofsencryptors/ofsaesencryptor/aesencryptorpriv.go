package ofsaesencryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	encryptormessages "github.com/thomas-osgood/ofs/ofsencryptors/internal/messages"
)

// function designed to decrypt bytes using a given AES key.
func (ae *AESEncryptor) decryptBytesAES(ciphertext []byte) (plaintext []byte, err error) {
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
		return nil, fmt.Errorf(encryptormessages.ERR_AES_CIPHERTEXT_NONCE_SIZE)
	}

	nonce, ciphertext = ciphertext[:nonceSize], ciphertext[nonceSize:]

	if plaintext, err = gcm.Open(nil, nonce, ciphertext, nil); err != nil {
		return nil, err
	}

	return plaintext, nil
}
