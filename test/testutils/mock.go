package testutils

import (
	"github.com/SIProjects/faucet-api/node"
	"github.com/SIProjects/faucet-api/test/fixture"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcutil"
)

type MockCache struct {
	Pending map[string]struct{}
	Fixture *fixture.CacheFixtures
}

func (c *MockCache) CanAddAddress(addr btcutil.Address) (bool, error) {
	return true, nil
}

func (c *MockCache) AddAddressToQueue(addr btcutil.Address) error {
	c.Pending[addr.String()] = struct{}{}
	return nil
}

func (c *MockCache) AddAddressPayout() error {
	return nil
}

func (c *MockCache) Close() {
}

func (c *MockCache) GetQueuedCount() (res int64, err error) {
	if c.Fixture != nil {
		res = int64(c.Fixture.GetQueuedCount)
	}

	return res, err
}

func (c *MockCache) GetNextAddresses(num int) ([]btcutil.Address, error) {
	return []btcutil.Address{}, nil
}

func NewMockCache(fx *fixture.CacheFixtures) *MockCache {
	return &MockCache{
		Pending: make(map[string]struct{}, 0),
		Fixture: fx,
	}
}

type MockNode struct {
	Fixture *fixture.NodeFixtures
}

func (n *MockNode) PayToAddresses(
	[]node.Payment,
) (string, map[btcutil.Address]btcutil.Amount, error) {
	var amounts map[btcutil.Address]btcutil.Amount
	return "", amounts, nil
}

func (n *MockNode) GetTransaction(
	hash string,
) (*btcjson.GetTransactionResult, error) {
	return nil, nil
}

func NewMockNode(fx *fixture.NodeFixtures) *MockNode {
	return &MockNode{
		Fixture: fx,
	}
}

func (n *MockNode) GetBalance() (res btcutil.Amount, err error) {
	if n.Fixture != nil {
		res, err = btcutil.NewAmount(n.Fixture.GetBalance)
	}
	return res, err
}

type MockScheduler struct {
}

func (n *MockScheduler) Start(done chan struct{}) {
}

func NewMockScheduler() *MockScheduler {
	return &MockScheduler{}
}
