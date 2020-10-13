package chain

import "github.com/btcsuite/btcd/chaincfg"

type ChainParams struct {
	PubKeyHashAddrID byte
	ScriptAddress    byte
	SecretKey        byte
}

var Testnet = ChainParams{
	PubKeyHashAddrID: byte(65),
	ScriptAddress:    byte(127),
	SecretKey:        byte(130),
}

func (c *ChainParams) params() *chaincfg.Params {
	return &chaincfg.Params{
		PubKeyHashAddrID: c.PubKeyHashAddrID,
		ScriptHashAddrID: c.ScriptAddress,
		PrivateKeyID:     c.SecretKey,
	}
}
