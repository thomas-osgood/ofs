package ofsrsaencryptor

// function designed to create, initialize and return a new
// rsa encryptor object.
func NewRsaEncryptor(opts ...RSAEncryptorOptFunc) (encryptor *RSAEncryptor, err error) {
	var curopt RSAEncryptorOptFunc
	var defaults RSAEncryptorOpt = RSAEncryptorOpt{}

	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	encryptor = new(RSAEncryptor)

	encryptor.privkeypem = defaults.PrivkeyPem
	encryptor.pubkeypem = defaults.PubkeyPem

	return encryptor, nil
}
