package tendermint

import (
	"cosmossdk/tendermint/bytes"
	"crypto/sha256"
)

type Address = bytes.HexBytes

func Sha256(bytes []byte) []byte {
	hasher := sha256.New()
	hasher.Write(bytes)
	return hasher.Sum(nil)
}
