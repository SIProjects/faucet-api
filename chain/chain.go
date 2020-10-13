package chain

import (
	"github.com/SIProjects/faucet-api/node"
	"github.com/btcsuite/btcutil"
)

type Chain struct {
	Node   node.Node
	Params ChainParams
}

func New(n node.Node, p ChainParams) *Chain {
	return &Chain{
		Node:   n,
		Params: p,
	}
}

func (c *Chain) DecodeAddress(address string) (btcutil.Address, error) {
	return btcutil.DecodeAddress(address, c.Params.params())
}
