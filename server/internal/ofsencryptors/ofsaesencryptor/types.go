package ofsaesencryptor

type AESEncryptor struct {
	key []byte

	// private key for AES encryption.
	privkey []byte
	// public key for AES encryption.
	pubkey []byte
}

type AESEncryptorOpt struct {
	Key []byte
}

type AESEncryptorOptFunc func(*AESEncryptorOpt) error
