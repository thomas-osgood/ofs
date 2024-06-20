package ofsrsaencryptor

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	ofscommon "github.com/thomas-osgood/ofs/internal/general"
	consts "github.com/thomas-osgood/ofs/ofsencryptors/internal/constants"
	encmessages "github.com/thomas-osgood/ofs/ofsencryptors/internal/messages"
	rsaconsts "github.com/thomas-osgood/ofs/ofsencryptors/ofsrsaencryptor/internal/constants"
	rsamessages "github.com/thomas-osgood/ofs/ofsencryptors/ofsrsaencryptor/internal/messages"
)

// function designed to decrypt a file's contents.
func (rsae *RSAEncryptor) decryptBytesRSA(ciphertext []byte) (plaintext []byte, err error) {
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

// function designed to encrypt a file's contents.
func (rsae *RSAEncryptor) encryptBytesRSA(plaintext []byte) (ciphertext []byte, err error) {
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

// func designed to manipulate the data contained within
// a given file. this will either encrypt or decrypt the
// data and write it back to the file.
func (rsae *RSAEncryptor) manipulateFileData(filename string, action int) (err error) {
	var content []byte
	var maxsize int
	var result []byte

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
		result, err = rsae.decryptBytesRSA(content)
	case consts.ACT_ENCRYPT:
		result, err = rsae.encryptBytesRSA(content)
	default:
		err = fmt.Errorf(encmessages.ERR_ACTION_UNKNOWN)
	}

	if err != nil {
		return err
	}

	return ofscommon.WriteFileBytes(filename, result)
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
	var key *rsa.PrivateKey

	key, err = rsae.constructPrivKey()
	if err != nil {
		return -1, err
	}

	maxsize = (key.N.BitLen() / 8) - rsaconsts.PKCS_HEADER_LEN

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

// function designed to build and return an RSA public key object using
// the public key bytes saved by the RSAEncryptor.
func (rsae *RSAEncryptor) constructPublicKey() (key *rsa.PublicKey, err error) {
	var der *pem.Block

	if (rsae.privkeybytes == nil) || (len(rsae.privkeybytes) < 1) {
		return nil, fmt.Errorf(rsamessages.ERR_PUBKEY_GEN)
	}

	der, _ = pem.Decode(rsae.pubkeybytes)
	if der == nil {
		return nil, fmt.Errorf(rsamessages.ERR_PUBKEY_DEC)
	}

	key, err = x509.ParsePKCS1PublicKey(der.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
