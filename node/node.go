package node

import "github.com/btcsuite/btcutil"

type Node interface {
	PayToAddress(address btcutil.Address, amount btcutil.Amount) (string, error)
}

type RPCNode struct {
	URL      string
	Username string
	Password string
}

func New(url, username, password string) (*RPCNode, error) {
	n := RPCNode{
		URL:      url,
		Username: username,
		Password: password,
	}
	return &n, nil
}

func (n *RPCNode) PayToAddress(address btcutil.Address, amount btcutil.Amount) (string, error) {
	return "", nil
}
