package gnft

import (
	"math/big"
)

type ERC721Metadata interface {
	Name() (string, error)
	Symbol() (string, error)
	TokenURI(tokenID *big.Int) (*TokenURI, error)
}
