package ofsbcrypt

import (
	bcryptconsts "github.com/thomas-osgood/ofs/ofsauthenticators/hashers/ofsbcrypt/internal/constants"
)

func NewBCryptHasher() (hasher *BCryptHasher, err error) {
	var defaults BCryptHasherOption = BCryptHasherOption{
		Cost: bcryptconsts.DEFAULT_COST,
	}

	hasher = new(BCryptHasher)
	hasher.cost = defaults.Cost

	return hasher, nil
}
