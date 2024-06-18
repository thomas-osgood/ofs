package ofsrsaencryptor

type RSAEncryptor struct {
	privkeypem []byte
	pubkeypem  []byte
}

type RSAEncryptorOpt struct {
	PrivkeyPem []byte
	PubkeyPem  []byte
}

type RSAEncryptorOptFunc func(*RSAEncryptorOpt) error
