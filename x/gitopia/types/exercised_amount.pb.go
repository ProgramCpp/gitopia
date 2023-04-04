// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/exercised_amount.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExercisedAmount struct {
	Address string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount  types.Coin `protobuf:"bytes,2,opt,name=amount,proto3,castkey=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
}

func (m *ExercisedAmount) Reset()         { *m = ExercisedAmount{} }
func (m *ExercisedAmount) String() string { return proto.CompactTextString(m) }
func (*ExercisedAmount) ProtoMessage()    {}
func (*ExercisedAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2f495f5b056fe64, []int{0}
}
func (m *ExercisedAmount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExercisedAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExercisedAmount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExercisedAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExercisedAmount.Merge(m, src)
}
func (m *ExercisedAmount) XXX_Size() int {
	return m.Size()
}
func (m *ExercisedAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_ExercisedAmount.DiscardUnknown(m)
}

var xxx_messageInfo_ExercisedAmount proto.InternalMessageInfo

func (m *ExercisedAmount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ExercisedAmount) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*ExercisedAmount)(nil), "gitopia.gitopia.gitopia.ExercisedAmount")
}

func init() { proto.RegisterFile("gitopia/exercised_amount.proto", fileDescriptor_e2f495f5b056fe64) }

var fileDescriptor_e2f495f5b056fe64 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0x2c, 0xc9,
	0x2f, 0xc8, 0x4c, 0xd4, 0x4f, 0xad, 0x48, 0x2d, 0x4a, 0xce, 0x2c, 0x4e, 0x4d, 0x89, 0x4f, 0xcc,
	0xcd, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0xca, 0xeb, 0xa1,
	0xd1, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x35, 0xfa, 0x20, 0x16, 0x44, 0xb9, 0x94, 0x5c,
	0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0xaa, 0x7e, 0x99, 0x61, 0x52, 0x6a,
	0x49, 0xa2, 0xa1, 0x7e, 0x72, 0x7e, 0x66, 0x1e, 0x44, 0x5e, 0xa9, 0x9f, 0x91, 0x8b, 0xdf, 0x15,
	0x66, 0x93, 0x23, 0xd8, 0x22, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x28, 0x89, 0x8b, 0x0d, 0xe2, 0x18, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x49, 0x3d, 0x88, 0xf1, 0x7a, 0x20, 0xe3, 0xf5, 0xa0, 0xc6,
	0xeb, 0x39, 0xe7, 0x67, 0xe6, 0x39, 0xe9, 0x9f, 0xb8, 0x27, 0xcf, 0xd0, 0x74, 0x5f, 0x5e, 0x3d,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xea, 0x16, 0x08, 0xa5, 0x5b,
	0x9c, 0x92, 0xad, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0x0c, 0xd6, 0x10, 0x04, 0x35, 0xd9, 0xc9, 0xe5,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0x90, 0x8c, 0x82, 0x85, 0x12,
	0x8c, 0xae, 0x80, 0xb3, 0xc0, 0x46, 0x26, 0xb1, 0x81, 0xbd, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x20, 0xbe, 0x43, 0xb6, 0x4f, 0x01, 0x00, 0x00,
}

func (m *ExercisedAmount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExercisedAmount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExercisedAmount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintExercisedAmount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintExercisedAmount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExercisedAmount(dAtA []byte, offset int, v uint64) int {
	offset -= sovExercisedAmount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExercisedAmount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovExercisedAmount(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovExercisedAmount(uint64(l))
	return n
}

func sovExercisedAmount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExercisedAmount(x uint64) (n int) {
	return sovExercisedAmount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExercisedAmount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExercisedAmount
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
			return fmt.Errorf("proto: ExercisedAmount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExercisedAmount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExercisedAmount
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
				return ErrInvalidLengthExercisedAmount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExercisedAmount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExercisedAmount
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
				return ErrInvalidLengthExercisedAmount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExercisedAmount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExercisedAmount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExercisedAmount
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
func skipExercisedAmount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExercisedAmount
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
					return 0, ErrIntOverflowExercisedAmount
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
					return 0, ErrIntOverflowExercisedAmount
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
				return 0, ErrInvalidLengthExercisedAmount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExercisedAmount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExercisedAmount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExercisedAmount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExercisedAmount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExercisedAmount = fmt.Errorf("proto: unexpected end of group")
)
