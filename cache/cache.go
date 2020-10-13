package cache

import "github.com/btcsuite/btcutil"

type Cache interface {
	CanAddAddress(addr btcutil.Address) (bool, error)
	AddAddressToQueue(addr btcutil.Address) error
	AddAddressPayout() error
	Close()
}
