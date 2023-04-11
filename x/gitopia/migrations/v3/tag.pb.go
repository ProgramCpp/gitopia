// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/tag.proto

package types

import (
	fmt "fmt"
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

type Tag struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RepositoryId uint64 `protobuf:"varint,2,opt,name=repositoryId,proto3" json:"repositoryId,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Sha          string `protobuf:"bytes,4,opt,name=sha,proto3" json:"sha,omitempty"`
	CreatedAt    int64  `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    int64  `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd09f5e76807ae94, []int{0}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return m.Size()
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetRepositoryId() uint64 {
	if m != nil {
		return m.RepositoryId
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *Tag) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Tag) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Tag)(nil), "gitopia.gitopia.gitopia.v3.Tag")
}

func init() { proto.RegisterFile("gitopia/tag.proto", fileDescriptor_bd09f5e76807ae94) }

var fileDescriptor_bd09f5e76807ae94 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcf, 0x2c, 0xc9,
	0x2f, 0xc8, 0x4c, 0xd4, 0x2f, 0x49, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x82,
	0x0a, 0xe9, 0xa1, 0xd3, 0x65, 0xc6, 0x4a, 0xb3, 0x19, 0xb9, 0x98, 0x43, 0x12, 0xd3, 0x85, 0xf8,
	0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x53, 0x84, 0x94,
	0xb8, 0x78, 0x8a, 0x52, 0x0b, 0xf2, 0x8b, 0x33, 0x4b, 0xf2, 0x8b, 0x2a, 0x3d, 0x53, 0x24, 0x98,
	0xc0, 0x32, 0x28, 0x62, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0xcc, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x71, 0x46, 0xa2, 0x04, 0x0b, 0x58, 0x08,
	0xc4, 0x14, 0x92, 0xe1, 0xe2, 0x4c, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d, 0x71, 0x2c, 0x91, 0x60,
	0x55, 0x60, 0xd4, 0x60, 0x0e, 0x42, 0x08, 0x80, 0x64, 0x4b, 0x0b, 0x52, 0xa0, 0xb2, 0x6c, 0x10,
	0x59, 0xb8, 0x80, 0x53, 0xd0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59,
	0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3, 0x7c, 0x0c, 0xa3, 0x2b,
	0xe0, 0xac, 0xdc, 0xcc, 0xf4, 0xa2, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0xfd, 0x32, 0x63, 0xeb,
	0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0xa0, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x81, 0xfa, 0xf9, 0x0d, 0x29, 0x01, 0x00, 0x00,
}

func (m *Tag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tag) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tag) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != 0 {
		i = encodeVarintTag(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x30
	}
	if m.CreatedAt != 0 {
		i = encodeVarintTag(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Sha) > 0 {
		i -= len(m.Sha)
		copy(dAtA[i:], m.Sha)
		i = encodeVarintTag(dAtA, i, uint64(len(m.Sha)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTag(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RepositoryId != 0 {
		i = encodeVarintTag(dAtA, i, uint64(m.RepositoryId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintTag(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTag(dAtA []byte, offset int, v uint64) int {
	offset -= sovTag(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tag) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTag(uint64(m.Id))
	}
	if m.RepositoryId != 0 {
		n += 1 + sovTag(uint64(m.RepositoryId))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTag(uint64(l))
	}
	l = len(m.Sha)
	if l > 0 {
		n += 1 + l + sovTag(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovTag(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovTag(uint64(m.UpdatedAt))
	}
	return n
}

func sovTag(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTag(x uint64) (n int) {
	return sovTag(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTag
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
			return fmt.Errorf("proto: Tag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepositoryId", wireType)
			}
			m.RepositoryId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RepositoryId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTag
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
				return ErrInvalidLengthTag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTag
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
				return ErrInvalidLengthTag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTag
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
func skipTag(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTag
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
					return 0, ErrIntOverflowTag
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
					return 0, ErrIntOverflowTag
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
				return 0, ErrInvalidLengthTag
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTag
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTag
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTag        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTag          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTag = fmt.Errorf("proto: unexpected end of group")
)
