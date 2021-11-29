package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/theredrad/gnft/standard"
	"github.com/theredrad/gnft/types"
	"io"
	"math/big"
	"net/http"
	"strings"
)

var (
	ErrInvalidResult = errors.New("invalid result")
	ErrNoResult      = errors.New("no result")
)

type GETH struct {
	client *ethclient.Client
}

type Contract struct {
	contract *bind.BoundContract
}

func NewGETH(rawurl string) (*GETH, error) {
	c, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return &GETH{
		client: c,
	}, nil
}

func (c *GETH) NewContract(address common.Address) (*Contract, error) {
	a, err := abi.JSON(strings.NewReader(standard.ABIERC721))
	if err != nil {
		return nil, err
	}

	return &Contract{
		contract: bind.NewBoundContract(address, a, c.client, c.client, c.client),
	}, nil
}

func (c *Contract) BalanceOf(address common.Address) (*big.Int, error) {
	panic("not implemented") // TODO
}

func (c *Contract) OwnerOf(tokenID *big.Int) (common.Address, error) {
	var res []interface{}
	err := c.contract.Call(&bind.CallOpts{}, &res, "ownerOf", tokenID)
	if err != nil {
		return common.Address{}, err
	}

	adr, err := addressFromResult(res)
	if err != nil {
		return common.Address{}, err
	}
	return adr, nil
}

func (c *Contract) SafeTransferFromWithData(from common.Address, to common.Address, tokenID *big.Int, data []byte) error {
	panic("not implemented") // TODO
}

func (c *Contract) SafeTransferFrom(from common.Address, to common.Address, tokenID *big.Int, data []byte) error {
	panic("not implemented") // TODO
}

func (c *Contract) TransferFrom(from common.Address, to common.Address, tokenID *big.Int) error {
	panic("not implemented") // TODO
}

func (c *Contract) Approved(approved common.Address, tokenID *big.Int) error {
	panic("not implemented") // TODO
}
func (c *Contract) SetApprovalForAll(operator common.Address, approved bool) error {
	panic("not implemented") // TODO
}

func (c *Contract) GetApproved(tokenID *big.Int) (common.Address, error) {
	panic("not implemented") // TODO
}

func (c *Contract) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	panic("not implemented") // TODO
}

func (c *Contract) Name() (string, error) {
	panic("not implemented") // TODO
}

func (c *Contract) Symbol() (string, error) {
	panic("not implemented") // TODO
}

func (c *Contract) TokenURI(tokenID *big.Int) (*types.TokenURI, error) {
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

	t := types.TokenURI{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Contract) TotalSupply() (*big.Int, error) {
	panic("not implemented") // TODO
}

func (c *Contract) TokenByIndex(index *big.Int) (*big.Int, error) {
	panic("not implemented") // TODO
}

func (c *Contract) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	panic("not implemented") // TODO
}

func (c *Contract) OnERC721Received(operator, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	panic("not implemented") // TODO
}

func stringFromResult(result []interface{}) (string, error) {
	if len(result) == 0 {
		return "", ErrNoResult
	}
	if adr, ok := result[0].(string); ok {
		return adr, nil
	}
	return "", ErrInvalidResult
}

func addressFromResult(result []interface{}) (common.Address, error) {
	if len(result) == 0 {
		return common.Address{}, ErrNoResult
	}
	if adr, ok := result[0].(common.Address); ok {
		return adr, nil
	}
	return common.Address{}, ErrInvalidResult
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
