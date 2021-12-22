package signing

import (
	cryptotypes "cosmossdk/types"
	"cosmossdk/types"
	"fmt"
)

// SignatureV2 is a convenience type that is easier to use in application logic
// than the protobuf SignerInfo's and raw signature bytes. It goes beyond the
// first sdk.Signature types by supporting sign modes and explicitly nested
// multi-signatures. It is intended to be used for both building and verifying
// signatures.
type SignatureV2 struct {
	// PubKey is the public key to use for verifying the signature
	PubKey cryptotypes.PubKey

	// Data is the actual data of the signature which includes SignMode's and
	// the signatures themselves for either single or multi-signatures.
	Data SignatureData

	// Sequence is the sequence of this account. Only populated in
	// SIGN_MODE_DIRECT.
	Sequence uint64
}

// SignatureDataToProto converts a SignatureData to SignatureDescriptor_Data.
// SignatureDescriptor_Data is considered an encoding type whereas SignatureData is used for
// business logic.
func SignatureDataToProto(data SignatureData) *SignatureDescriptor_Data {
	switch data := data.(type) {
	case *SingleSignatureData:
		return &SignatureDescriptor_Data{
			Sum: &SignatureDescriptor_Data_Single_{
				Single: &SignatureDescriptor_Data_Single{
					Mode:      data.SignMode,
					Signature: data.Signature,
				},
			},
		}
	case *MultiSignatureData:
		descDatas := make([]*SignatureDescriptor_Data, len(data.Signatures))

		for j, d := range data.Signatures {
			descDatas[j] = SignatureDataToProto(d)
		}

		return &SignatureDescriptor_Data{
			Sum: &SignatureDescriptor_Data_Multi_{
				Multi: &SignatureDescriptor_Data_Multi{
					Bitarray:   data.BitArray,
					Signatures: descDatas,
				},
			},
		}
	default:
		panic(fmt.Errorf("unexpected case %+v", data))
	}
}

// SignatureDataFromProto converts a SignatureDescriptor_Data to SignatureData.
// SignatureDescriptor_Data is considered an encoding type whereas SignatureData is used for
// business logic.
func SignatureDataFromProto(descData *SignatureDescriptor_Data) SignatureData {
	switch descData := descData.Sum.(type) {
	case *SignatureDescriptor_Data_Single_:
		return &SingleSignatureData{
			SignMode:  descData.Single.Mode,
			Signature: descData.Single.Signature,
		}
	case *SignatureDescriptor_Data_Multi_:
		multi := descData.Multi
		datas := make([]SignatureData, len(multi.Signatures))

		for j, d := range multi.Signatures {
			datas[j] = SignatureDataFromProto(d)
		}

		return &MultiSignatureData{
			BitArray:   multi.Bitarray,
			Signatures: datas,
		}
	default:
		panic(fmt.Errorf("unexpected case %+v", descData))
	}
}

type UnpackInterfacesMessage interface {
	// UnpackInterfaces is implemented in order to unpack values packed within
	// Any's using the AnyUnpacker. It should generally be implemented as
	// follows:
	//   func (s *MyStruct) UnpackInterfaces(unpacker AnyUnpacker) error {
	//		var x AnyInterface
	//		// where X is an Any field on MyStruct
	//		err := unpacker.UnpackAny(s.X, &x)
	//		if err != nil {
	//			return nil
	//		}
	//		// where Y is a field on MyStruct that implements UnpackInterfacesMessage itself
	//		err = s.Y.UnpackInterfaces(unpacker)
	//		if err != nil {
	//			return nil
	//		}
	//		return nil
	//	 }
	UnpackInterfaces(unpacker AnyUnpacker) error
}

type AnyUnpacker interface {
	// UnpackAny unpacks the value in any to the interface pointer passed in as
	// iface. Note that the type in any must have been registered in the
	// underlying whitelist registry as a concrete type for that interface
	// Ex:
	//    var msg sdk.Msg
	//    err := cdc.UnpackAny(any, &msg)
	//    ...
	UnpackAny(any *types.Any, iface interface{}) error
}

var _, _ UnpackInterfacesMessage = &SignatureDescriptors{}, &SignatureDescriptor{}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (sds *SignatureDescriptors) UnpackInterfaces(unpacker AnyUnpacker) error {
	for _, sig := range sds.Signatures {
		err := sig.UnpackInterfaces(unpacker)

		if err != nil {
			return err
		}
	}

	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (sd *SignatureDescriptor) UnpackInterfaces(unpacker AnyUnpacker) error {
	return unpacker.UnpackAny(sd.PublicKey, new(cryptotypes.PubKey))
}
