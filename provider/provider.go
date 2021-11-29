package provider

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/theredrad/gnft/standard"
)

type Provider interface {
	NewContract(address common.Address) standard.NFT
}
