package ofsaesencryptor

import (
	encryptorvalidations "github.com/thomas-osgood/ofs/ofsencryptors/ofsaesencryptor/internal/validations"
)

// function designed to create, initialize and return a new
// aes encryptor object.
//
// if WithKey() is not specified, an encryption key will be
// randomly generated upon initialization and AES-256 will
// be implemented.
func NewAesEncryptor(opts ...AESEncryptorOptFunc) (encryptor *AESEncryptor, err error) {
	var defaults AESEncryptorOpt = AESEncryptorOpt{Key: make([]byte, 0)}
	var curopt AESEncryptorOptFunc

	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	encryptor = new(AESEncryptor)

	encryptor.key = defaults.Key

	return encryptor, nil
}

// specify a key to use for the AES ecnryptor.
func WithKey(key []byte) AESEncryptorOptFunc {
	return func(ao *AESEncryptorOpt) (err error) {

		err = encryptorvalidations.ValidateKey(key)
		if err != nil {
			return err
		}

		ao.Key = key

		return nil
	}
}
