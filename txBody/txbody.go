package txBody

import (
	"cosmossdk/protoTx/proto"
	"cosmossdk/types"
	"fmt"
	"io"
	math_bits "math/bits"
)

var fileDescriptor_96d1575ffde80842 = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x5e, 0xef, 0x5f, 0xd6, 0x27, 0x49, 0x4b, 0x47, 0x11, 0xda, 0x6c, 0x54, 0x37, 0x18, 0x15,
	0xf6, 0x26, 0x76, 0x9b, 0x5e, 0xf0, 0x23, 0x24, 0xc8, 0x16, 0xaa, 0x54, 0xa5, 0x20, 0x4d, 0x72,
	0xd5, 0x1b, 0x6b, 0xec, 0x9d, 0x78, 0x47, 0x5d, 0xcf, 0x2c, 0x9e, 0x71, 0xb1, 0x1f, 0x02, 0xa9,
	0x42, 0x42, 0xbc, 0x03, 0x2f, 0xc0, 0x2b, 0xf4, 0xb2, 0x97, 0x5c, 0x41, 0x95, 0x3c, 0x08, 0x68,
	0xc6, 0x63, 0x27, 0x82, 0x55, 0x72, 0xc3, 0x95, 0xe7, 0x9c, 0xf9, 0xce, 0x37, 0x9f, 0xcf, 0x1f,
	0x4c, 0x12, 0x21, 0x33, 0x21, 0x43, 0x55, 0x86, 0xaf, 0x1e, 0xc6, 0x54, 0x91, 0x87, 0xa1, 0x2a,
	0x83, 0x55, 0x2e, 0x94, 0x40, 0x77, 0xea, 0xbb, 0x40, 0x95, 0x81, 0xbd, 0x9b, 0xec, 0xa4, 0x22,
	0x15, 0xe6, 0x36, 0xd4, 0xa7, 0x1a, 0x38, 0x39, 0xb0, 0x24, 0x49, 0x5e, 0xad, 0x94, 0x08, 0xb3,
	0x62, 0xa9, 0x98, 0x64, 0x69, 0xcb, 0xd8, 0x38, 0x2c, 0xdc, 0xb3, 0xf0, 0x98, 0x48, 0xda, 0x62,
	0x12, 0xc1, 0xb8, 0xbd, 0xff, 0xf8, 0x52, 0x93, 0x64, 0x29, 0x67, 0xfc, 0x92, 0xc9, 0xda, 0x16,
	0xb8, 0x9b, 0x0a, 0x91, 0x2e, 0x69, 0x68, 0xac, 0xb8, 0x38, 0x0b, 0x09, 0xaf, 0xea, 0x2b, 0xff,
	0x27, 0x07, 0xba, 0xa7, 0x25, 0x3a, 0x80, 0x7e, 0x2c, 0xe6, 0xd5, 0xd8, 0xd9, 0x77, 0xa6, 0x9b,
	0x87, 0xbb, 0xc1, 0x7f, 0xfe, 0x28, 0x38, 0x2d, 0x67, 0x62, 0x5e, 0x61, 0x03, 0x43, 0x9f, 0x82,
	0x4b, 0x0a, 0xb5, 0x88, 0x18, 0x3f, 0x13, 0xe3, 0xae, 0x89, 0xd9, 0x5b, 0x13, 0x73, 0x54, 0xa8,
	0xc5, 0x53, 0x7e, 0x26, 0xf0, 0x88, 0xd8, 0x13, 0xf2, 0x00, 0xb4, 0x36, 0xa2, 0x8a, 0x9c, 0xca,
	0x71, 0x6f, 0xbf, 0x37, 0xdd, 0xc2, 0x57, 0x3c, 0x3e, 0x87, 0xc1, 0x69, 0x89, 0xc9, 0x8f, 0xe8,
	0x2e, 0x80, 0x7e, 0x2a, 0x8a, 0x2b, 0x45, 0xa5, 0xd1, 0xb5, 0x85, 0x5d, 0xed, 0x99, 0x69, 0x07,
	0xfa, 0x08, 0x6e, 0xb7, 0x0a, 0x2c, 0xa6, 0x6b, 0x30, 0xdb, 0xcd, 0x53, 0x35, 0xee, 0xa6, 0xf7,
	0x7e, 0x76, 0x60, 0xe3, 0x84, 0xa5, 0xfc, 0x6b, 0x91, 0xfc, 0x5f, 0x4f, 0xee, 0xc2, 0x28, 0x59,
	0x10, 0xc6, 0x23, 0x36, 0x1f, 0xf7, 0xf6, 0x9d, 0xa9, 0x8b, 0x37, 0x8c, 0xfd, 0x74, 0x8e, 0xee,
	0xc3, 0x2d, 0x92, 0x24, 0xa2, 0xe0, 0x2a, 0xe2, 0x45, 0x16, 0xd3, 0x7c, 0xdc, 0xdf, 0x77, 0xa6,
	0x7d, 0xbc, 0x6d, 0xbd, 0xdf, 0x19, 0xa7, 0xff, 0x4b, 0x17, 0x86, 0x75, 0xbe, 0xd1, 0x03, 0x18,
	0x65, 0x54, 0x4a, 0x92, 0x1a, 0x45, 0xbd, 0xe9, 0xe6, 0xe1, 0x4e, 0x50, 0x57, 0x33, 0x68, 0xaa,
	0x19, 0x1c, 0xf1, 0x0a, 0xb7, 0x28, 0x84, 0xa0, 0x9f, 0xd1, 0xac, 0x2e, 0x8b, 0x8b, 0xcd, 0x59,
	0xbf, 0xab, 0x58, 0x46, 0x45, 0xa1, 0xa2, 0x05, 0x65, 0xe9, 0x42, 0x19, 0x61, 0x7d, 0xbc, 0x6d,
	0xbd, 0xc7, 0xc6, 0x89, 0x66, 0x70, 0x87, 0x96, 0x8a, 0x72, 0xc9, 0x04, 0x8f, 0xc4, 0x4a, 0x31,
	0xc1, 0xe5, 0xf8, 0xef, 0x8d, 0x6b, 0x9e, 0x7d, 0xaf, 0xc5, 0x7f, 0x5f, 0xc3, 0xd1, 0x0b, 0xf0,
	0xb8, 0xe0, 0x51, 0x92, 0x33, 0xc5, 0x12, 0xb2, 0x8c, 0xd6, 0x10, 0xde, 0xbe, 0x86, 0x70, 0x8f,
	0x0b, 0xfe, 0xd8, 0xc6, 0x7e, 0xf3, 0x2f, 0x6e, 0xff, 0x15, 0x8c, 0x9a, 0x96, 0x42, 0x5f, 0xc1,
	0x96, 0x2e, 0x23, 0xcd, 0x4d, 0x3d, 0x9a, 0xe4, 0xdc, 0x5d, 0xd3, 0x85, 0x27, 0x06, 0x66, 0xfa,
	0x70, 0x53, 0xb6, 0x67, 0x89, 0xa6, 0xd0, 0x3b, 0xa3, 0xd4, 0xb6, 0xef, 0xfb, 0x6b, 0x02, 0x9f,
	0x50, 0x8a, 0x35, 0xc4, 0xff, 0xd5, 0x01, 0xb8, 0x64, 0x41, 0x8f, 0x00, 0x56, 0x45, 0xbc, 0x64,
	0x49, 0xf4, 0x92, 0x36, 0x23, 0xb3, 0xfe, 0x6f, 0xdc, 0x1a, 0xf7, 0x8c, 0x9a, 0x91, 0xc9, 0xc4,
	0x9c, 0xde, 0x34, 0x32, 0xcf, 0xc5, 0x9c, 0xd6, 0x23, 0x93, 0xd9, 0x13, 0x9a, 0xc0, 0x48, 0xd2,
	0x1f, 0x0a, 0xca, 0x13, 0x6a, 0xcb, 0xd6, 0xda, 0xfe, 0xbb, 0x2e, 0x8c, 0x9a, 0x10, 0xf4, 0x05,
	0x0c, 0x25, 0xe3, 0xe9, 0x92, 0x5a, 0x4d, 0xfe, 0x35, 0xfc, 0xc1, 0x89, 0x41, 0x1e, 0x77, 0xb0,
	0x8d, 0x41, 0x9f, 0xc1, 0xc0, 0xec, 0x1f, 0x2b, 0xee, 0x83, 0xeb, 0x82, 0x9f, 0x6b, 0xe0, 0x71,
	0x07, 0xd7, 0x11, 0x93, 0x23, 0x18, 0xd6, 0x74, 0xe8, 0x13, 0xe8, 0x6b, 0xdd, 0x46, 0xc0, 0xad,
	0xc3, 0x0f, 0xaf, 0x70, 0x34, 0x1b, 0xe9, 0x6a, 0x55, 0x34, 0x1f, 0x36, 0x01, 0x93, 0xd7, 0x0e,
	0x0c, 0x0c, 0x2b, 0x7a, 0x06, 0xa3, 0x98, 0x29, 0x92, 0xe7, 0xa4, 0xc9, 0x6d, 0xd8, 0xd0, 0xd4,
	0x7b, 0x33, 0x68, 0xd7, 0x64, 0xc3, 0xf5, 0x58, 0x64, 0x2b, 0x92, 0xa8, 0x19, 0x53, 0x47, 0x3a,
	0x0c, 0xb7, 0x04, 0xe8, 0x73, 0x80, 0x36, 0xeb, 0x7a, 0x5c, 0x7b, 0x37, 0xa5, 0xdd, 0x6d, 0xd2,
	0x2e, 0x67, 0x03, 0xe8, 0xc9, 0x22, 0xf3, 0x7f, 0x77, 0xa0, 0xf7, 0x84, 0x52, 0x94, 0xc0, 0x90,
	0x64, 0x7a, 0x48, 0x6d, 0xab, 0xb5, 0x4b, 0x52, 0xaf, 0xe7, 0x2b, 0x52, 0x18, 0x9f, 0x3d, 0x78,
	0xf3, 0xe7, 0xbd, 0xce, 0x6f, 0x7f, 0xdd, 0x9b, 0xa6, 0x4c, 0x2d, 0x8a, 0x38, 0x48, 0x44, 0x16,
	0x36, 0xab, 0xdf, 0x7c, 0x0e, 0xe4, 0xfc, 0x65, 0xa8, 0xaa, 0x15, 0x95, 0x26, 0x40, 0x62, 0x4b,
	0x8d, 0xf6, 0xc0, 0x4d, 0x89, 0x8c, 0x96, 0x2c, 0x63, 0xca, 0x14, 0xa2, 0x8f, 0x47, 0x29, 0x91,
	0xdf, 0x6a, 0x1b, 0xed, 0xc0, 0x60, 0x45, 0x2a, 0x9a, 0xdb, 0xad, 0x52, 0x1b, 0x68, 0x0c, 0x1b,
	0x69, 0x4e, 0xb8, 0xb2, 0xcb, 0xc4, 0xc5, 0x8d, 0x39, 0xfb, 0xf2, 0xcd, 0xb9, 0xe7, 0xbc, 0x3d,
	0xf7, 0x9c, 0x77, 0xe7, 0x9e, 0xf3, 0xfa, 0xc2, 0xeb, 0xbc, 0xbd, 0xf0, 0x3a, 0x7f, 0x5c, 0x78,
	0x9d, 0x17, 0xf7, 0x6f, 0x16, 0x16, 0xaa, 0x32, 0x1e, 0x9a, 0x66, 0x7e, 0xf4, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0xb7, 0x75, 0x7d, 0xfd, 0x06, 0x00, 0x00,
}

