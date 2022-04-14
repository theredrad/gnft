package gnft

import (
	"math/big"
)

func ToBigInt(result []interface{}) (*big.Int, error) {
	if len(result) == 0 {
		return nil, ErrNoResult
	}
	if b, ok := result[0].(*big.Int); ok {
		return b, nil
	}

	return nil, ErrInvalidResult
}

func ToString(result []interface{}) (string, error) {
	if len(result) == 0 {
		return "", ErrNoResult
	}
	if s, ok := result[0].(string); ok {
		return s, nil
	}

	return "", ErrInvalidResult
}
