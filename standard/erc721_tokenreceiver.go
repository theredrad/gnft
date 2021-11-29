package standard

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ERC721TokenReceiver interface {
	OnERC721Received(operator, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error)
}
