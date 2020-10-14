package testutils

import (
	"github.com/SIProjects/faucet-api/node"
	"github.com/btcsuite/btcutil"
)

type MockCache struct {
	Pending map[string]struct{}
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

func (c *MockCache) GetNextAddresses(num int) ([]btcutil.Address, error) {
	return []btcutil.Address{}, nil
}

func NewMockCache() *MockCache {
	return &MockCache{
		Pending: make(map[string]struct{}, 0),
	}
}

type MockNode struct {
}

func (n *MockNode) PayToAddresses(
	[]node.Payment,
) (string, map[btcutil.Address]btcutil.Amount, error) {
	var amounts map[btcutil.Address]btcutil.Amount
	return "", amounts, nil
}

func NewMockNode() *MockNode {
	return &MockNode{}
}

type MockScheduler struct {
}

func (n *MockScheduler) Start(done chan struct{}) {
}

func NewMockScheduler() *MockScheduler {
	return &MockScheduler{}
}
