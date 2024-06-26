package ofssha256

func NewSha256Hasher() (hasher *SHA256Hasher, err error) {
	hasher = new(SHA256Hasher)
	return hasher, nil
}
