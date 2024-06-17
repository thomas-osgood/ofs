package ofsaesencryptor

// function designed to create, initialize and return a new
// aes encryptor object.
func NewAesEncryptor(opts ...AESEncryptorOptFunc) (encryptor *AESEncryptor, err error) {
	var defaults AESEncryptorOpt = AESEncryptorOpt{}
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

		err = validateKey(key)
		if err != nil {
			return err
		}

		ao.Key = key

		return nil
	}
}