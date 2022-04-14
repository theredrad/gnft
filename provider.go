package gnft

import (
	"github.com/ethereum/go-ethereum/common"
)

type Provider interface {
	NewERC721Contract(address common.Address) NFT721
}
