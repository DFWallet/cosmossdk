package txBody

import (
	"github.com/DFWallet/cosmossdk/protoTx/proto"
	"github.com/DFWallet/cosmossdk/types"
	"github.com/DFWallet/cosmossdk/types/signing"
	"fmt"
	"io"
)

// SignerInfo describes the public key and signing mode of a single top-level
// signer.
type SignerInfo struct {
	// public_key is the public key of the signer. It is optional for accounts
	// that already exist in state. If unset, the verifier can use the required \
	// signer address for this position and lookup the public key.
	PublicKey *types.Any `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// mode_info describes the signing mode of the signer and is a nested
	// structure to support nested multisig pubkey's
	ModeInfo *ModeInfo `protobuf:"bytes,2,opt,name=mode_info,json=modeInfo,proto3" json:"mode_info,omitempty"`
	// sequence is the sequence of the account, which describes the
	// number of committed transactions signed by a given address. It is used to
	// prevent replay attacks.
	Sequence uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *SignerInfo) Reset()         { *m = SignerInfo{} }
func (m *SignerInfo) String() string { return proto.CompactTextString(m) }
func (*SignerInfo) ProtoMessage()    {}
func (*SignerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{5}
}
func (m *SignerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &types.Any{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ModeInfo == nil {
				m.ModeInfo = &ModeInfo{}
			}
			if err := m.ModeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *ModeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ModeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Single", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ModeInfo_Single{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ModeInfo_Single_{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multi", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ModeInfo_Multi{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ModeInfo_Multi_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *ModeInfo_Single) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Single: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Single: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= signing.SignMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *ModeInfo_Single_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *ModeInfo_Single_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Single != nil {
		{
			size, err := m.Single.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ModeInfo_Single) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mode != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ModeInfo_Single_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Single != nil {
		l = m.Single.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}
func (m *ModeInfo_Single) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mode != 0 {
		n += 1 + sovTx(uint64(m.Mode))
	}
	return n
}

type SignMode int32

func (x SignMode) String() string {
	return proto.EnumName(SignMode_name, int32(x))
}

var SignMode_name = map[int32]string{
	0:   "SIGN_MODE_UNSPECIFIED",
	1:   "SIGN_MODE_DIRECT",
	2:   "SIGN_MODE_TEXTUAL",
	127: "SIGN_MODE_LEGACY_AMINO_JSON",
}
var fileDescriptor_9a54958ff3d0b1b9 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xed, 0x26, 0x8d, 0xd2, 0x29, 0x42, 0x61, 0x49, 0xa5, 0xc4, 0x20, 0x13, 0x95, 0x03,
	0x11, 0x52, 0xd6, 0x6a, 0x72, 0x40, 0x70, 0xcb, 0x1f, 0x93, 0x86, 0x36, 0x09, 0xd8, 0xa9, 0x04,
	0x5c, 0x2c, 0xdb, 0xd9, 0x1a, 0xab, 0xb1, 0xd7, 0x78, 0xd7, 0xa8, 0x3e, 0xf1, 0x0a, 0xbc, 0x06,
	0xcf, 0xc1, 0x85, 0x63, 0x8f, 0x1c, 0x51, 0xf2, 0x0c, 0xdc, 0x51, 0xec, 0x38, 0x09, 0x52, 0x11,
	0x22, 0x27, 0x6b, 0x66, 0xbe, 0xfd, 0xcd, 0xb7, 0x9a, 0x59, 0xc3, 0x13, 0x9b, 0x32, 0x8f, 0x32,
	0x85, 0x5f, 0x2b, 0xcc, 0x75, 0x7c, 0xd7, 0x77, 0x94, 0x4f, 0x27, 0x16, 0xe1, 0xe6, 0x49, 0x16,
	0xe3, 0x20, 0xa4, 0x9c, 0xa2, 0x6a, 0x2a, 0xc4, 0xfc, 0x1a, 0x67, 0x85, 0x95, 0x50, 0x6a, 0xac,
	0x18, 0x76, 0x18, 0x07, 0x9c, 0x2a, 0x5e, 0x34, 0xe3, 0x2e, 0x73, 0x37, 0xa0, 0x2c, 0x91, 0x92,
	0xa4, 0xaa, 0x43, 0xa9, 0x33, 0x23, 0x4a, 0x12, 0x59, 0xd1, 0xa5, 0x62, 0xfa, 0x71, 0x5a, 0x3a,
	0xbe, 0x84, 0xb2, 0xee, 0x3a, 0xbe, 0xc9, 0xa3, 0x90, 0xf4, 0x08, 0xb3, 0x43, 0x37, 0xe0, 0x34,
	0x64, 0x68, 0x04, 0xc0, 0xb2, 0x3c, 0xab, 0x88, 0xb5, 0x5c, 0xfd, 0xb0, 0x89, 0xf1, 0x5f, 0x1d,
	0xe1, 0x5b, 0x20, 0xda, 0x16, 0xe1, 0xf8, 0x57, 0x1e, 0xee, 0xdf, 0xa2, 0x41, 0x2d, 0x80, 0x20,
	0xb2, 0x66, 0xae, 0x6d, 0x5c, 0x91, 0xb8, 0x22, 0xd6, 0xc4, 0xfa, 0x61, 0xb3, 0x8c, 0x53, 0xbf,
	0x38, 0xf3, 0x8b, 0xdb, 0x7e, 0xac, 0x1d, 0xa4, 0xba, 0x33, 0x12, 0xa3, 0x3e, 0xe4, 0xa7, 0x26,
	0x37, 0x2b, 0x7b, 0x89, 0xbc, 0xf5, 0x7f, 0xb6, 0x70, 0xcf, 0xe4, 0xa6, 0x96, 0x00, 0x90, 0x04,
	0x45, 0x46, 0x3e, 0x46, 0xc4, 0xb7, 0x49, 0x25, 0x57, 0x13, 0xeb, 0x79, 0x6d, 0x1d, 0x4b, 0xdf,
	0x72, 0x90, 0x5f, 0x4a, 0xd1, 0x04, 0x0a, 0xcc, 0xf5, 0x9d, 0x19, 0x59, 0xd9, 0x7b, 0xb1, 0x43,
	0x3f, 0xac, 0x27, 0x84, 0x53, 0x41, 0x5b, 0xb1, 0xd0, 0x1b, 0xd8, 0x4f, 0xa6, 0xb4, 0xba, 0xc4,
	0xf3, 0x5d, 0xa0, 0xc3, 0x25, 0xe0, 0x54, 0xd0, 0x52, 0x92, 0x64, 0x40, 0x21, 0x6d, 0x83, 0x9e,
	0x41, 0xde, 0xa3, 0xd3, 0xd4, 0xf0, 0xdd, 0xe6, 0xe3, 0x7f, 0xb0, 0x87, 0x74, 0x4a, 0xb4, 0xe4,
	0x00, 0x7a, 0x08, 0x07, 0xeb, 0xa1, 0x25, 0xce, 0xee, 0x68, 0x9b, 0x84, 0xf4, 0x55, 0x84, 0xfd,
	0xa4, 0x27, 0x3a, 0x83, 0xa2, 0xe5, 0x72, 0x33, 0x0c, 0xcd, 0x6c, 0x68, 0x4a, 0xd6, 0x24, 0xdd,
	0x49, 0xbc, 0x5e, 0xc1, 0xac, 0x53, 0x97, 0x7a, 0x81, 0x69, 0xf3, 0x8e, 0xcb, 0xdb, 0xcb, 0x63,
	0xda, 0x1a, 0x80, 0xf4, 0x3f, 0x76, 0x6d, 0x2f, 0xd9, 0xb5, 0x9d, 0x86, 0xba, 0x85, 0xe9, 0xec,
	0x43, 0x8e, 0x45, 0xde, 0x53, 0x06, 0xc5, 0xec, 0x8a, 0xa8, 0x0a, 0x47, 0xfa, 0xa0, 0x3f, 0x32,
	0x86, 0xe3, 0x9e, 0x6a, 0x5c, 0x8c, 0xf4, 0xd7, 0x6a, 0x77, 0xf0, 0x72, 0xa0, 0xf6, 0x4a, 0x02,
	0x2a, 0x43, 0x69, 0x53, 0xea, 0x0d, 0x34, 0xb5, 0x3b, 0x29, 0x89, 0xe8, 0x08, 0xee, 0x6d, 0xb2,
	0x13, 0xf5, 0xed, 0xe4, 0xa2, 0x7d, 0x5e, 0xda, 0x43, 0x8f, 0xe0, 0xc1, 0x26, 0x7d, 0xae, 0xf6,
	0xdb, 0xdd, 0x77, 0x46, 0x7b, 0x38, 0x18, 0x8d, 0x8d, 0x57, 0xfa, 0x78, 0x54, 0xfa, 0xdc, 0xe9,
	0x7f, 0x9f, 0xcb, 0xe2, 0xcd, 0x5c, 0x16, 0x7f, 0xce, 0x65, 0xf1, 0xcb, 0x42, 0x16, 0x6e, 0x16,
	0xb2, 0xf0, 0x63, 0x21, 0x0b, 0xef, 0x1b, 0x8e, 0xcb, 0x3f, 0x44, 0x16, 0xb6, 0xa9, 0xa7, 0x64,
	0x6f, 0x38, 0xf9, 0x34, 0xd8, 0xf4, 0x4a, 0xe1, 0x71, 0x40, 0xb6, 0x7f, 0x0c, 0x56, 0x21, 0x79,
	0x01, 0xad, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x87, 0x55, 0xbf, 0x34, 0x04, 0x00, 0x00,
}

func (SignMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9a54958ff3d0b1b9, []int{0}
}

type ModeInfo_Single struct {
	// mode is the signing mode of the single signer
	Mode signing.SignMode `protobuf:"varint,1,opt,name=mode,proto3,enum=cosmos.tx.signing.v1beta1.SignMode" json:"mode,omitempty"`
}

type ModeInfo_Single_ struct {
	Single *ModeInfo_Single `protobuf:"bytes,1,opt,name=single,proto3,oneof" json:"single,omitempty"`
}

type ModeInfo_Multi_ struct {
	Multi *ModeInfo_Multi `protobuf:"bytes,2,opt,name=multi,proto3,oneof" json:"multi,omitempty"`
}

func (*ModeInfo_Single_) isModeInfo_Sum() {}
func (*ModeInfo_Multi_) isModeInfo_Sum()  {}


func (m *ModeInfo_Multi_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModeInfo_Multi_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Multi != nil {
		l = m.Multi.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}
func (m *ModeInfo_Multi) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bitarray != nil {
		l = m.Bitarray.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ModeInfos) > 0 {
		for _, e := range m.ModeInfos {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}
func (m *ModeInfo_Multi_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Multi != nil {
		{
			size, err := m.Multi.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ModeInfo_Multi) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModeInfos) > 0 {
		for iNdEx := len(m.ModeInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ModeInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Bitarray != nil {
		{
			size, err := m.Bitarray.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

type ModeInfo_Multi struct {
	// bitarray specifies which keys within the multisig are signing
	Bitarray *types.CompactBitArray `protobuf:"bytes,1,opt,name=bitarray,proto3" json:"bitarray,omitempty"`
	// mode_infos is the corresponding modes of the signers of the multisig
	// which could include nested multisig public keys
	ModeInfos []*ModeInfo `protobuf:"bytes,2,rep,name=mode_infos,json=modeInfos,proto3" json:"mode_infos,omitempty"`
}

func (m *ModeInfo_Multi) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Multi: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Multi: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bitarray", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bitarray == nil {
				m.Bitarray = &types.CompactBitArray{}
			}
			if err := m.Bitarray.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModeInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModeInfos = append(m.ModeInfos, &ModeInfo{})
			if err := m.ModeInfos[len(m.ModeInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *SignerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if m.ModeInfo != nil {
		{
			size, err := m.ModeInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *SignerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignerInfo.Merge(m, src)
}
func (m *SignerInfo) XXX_Size() int {
	return m.Size()
}
func (m *SignerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ModeInfo != nil {
		l = m.ModeInfo.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovTx(uint64(m.Sequence))
	}
	return n
}
func (m *ModeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}
func (m *SignerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SignerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SignerInfo proto.InternalMessageInfo

func (m *SignerInfo) GetPublicKey() *types.Any {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SignerInfo) GetModeInfo() *ModeInfo {
	if m != nil {
		return m.ModeInfo
	}
	return nil
}

func (m *SignerInfo) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// Fee includes the amount of coins paid in fees and the maximum
// gas to be used by the transaction. The ratio yields an effective "gasprice",
// which must be above some miminum to be accepted into the mempool.
type Fee struct {
	// amount is the amount of coins to be paid as a fee
	Amount types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	// gas_limit is the maximum gas that can be used in transaction processing
	// before an out of gas error occurs
	GasLimit uint64 `protobuf:"varint,2,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// if unset, the first signer is responsible for paying the fees. If set, the specified account must pay the fees.
	// the payer must be a tx signer (and thus have signed this field in AuthInfo).
	// setting this field does *not* change the ordering of required signers for the transaction.
	Payer string `protobuf:"bytes,3,opt,name=payer,proto3" json:"payer,omitempty"`
	// if set, the fee payer (either the first signer or the value of the payer field) requests that a fee grant be used
	// to pay fees instead of the fee payer's own balance. If an appropriate fee grant does not exist or the chain does
	// not support fee grants, this will fail
	Granter string `protobuf:"bytes,4,opt,name=granter,proto3" json:"granter,omitempty"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{7}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Payer) > 0 {
		i -= len(m.Payer)
		copy(dAtA[i:], m.Payer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Payer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GasLimit != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}

func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.GasLimit != 0 {
		n += 1 + sovTx(uint64(m.GasLimit))
	}
	l = len(m.Payer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetAmount() types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Fee) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *Fee) GetPayer() string {
	if m != nil {
		return m.Payer
	}
	return ""
}

func (m *Fee) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

// ModeInfo describes the signing mode of a single or nested multisig signer.
type ModeInfo struct {
	// sum is the oneof that specifies whether this represents a single or nested
	// multisig signer
	//
	// Types that are valid to be assigned to Sum:
	//	*ModeInfo_Single_
	//	*ModeInfo_Multi_
	Sum isModeInfo_Sum `protobuf_oneof:"sum"`
}

func (m *ModeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

type isModeInfo_Sum interface {
	isModeInfo_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}


// SignDoc is the type used for generating sign bytes for SIGN_MODE_DIRECT.
type SignDoc struct {
	// body_bytes is protobuf serialization of a TxBody that matches the
	// representation in TxRaw.
	BodyBytes []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	// auth_info_bytes is a protobuf serialization of an AuthInfo that matches the
	// representation in TxRaw.
	AuthInfoBytes []byte `protobuf:"bytes,2,opt,name=auth_info_bytes,json=authInfoBytes,proto3" json:"auth_info_bytes,omitempty"`
	// chain_id is the unique identifier of the chain this transaction targets.
	// It prevents signed transactions from being used on another chain by an
	// attacker
	ChainId string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// account_number is the account number of the account in state
	AccountNumber uint64 `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
}

