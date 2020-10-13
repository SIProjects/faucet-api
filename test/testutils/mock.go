package testutils

import "github.com/btcsuite/btcutil"

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

func NewMockCache() *MockCache {
	return &MockCache{
		Pending: make(map[string]struct{}, 0),
	}
}

type MockNode struct {
}

func (n *MockNode) PayToAddress(address btcutil.Address, amount btcutil.Amount) (string, error) {
	return "", nil
}

func NewMockNode() *MockNode {
	return &MockNode{}
}
