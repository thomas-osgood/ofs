package ofsbcrypt

import "golang.org/x/crypto/bcrypt"

// function designed to take in a password and hash and compare the two.
//
// if the password and hash match, nil will be returned.
func (bch *BCryptHasher) ComparePassHash(password []byte, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

// function designed to calculate and return the hash of a given password.
func (bch *BCryptHasher) HashPassword(password []byte) (hash []byte, err error) {
	return bcrypt.GenerateFromPassword(password, bch.cost)
}
