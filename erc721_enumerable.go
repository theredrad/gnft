package gnft

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ERC721Enumerable interface {
	TotalSupply() (*big.Int, error)
	TokenByIndex(index *big.Int) (*big.Int, error)
	TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error)
}
