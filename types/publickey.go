package types

import (
	"cosmossdk/protoTx/proto"
	"cosmossdk/tendermint"
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