package ofsnohash

// function designed to create, initialize and return
// a point to an instance of a new NoHasher object.
func NewNoHasher() (hasher *NoHasher, err error) {
	return &NoHasher{}, nil
}
