package ofsrsaencryptor

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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
	var maxsize int

	maxsize, err = rsae.maxEncryptionSize()
	if err != nil {
		return err
	}

	content, err = ofscommon.ReadFileBytes(filename)
	if err != nil {
		return err
	}

	// if the file content length exceeds the maximum length
	// the current RSA object can encrypt, return an error.
	if len(content) > maxsize {
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
//
// references:
//
// https://serverfault.com/questions/325467/i-have-a-keypair-how-do-i-determine-the-key-length
//
// https://stackoverflow.com/questions/42707353/how-to-verify-rsa-key-length-in-go
func (rsae *RSAEncryptor) maxEncryptionSize() (maxsize int, err error) {
	return maxsize, nil
}

// function designed to build and return an RSA private key object
// using the private key bytes saved by the RSAEncryptor.
func (rsae *RSAEncryptor) constructPrivKey() (key *rsa.PrivateKey, err error) {
	var der *pem.Block

	if (rsae.privkeybytes == nil) || (len(rsae.privkeybytes) < 1) {
		return nil, fmt.Errorf(rsamessages.ERR_PRIVKEY_GEN)
	}

	der, _ = pem.Decode(rsae.privkeybytes)
	if der == nil {
		return nil, fmt.Errorf(rsamessages.ERR_PRIVKEY_DEC)
	}

	key, err = x509.ParsePKCS1PrivateKey(der.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
