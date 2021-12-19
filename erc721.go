package gnft

import (
	"math/big"
)

type ERC721 interface {
	BalanceOf(address string) (*big.Int, error)
	OwnerOf(tokenID *big.Int) (string, error)
	SafeTransferFromWithData(from string, to string, tokenID *big.Int, data []byte) error
	SafeTransferFrom(from string, to string, tokenID *big.Int, data []byte) error
	TransferFrom(from string, to string, tokenID *big.Int) error
	Transfer(from, to string) TransferObject
	Approved(approved string, tokenID *big.Int) error
	SetApprovalForAll(operator string, approved bool) error
	GetApproved(tokenID *big.Int) (string, error)
	IsApprovedForAll(owner string, operator string) (bool, error)
}

type TransferObject interface {
	Send() error
	SetFee() TransferObject
	DataBytes() []byte
}
