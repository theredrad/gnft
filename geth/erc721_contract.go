package geth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/theredrad/gnft"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ERC721Contract struct {
	contract *bind.BoundContract
}

func NewERC721Contract(c *GETH, address string) (*ERC721Contract, error) {
	a, err := abi.JSON(strings.NewReader(gnft.ABIERC721))
	if err != nil {
		return nil, err
	}

	return &ERC721Contract{
		contract: bind.NewBoundContract(common.HexToAddress(address), a, c.client, c.client, c.client),
	}, nil
}

func (c *ERC721Contract) BalanceOf(address string) (*big.Int, error) {
	var res []interface{}

	add := common.HexToAddress(address)
	err := c.contract.Call(&bind.CallOpts{}, &res, "balanceOf", add)
	if err != nil {
		return nil, err
	}

	b, err := gnft.ToBigInt(res)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *ERC721Contract) OwnerOf(tokenID *big.Int) (string, error) {
	var res []interface{}
	err := c.contract.Call(&bind.CallOpts{}, &res, "ownerOf", tokenID)
	if err != nil {
		return "", err
	}

	adr, err := addressFromResult(res)
	if err != nil {
		return "", err
	}
	return adr, nil
}

func (c *ERC721Contract) SafeTransferFromWithData(from, to string, tokenID *big.Int, data []byte) error {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) SafeTransferFrom(from, to string, tokenID *big.Int, data []byte) error {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) TransferFrom(from, to string, tokenID *big.Int) error {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) Approved(approved string, tokenID *big.Int) error {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) SetApprovalForAll(operator string, approved bool) error {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) GetApproved(tokenID *big.Int) (string, error) {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) IsApprovedForAll(owner, operator string) (bool, error) {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) Name() (string, error) {
	var res []interface{}
	err := c.contract.Call(&bind.CallOpts{}, &res, "name")
	if err != nil {
		return "", err
	}

	s, err := gnft.ToString(res)
	if err != nil {
		return "", err
	}
	return s, nil
}

func (c *ERC721Contract) Symbol() (string, error) {
	var res []interface{}
	err := c.contract.Call(&bind.CallOpts{}, &res, "symbol")
	if err != nil {
		return "", err
	}

	s, err := gnft.ToString(res)
	if err != nil {
		return "", err
	}
	return s, nil
}

func (c *ERC721Contract) TokenURI(tokenID *big.Int) (*gnft.TokenURI, error) {
	var res []interface{}
	err := c.contract.Call(&bind.CallOpts{}, &res, "tokenURI", tokenID)
	if err != nil {
		return nil, err
	}

	uri, err := stringFromResult(res)
	if err != nil {
		return nil, err
	}

	b, err := fetchTokenURI(uri)
	if err != nil {
		return nil, err
	}

	t := gnft.TokenURI{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *ERC721Contract) TotalSupply() (*big.Int, error) {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) TokenByIndex(index *big.Int) (*big.Int, error) {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) TokenOfOwnerByIndex(owner string, index *big.Int) (*big.Int, error) {
	panic("not implemented") // TODO
}

func (c *ERC721Contract) OnERC721Received(operator, from string, tokenId *big.Int, data []byte) ([4]byte, error) {
	panic("not implemented") // TODO
}

func stringFromResult(result []interface{}) (string, error) {
	if len(result) == 0 {
		return "", gnft.ErrNoResult
	}
	if adr, ok := result[0].(string); ok {
		return adr, nil
	}
	return "", gnft.ErrInvalidResult
}

func addressFromResult(result []interface{}) (string, error) {
	if len(result) == 0 {
		return "", gnft.ErrNoResult
	}
	if adr, ok := result[0].(common.Address); ok {
		return adr.Hex(), nil
	}
	return "", gnft.ErrInvalidResult
}

func fetchTokenURI(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("http response code: %d", resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
