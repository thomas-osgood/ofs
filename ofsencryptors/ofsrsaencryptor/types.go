package ofsrsaencryptor

type RSAEncryptor struct {
	privkeybytes []byte
	pubkeybytes  []byte
}

type RSAEncryptorOpt struct {
	PrivkeyBytes []byte
	PubkeyBytes  []byte
}

type RSAEncryptorOptFunc func(*RSAEncryptorOpt) error