// TxBody is the body of a transaction that all signers sign over.
type TxBody struct {
	// messages is a list of messages to be executed. The required signers of
	// those messages define the number and order of elements in AuthInfo's
	// signer_infos and Tx's signatures. Each required signer address is added to
	// the list only the first time it occurs.
	// By convention, the first required signer (usually from the first message)
	// is referred to as the primary signer and pays the fee for the whole
	// transaction.
	Messages []*types.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	// memo is any arbitrary note/comment to be added to the transaction.
	// WARNING: in clients, any publicly exposed text should not be called memo,
	// but should be called `note` instead (see https://github.com/cosmos/cosmos-sdk/issues/9122).
	Memo string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	// timeout is the block height after which this transaction will not
	// be processed by the chain
	TimeoutHeight uint64 `protobuf:"varint,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	// extension_options are arbitrary options that can be added by chains
	// when the default options are not sufficient. If any of these are present
	// and can't be handled, the transaction will be rejected
	ExtensionOptions []*types.Any `protobuf:"bytes,1023,rep,name=extension_options,json=extensionOptions,proto3" json:"extension_options,omitempty"`
	// extension_options are arbitrary options that can be added by chains
	// when the default options are not sufficient. If any of these are present
	// and can't be handled, they will be ignored
	NonCriticalExtensionOptions []*types.Any `protobuf:"bytes,2047,rep,name=non_critical_extension_options,json=nonCriticalExtensionOptions,proto3" json:"non_critical_extension_options,omitempty"`
}

func (m *TxBody) Reset()         { *m = TxBody{} }
func (m *TxBody) String() string { return proto.CompactTextString(m) }
func (*TxBody) ProtoMessage()    {}
func (*TxBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d1575ffde80842, []int{3}
}
func (m *TxBody) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxBody.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxBody.Merge(m, src)
}
func (m *TxBody) XXX_Size() int {
	return m.Size()
}
func (m *TxBody) XXX_DiscardUnknown() {
	xxx_messageInfo_TxBody.DiscardUnknown(m)
}

var xxx_messageInfo_TxBody proto.InternalMessageInfo

func (m *TxBody) GetMessages() []*types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *TxBody) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *TxBody) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *TxBody) GetExtensionOptions() []*types.Any {
	if m != nil {
		return m.ExtensionOptions
	}
	return nil
}

func (m *TxBody) GetNonCriticalExtensionOptions() []*types.Any {
	if m != nil {
		return m.NonCriticalExtensionOptions
	}
	return nil
}

func (m *TxBody) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NonCriticalExtensionOptions) > 0 {
		for iNdEx := len(m.NonCriticalExtensionOptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NonCriticalExtensionOptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7f
			i--
			dAtA[i] = 0xfa
		}
	}
	if len(m.ExtensionOptions) > 0 {
		for iNdEx := len(m.ExtensionOptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExtensionOptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3f
			i--
			dAtA[i] = 0xfa
		}
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func (m *TxBody) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovTx(uint64(m.TimeoutHeight))
	}
	if len(m.ExtensionOptions) > 0 {
		for _, e := range m.ExtensionOptions {
			l = e.Size()
			n += 2 + l + sovTx(uint64(l))
		}
	}
	if len(m.NonCriticalExtensionOptions) > 0 {
		for _, e := range m.NonCriticalExtensionOptions {
			l = e.Size()
			n += 2 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *TxBody) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TxBody: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxBody: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
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
			m.Messages = append(m.Messages, &types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
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
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 1023:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtensionOptions", wireType)
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
			m.ExtensionOptions = append(m.ExtensionOptions, &types.Any{})
			if err := m.ExtensionOptions[len(m.ExtensionOptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2047:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonCriticalExtensionOptions", wireType)
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
			m.NonCriticalExtensionOptions = append(m.NonCriticalExtensionOptions, &types.Any{})
			if err := m.NonCriticalExtensionOptions[len(m.NonCriticalExtensionOptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)