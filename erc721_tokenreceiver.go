package gnft

import (
	"math/big"
)

type ERC721TokenReceiver interface {
	OnERC721Received(operator, from string, tokenId *big.Int, data []byte) ([4]byte, error)
}
