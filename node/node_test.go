package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressDecoding(t *testing.T) {
	asserts := assert.New(t)
	addresses := [...]string{
		"TJ8Yi3uxJZb8Wjvt5gVuoJcYHZi9FyVKKA",
		"TSWyjnRGHobTPaGfafrT1Q1FkqUU8dumYW",
	}

	for _, address := range addresses {
		n, err := New("", "", "", Testnet)

		_, err = n.DecodeAddress(address)

		asserts.NoError(err)
	}
}
