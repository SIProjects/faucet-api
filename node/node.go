package node

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
)

type Node interface {
	PayToAddresses(xs []Payment) (string, map[btcutil.Address]btcutil.Amount, error)
	GetTransaction(txid string) (*btcjson.GetTransactionResult, error)
	GetBalance() (btcutil.Amount, error)
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

func (n *RPCNode) PayToAddresses(
	xs []Payment,
) (string, map[btcutil.Address]btcutil.Amount, error) {

	amounts := make(map[btcutil.Address]btcutil.Amount)

	for _, x := range xs {
		amounts[x.Address] = x.Amount
	}

	hash, err := n.Client.SendMany("", amounts)

	if err != nil {
		return "", amounts, err
	}

	return hash.String(), amounts, nil
}

func (n *RPCNode) GetTransaction(
	txid string,
) (*btcjson.GetTransactionResult, error) {
	hash, err := chainhash.NewHashFromStr(txid)

	if err != nil {
		return nil, err
	}

	return n.Client.GetTransaction(hash)
}

func (n *RPCNode) GetBalance() (btcutil.Amount, error) {
	return n.Client.GetBalance("*")
}
