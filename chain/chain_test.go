package chain_test

import (
	"testing"

	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/node"
	"github.com/stretchr/testify/assert"
)

func TestAddressDecoding(t *testing.T) {
	asserts := assert.New(t)
	addresses := [...]string{
		"TJ8Yi3uxJZb8Wjvt5gVuoJcYHZi9FyVKKA",
		"TSWyjnRGHobTPaGfafrT1Q1FkqUU8dumYW",
	}

	for _, address := range addresses {
		n, err := node.New("", "", "")
		ch := chain.New(n, chain.Testnet)

		_, err = ch.DecodeAddress(address)

		asserts.NoError(err)
	}
}
