package general

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	encmsg "github.com/thomas-osgood/ofs/internal/messages"
)

func GenerateRandomBytes(minlen int, maxlen int) (randbytes []byte, err error) {
	var baseval *big.Int
	var biglen *big.Int
	var bigmin *big.Int
	var charset []byte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var length int
	var randidx *big.Int

	// validate min/max parameters
	if (minlen <= 0) || (maxlen <= 0) {
		return nil, errors.New("min and max lengths mus be greater than zero")
	} else if minlen > maxlen {
		return nil, errors.New("min length must be less than or equal to max length")
	}

	// this is the number that will be used to generate
	// the random number. this is the difference of the
	// max value and min value because the final random
	// number will be calculated by adding the min value
	// so the number falls within the range MIN <= x <= MAX.
	baseval = big.NewInt(int64(maxlen - minlen))

	// convert the minimum value to a big.Int so it can be
	// used to adjust the randomly generated length.
	bigmin = big.NewInt(int64(minlen))

	// use the crypto/rand library to generate a length
	// for the string.
	biglen, err = rand.Int(rand.Reader, big.NewInt(baseval.Int64()))
	if err != nil {
		return nil, fmt.Errorf(encmsg.ERR_RANDLEN_GEN, err)
	}

	// adjust the generated number to fit within the range.
	biglen = biglen.Add(biglen, bigmin)

	length = int(biglen.Int64())

	randbytes = make([]byte, 0)

	for i := 0; i < length; i++ {
		// calculate the random index to choose. if
		// there is an error, choose index 0.
		randidx, err = rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			randidx = big.NewInt(0)
		}
		// append the char at the randomly generated
		// index to the randomly generated string.
		randbytes = append(randbytes, charset[randidx.Int64()])
	}

	return randbytes, nil
}
