package standard

import (
	"github.com/theredrad/gnft/types"
	"math/big"
)

type ERC721Metadata interface {
	Name() (string, error)
	Symbol() (string, error)
	TokenURI(tokenID *big.Int) (*types.TokenURI, error)
}
