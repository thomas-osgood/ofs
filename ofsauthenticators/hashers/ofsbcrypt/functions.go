package ofsbcrypt

import (
	bcryptconsts "github.com/thomas-osgood/ofs/ofsauthenticators/hashers/ofsbcrypt/internal/constants"
)

// function designed to create, initialize and return a BCryptHasher object.
func NewBCryptHasher(opts ...BCryptHasherOptFunc) (hasher *BCryptHasher, err error) {
	var curopt BCryptHasherOptFunc
	var defaults BCryptHasherOption = BCryptHasherOption{
		Cost: bcryptconsts.DEFAULT_COST,
	}

	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	hasher = new(BCryptHasher)
	hasher.cost = defaults.Cost

	return hasher, nil
}
