package ofsrsaencryptor

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	consts "github.com/thomas-osgood/ofs/ofsencryptors/ofsrsaencryptor/internal/constants"
)

// function designed to generate a private-public key pair
// for an RSA encryptor.
func genKeyPair() (privkey []byte, pubkey []byte, err error) {
	var priv *rsa.PrivateKey
	var pub *rsa.PublicKey = new(rsa.PublicKey)

	priv, err = rsa.GenerateKey(rand.Reader, consts.KEYSIZE)
	if err != nil {
		return nil, nil, err
	}

	privkey = pem.EncodeToMemory(&pem.Block{
		Type:  consts.RSA_TYPE_PRIVKEY,
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})

	pubkey = pem.EncodeToMemory(&pem.Block{
		Type:  consts.RSA_TYPE_PUBKEY,
		Bytes: x509.MarshalPKCS1PublicKey(pub),
	})

	return privkey, pubkey, nil
}
