package node

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
)

type Node interface {
	PayToAddresses(xs []Payment) (string, error)
}

type RPCNode struct {
	URL      string
	Username string
	Password string
	Client   *rpcclient.Client
}

func New(url, username, password string) (*RPCNode, error) {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         url,
		User:         username,
		Pass:         password,
	}, nil)

	if err != nil {
		return nil, err
	}

	n := RPCNode{
		URL:      url,
		Username: username,
		Password: password,
		Client:   client,
	}
	return &n, nil
}

type Payment struct {
	Address btcutil.Address
	Amount  btcutil.Amount
}

func (n *RPCNode) PayToAddresses(xs []Payment) (string, error) {

	amounts := make(map[btcutil.Address]btcutil.Amount)

	for _, x := range xs {
		amounts[x.Address] = x.Amount
	}

	hash, err := n.Client.SendMany("", amounts)

	if err != nil {
		return "", err
	}

	return hash.String(), nil
}
