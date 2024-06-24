package ofsbcrypt

import (
	"fmt"

	bcryptconsts "github.com/thomas-osgood/ofs/ofsauthenticators/hashers/ofsbcrypt/internal/constants"
	bcryptmsg "github.com/thomas-osgood/ofs/ofsauthenticators/hashers/ofsbcrypt/internal/messages"
	"golang.org/x/crypto/bcrypt"
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

// set the hash cost.
func WithCost(cost int) BCryptHasherOptFunc {
	return func(bho *BCryptHasherOption) error {
		if (cost < bcrypt.MinCost) || (cost > bcrypt.MaxCost) {
			return fmt.Errorf(bcryptmsg.ERR_INVALID_COST, bcrypt.MinCost, bcrypt.MaxCost)
		}
		bho.Cost = cost
		return nil
	}
}
