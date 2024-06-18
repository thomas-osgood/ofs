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

	// if no key is specified and WithAutokey has not been called,
	// auto-generate an AES-256 key.
	if (defaults.Key == nil) || (len(defaults.Key) < 1) {
		defaults.Key, err = autogenKey(AES_256)
		if err != nil {
			return nil, err
		}
	}

	encryptor = new(AESEncryptor)

	encryptor.key = defaults.Key

	return encryptor, nil
}

// auto-generate a key for the given AES type.
func WithAutokey(keytype int) AESEncryptorOptFunc {
	return func(ao *AESEncryptorOpt) (err error) {
		var key []byte

		key, err = autogenKey(keytype)
		if err != nil {
			return err
		}

		ao.Key = key

		return nil
	}
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
