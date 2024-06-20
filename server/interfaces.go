package server

// interface defining an object that can perform various
// encryption methods on files.
//
// this can be used to encrypt and decrypt files on the server.
type OFSEncryptor interface {
	// decrypt a given encrypted file.
	DecryptFile(string) error
	// encrypt a given file.
	EncryptFile(string) error
}

// interface defining an object that can perform vaious encryption
// methods on bytes.
type BytesEncryptor interface {
	// decrypt provided ciphertext and return the
	// resulting plaintext.
	DecryptBytes([]byte) ([]byte, error)
	// encrypt provided plaintext and return the
	// resulting ciphertext.
	EncryptBytes([]byte) ([]byte, error)
}
