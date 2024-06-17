package ofsaesencryptor

// function designed to create, initialize and return a new
// aes encryptor object.
func NewAesEncryptor() (encryptor *AESEncryptor, err error) {

	encryptor = new(AESEncryptor)

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
