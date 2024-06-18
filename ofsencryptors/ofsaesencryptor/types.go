package ofsaesencryptor

type AESEncryptor struct {
	key []byte
}

type AESEncryptorOpt struct {
	Key []byte
}

type AESEncryptorOptFunc func(*AESEncryptorOpt) error
