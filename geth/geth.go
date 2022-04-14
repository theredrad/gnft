package geth

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type GETH struct {
	client *ethclient.Client
}

func New(rawurl string) (*GETH, error) {
	c, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return &GETH{
		client: c,
	}, nil
}
