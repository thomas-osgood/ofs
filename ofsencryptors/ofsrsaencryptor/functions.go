package ofsrsaencryptor

import "crypto/rsa"

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

	err = defaults.detectKey()
	if err != nil {
		return nil, err
	}

	encryptor = new(RSAEncryptor)

	encryptor.privkeybytes = defaults.PrivkeyBytes
	encryptor.pubkeybytes = defaults.PubkeyBytes

	return encryptor, nil
}

// auto-generate a public-private key pair to be used by
// the encryptor.
func WithRSAKeyAuto() RSAEncryptorOptFunc {
	return func(ro *RSAEncryptorOpt) (err error) {
		ro.PrivkeyBytes, ro.PubkeyBytes, err = genKeyPair()
		return err
	}
}

// set the public-private key pair used by the encryptor by
// passing in an RSA Private Key object.
func WithRSAKeyPair(privkey *rsa.PrivateKey) RSAEncryptorOptFunc {
	return func(ro *RSAEncryptorOpt) error {
		ro.PrivkeyBytes, ro.PubkeyBytes = genPubPrivBytes(privkey)
		return nil
	}
}
