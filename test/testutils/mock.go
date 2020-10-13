package testutils

import "github.com/btcsuite/btcutil"

type MockCache struct {
}

func (c *MockCache) CanAddAddress(addr btcutil.Address) (bool, error) {
	return true, nil
}
func (c *MockCache) AddAddressToQueue(addr btcutil.Address) error {
	return nil
}
func (c *MockCache) AddAddressPayout() error {
	return nil
}
func (c *MockCache) Close() {
}

func NewMockCache() *MockCache {
	return &MockCache{}
}
