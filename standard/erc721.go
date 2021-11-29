package standard

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ERC721 interface {
	BalanceOf(address common.Address) (*big.Int, error)
	OwnerOf(tokenID *big.Int) (common.Address, error)
	SafeTransferFromWithData(from common.Address, to common.Address, tokenID *big.Int, data []byte) error
	SafeTransferFrom(from common.Address, to common.Address, tokenID *big.Int, data []byte) error
	TransferFrom(from common.Address, to common.Address, tokenID *big.Int) error
	Approved(approved common.Address, tokenID *big.Int) error
	SetApprovalForAll(operator common.Address, approved bool) error
	GetApproved(tokenID *big.Int) (common.Address, error)
	IsApprovedForAll(owner common.Address, operator common.Address) (bool, error)
}
