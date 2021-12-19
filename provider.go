package gnft

import (
	"github.com/ethereum/go-ethereum/common"
)

type Provider interface {
	NewContract(address common.Address) NFT
}
