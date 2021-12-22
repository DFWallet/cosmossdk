package types

import (
	"github.com/DFWallet/cosmossdk/protoTx/proto"
	"github.com/DFWallet/cosmossdk/tendermint"
)

// PubKey defines a public key and extends proto.Message.
type PubKey interface {
	proto.Message

	Address() Address
	Bytes() []byte
	VerifySignature(msg []byte, sig []byte) bool
	Equals(PubKey) bool
	Type() string
}

type (
	Address = tendermint.Address
)