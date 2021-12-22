package auth

import (
	tx "github.com/DFWallet/cosmossdk/txBody"
	cryptotypes "github.com/DFWallet/cosmossdk/types"
	"github.com/DFWallet/cosmossdk/types/signing"
	"fmt"
)

func SignatureDataToModeInfoAndSig(data signing.SignatureData) (*tx.ModeInfo, []byte) {
	if data == nil {
		return nil, nil
	}

	switch data := data.(type) {
	case *signing.SingleSignatureData:
		return &tx.ModeInfo{
			Sum: &tx.ModeInfo_Single_{
				Single: &tx.ModeInfo_Single{Mode: data.SignMode},
			},
		}, data.Signature
	case *signing.MultiSignatureData:
		n := len(data.Signatures)
		modeInfos := make([]*tx.ModeInfo, n)
		sigs := make([][]byte, n)

		for i, d := range data.Signatures {
			modeInfos[i], sigs[i] = SignatureDataToModeInfoAndSig(d)
		}

		multisig := cryptotypes.MultiSignature{
			Signatures: sigs,
		}
		sig, err := multisig.Marshal()
		if err != nil {
			panic(err)
		}

		return &tx.ModeInfo{
			Sum: &tx.ModeInfo_Multi_{
				Multi: &tx.ModeInfo_Multi{
					Bitarray:  data.BitArray,
					ModeInfos: modeInfos,
				},
			},
		}, sig
	default:
		panic(fmt.Sprintf("unexpected signature data type %T", data))
	}
}
