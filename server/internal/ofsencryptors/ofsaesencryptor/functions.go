package ofsaesencryptor

// function designed to create, initialize and return a new
// aes encryptor object.
func NewAesEncryptor() (encryptor *AESEncryptor, err error) {

	encryptor = new(AESEncryptor)

	return encryptor, nil
}
