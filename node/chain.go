package node

import "github.com/btcsuite/btcd/chaincfg"

type Chain struct {
	PubKeyHashAddrID byte
	ScriptAddress    byte
	SecretKey        byte
}

var Testnet = Chain{
	PubKeyHashAddrID: byte(65),
	ScriptAddress:    byte(127),
	SecretKey:        byte(130),
}

func (c *Chain) params() *chaincfg.Params {
	return &chaincfg.Params{
		PubKeyHashAddrID: c.PubKeyHashAddrID,
		ScriptHashAddrID: c.ScriptAddress,
		PrivateKeyID:     c.SecretKey,
	}
}