func (m *SignDoc) Reset()         { *m = SignDoc{} }
func (m *SignDoc) String() string { return proto.CompactTextString(m) }
func (*SignDoc) ProtoMessage()    {}
func (*SignDoc) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{2}
}
func (m *SignDoc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SignDoc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *SignDoc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignDoc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignDoc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BodyBytes = append(m.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.BodyBytes == nil {
				m.BodyBytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthInfoBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthInfoBytes = append(m.AuthInfoBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthInfoBytes == nil {
				m.AuthInfoBytes = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignDoc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignDoc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignDoc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccountNumber != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AuthInfoBytes) > 0 {
		i -= len(m.AuthInfoBytes)
		copy(dAtA[i:], m.AuthInfoBytes)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AuthInfoBytes)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BodyBytes) > 0 {
		i -= len(m.BodyBytes)
		copy(dAtA[i:], m.BodyBytes)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BodyBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *SignDoc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignDoc.Merge(m, src)
}
func (m *SignDoc) XXX_Size() int {
	return m.Size()
}
func (m *SignDoc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BodyBytes)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AuthInfoBytes)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovTx(uint64(m.AccountNumber))
	}
	return n
}
func (m *SignDoc) XXX_DiscardUnknown() {
	xxx_messageInfo_SignDoc.DiscardUnknown(m)
}

var xxx_messageInfo_SignDoc proto.InternalMessageInfo

func (m *SignDoc) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

func (m *SignDoc) GetAuthInfoBytes() []byte {
	if m != nil {
		return m.AuthInfoBytes
	}
	return nil
}

func (m *SignDoc) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *SignDoc) GetAccountNumber() uint64 {
	if m != nil {
		return m.AccountNumber
	}
	return 0
}

