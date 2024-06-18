package server

// interface defining an object that can perform various
// encryption methods.
//
// this can be used to encrypt and decrypt files on the server.
type OFSEncryptor interface {
	// decrypt a given encrypted file.
	DecryptFile(string) error
	// encrypt a given file.
	EncryptFile(string) error
}
