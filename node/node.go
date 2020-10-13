package node

import "github.com/btcsuite/btcutil"

type Node interface {
	DecodeAddress(address string) (btcutil.Address, error)
	PayToAddress(address btcutil.Address, amount btcutil.Amount) (string, error)
}

type RPCNode struct {
	URL         string
	Username    string
	Password    string
	ChainParams Chain
}

func New(url, username, password string, chain Chain) (*RPCNode, error) {
	n := RPCNode{
		URL:         url,
		Username:    username,
		Password:    password,
		ChainParams: chain,
	}
	return &n, nil
}

func (n *RPCNode) DecodeAddress(address string) (btcutil.Address, error) {
	return btcutil.DecodeAddress(address, n.ChainParams.params())
}

func (n *RPCNode) PayToAddress(address btcutil.Address, amount btcutil.Amount) (string, error) {
	return "", nil
}
