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

	priv, err = rsa.GenerateKey(rand.Reader, consts.KEYSIZE)
	if err != nil {
		return nil, nil, err
	}

	privkey, pubkey = genPubPrivBytes(priv)

	return privkey, pubkey, nil
}

// function designed to generate the public and private key bytes
// for a key pair based on a given RSA private key. this utilizes
// the MarshalPKCS1 functions to generate the private and public
// key pair bytes.
func genPubPrivBytes(priv *rsa.PrivateKey) (privkey []byte, pubkey []byte) {
	var pub *rsa.PublicKey = new(rsa.PublicKey)

	privkey = pem.EncodeToMemory(&pem.Block{
		Type:  consts.RSA_TYPE_PRIVKEY,
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})

	pubkey = pem.EncodeToMemory(&pem.Block{
		Type:  consts.RSA_TYPE_PUBKEY,
		Bytes: x509.MarshalPKCS1PublicKey(pub),
	})

	return privkey, pubkey
}
