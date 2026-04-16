package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	eth        *ethclient.Client
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	from       common.Address
}

func NewClient(rpcURL, privateKeyHex, chainIDStr string) (*Client, error) {
	eth, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}

	chainID := new(big.Int)
	if _, ok := chainID.SetString(chainIDStr, 10); !ok {
		return nil, fmt.Errorf("failed to parse chain id: %s", chainIDStr)
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	pk, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	from := crypto.PubkeyToAddress(pk.PublicKey)

	return &Client{
		eth:        eth,
		privateKey: pk,
		chainID:    chainID,
		from:       from,
	}, nil
}

func (c *Client) NewTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.eth.PendingNonceAt(ctx, c.from)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(120))
	gasPrice = new(big.Int).Div(gasPrice, big.NewInt(100))

	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(500000)
	auth.Context = ctx

	return auth, nil
}

func (c *Client) Eth() *ethclient.Client {
	return c.eth
}

func (c *Client) From() common.Address {
	return c.from
}
