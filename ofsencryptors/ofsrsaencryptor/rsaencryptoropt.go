package ofsrsaencryptor

// function designed to detect whether the key is empty.
// if it is, a new key will be autogenerated and assigned
// to RSAEncryptorOpt.
func (reo *RSAEncryptorOpt) detectKey() (err error) {

	// if no key option was provided, auto-generate a key.
	if (reo.PrivkeyBytes == nil) || (reo.PubkeyBytes == nil) {
		reo.PrivkeyBytes, reo.PubkeyBytes, err = genKeyPair()
	}

	return err
}
