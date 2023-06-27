// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/gitopia/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_cosmos_gogoproto_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DistributionProportion struct {
	Proportion github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=proportion,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"proportion" yaml:"proportion"`
	Address    string                                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *DistributionProportion) Reset()         { *m = DistributionProportion{} }
func (m *DistributionProportion) String() string { return proto.CompactTextString(m) }
func (*DistributionProportion) ProtoMessage()    {}
func (*DistributionProportion) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdae11692a018c3a, []int{0}
}
func (m *DistributionProportion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportion.Merge(m, src)
}
func (m *DistributionProportion) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportion) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportion.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportion proto.InternalMessageInfo

func (m *DistributionProportion) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type PoolProportions struct {
	Ecosystem *DistributionProportion `protobuf:"bytes,1,opt,name=ecosystem,proto3" json:"ecosystem,omitempty" yaml:"ecosystem"`
	Team      *DistributionProportion `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty" yaml:"team"`
}

func (m *PoolProportions) Reset()         { *m = PoolProportions{} }
func (m *PoolProportions) String() string { return proto.CompactTextString(m) }
func (*PoolProportions) ProtoMessage()    {}
func (*PoolProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdae11692a018c3a, []int{1}
}
func (m *PoolProportions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolProportions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolProportions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolProportions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolProportions.Merge(m, src)
}
func (m *PoolProportions) XXX_Size() int {
	return m.Size()
}
func (m *PoolProportions) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolProportions.DiscardUnknown(m)
}

var xxx_messageInfo_PoolProportions proto.InternalMessageInfo

func (m *PoolProportions) GetEcosystem() *DistributionProportion {
	if m != nil {
		return m.Ecosystem
	}
	return nil
}

func (m *PoolProportions) GetTeam() *DistributionProportion {
	if m != nil {
		return m.Team
	}
	return nil
}

// Params defines the parameters for the module.
type Params struct {
	NextInflationTime time.Time                `protobuf:"bytes,1,opt,name=next_inflation_time,json=nextInflationTime,proto3,stdtime" json:"next_inflation_time" yaml:"next_inflation_time"`
	PoolProportions   PoolProportions          `protobuf:"bytes,2,opt,name=pool_proportions,json=poolProportions,proto3" json:"pool_proportions" yaml:"pool_proportions"`
	TeamProportions   []DistributionProportion `protobuf:"bytes,3,rep,name=team_proportions,json=teamProportions,proto3" json:"team_proportions" yaml:"team_proportions"`
	GenesisTime       time.Time                `protobuf:"bytes,4,opt,name=genesis_time,json=genesisTime,proto3,stdtime" json:"genesis_time" yaml:"genesis_time"`
	GitServer         string                   `protobuf:"bytes,5,opt,name=git_server,json=gitServer,proto3" json:"git_server,omitempty" yaml:"git_server"`
	StorageProvider   string                   `protobuf:"bytes,6,opt,name=storage_provider,json=storageProvider,proto3" json:"storage_provider,omitempty" yaml:"storage_provider"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdae11692a018c3a, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetNextInflationTime() time.Time {
	if m != nil {
		return m.NextInflationTime
	}
	return time.Time{}
}

func (m *Params) GetPoolProportions() PoolProportions {
	if m != nil {
		return m.PoolProportions
	}
	return PoolProportions{}
}

func (m *Params) GetTeamProportions() []DistributionProportion {
	if m != nil {
		return m.TeamProportions
	}
	return nil
}

func (m *Params) GetGenesisTime() time.Time {
	if m != nil {
		return m.GenesisTime
	}
	return time.Time{}
}

func (m *Params) GetGitServer() string {
	if m != nil {
		return m.GitServer
	}
	return ""
}

func (m *Params) GetStorageProvider() string {
	if m != nil {
		return m.StorageProvider
	}
	return ""
}

func init() {
	proto.RegisterType((*DistributionProportion)(nil), "gitopia.gitopia.gitopia.DistributionProportion")
	proto.RegisterType((*PoolProportions)(nil), "gitopia.gitopia.gitopia.PoolProportions")
	proto.RegisterType((*Params)(nil), "gitopia.gitopia.gitopia.Params")
}

func init() { proto.RegisterFile("gitopia/gitopia/params.proto", fileDescriptor_cdae11692a018c3a) }

var fileDescriptor_cdae11692a018c3a = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0x6e, 0x7e, 0xeb, 0xaf, 0xa8, 0x2e, 0xa2, 0x5d, 0x56, 0x58, 0x55, 0xa4, 0x78, 0xf2, 0x61,
	0xda, 0x81, 0x25, 0xd2, 0xe0, 0xb4, 0x63, 0x99, 0x40, 0xdc, 0xaa, 0xb0, 0x13, 0x07, 0xaa, 0x34,
	0xf5, 0x8c, 0x45, 0x53, 0x47, 0xb6, 0x5b, 0xad, 0xe2, 0x4b, 0xec, 0xc8, 0x91, 0x3b, 0x5f, 0x64,
	0x07, 0x0e, 0x93, 0xb8, 0x20, 0x0e, 0x01, 0xb5, 0xdf, 0x20, 0x9f, 0x00, 0xf9, 0x4f, 0x9a, 0x12,
	0x6d, 0x42, 0xda, 0xc9, 0xf6, 0xeb, 0xf7, 0x7d, 0x9e, 0xf7, 0x7d, 0x9e, 0x38, 0xa0, 0x4b, 0xa8,
	0x64, 0x29, 0x8d, 0x82, 0x34, 0xe2, 0x51, 0x22, 0xfc, 0x94, 0x33, 0xc9, 0xdc, 0x7d, 0x1b, 0xf5,
	0x2b, 0x6b, 0xbf, 0x4b, 0x18, 0x61, 0x3a, 0x27, 0x50, 0x3b, 0x93, 0xde, 0x87, 0x84, 0x31, 0x32,
	0xc5, 0x81, 0x3e, 0x8d, 0xe7, 0x17, 0x81, 0xa4, 0x09, 0x16, 0x32, 0x4a, 0x52, 0x93, 0x80, 0xbe,
	0x3a, 0xe0, 0xc9, 0x19, 0x15, 0x92, 0xd3, 0xf1, 0x5c, 0x52, 0x36, 0x1b, 0x72, 0x96, 0x32, 0xae,
	0x76, 0x6e, 0x0c, 0x40, 0xba, 0x39, 0xf5, 0x9c, 0x03, 0xe7, 0xa8, 0x39, 0x78, 0x79, 0x9d, 0xc1,
	0xda, 0xcf, 0x0c, 0x1e, 0x12, 0x2a, 0x3f, 0xcc, 0xc7, 0x7e, 0xcc, 0x92, 0x20, 0x66, 0x22, 0x61,
	0xc2, 0x2e, 0xc7, 0x62, 0xf2, 0x31, 0x90, 0xcb, 0x14, 0x0b, 0xff, 0x0c, 0xc7, 0x79, 0x06, 0x77,
	0x97, 0x51, 0x32, 0x3d, 0x45, 0x25, 0x12, 0x0a, 0xb7, 0x60, 0xdd, 0x67, 0xe0, 0x41, 0x34, 0x99,
	0x70, 0x2c, 0x44, 0xef, 0x3f, 0xcd, 0xe0, 0xe6, 0x19, 0x7c, 0x64, 0x6a, 0xec, 0x05, 0x0a, 0x8b,
	0x14, 0xf4, 0xcd, 0x01, 0xed, 0x21, 0x63, 0xd3, 0xb2, 0x4b, 0xe1, 0xc6, 0xa0, 0x89, 0x63, 0x26,
	0x96, 0x42, 0xe2, 0x44, 0x77, 0xd9, 0x3a, 0x09, 0xfc, 0x3b, 0x54, 0xf2, 0x6f, 0x1f, 0x75, 0xd0,
	0xcd, 0x33, 0xd8, 0x31, 0xa4, 0x1b, 0x2c, 0x14, 0x96, 0xb8, 0xee, 0x39, 0xa8, 0x4b, 0x1c, 0x25,
	0xba, 0xc7, 0x7b, 0xe0, 0xb7, 0xf3, 0x0c, 0xb6, 0x0c, 0xbe, 0x82, 0x41, 0xa1, 0x46, 0x43, 0xdf,
	0xeb, 0xa0, 0x31, 0xd4, 0xee, 0xba, 0x1c, 0xec, 0xcd, 0xf0, 0xa5, 0x1c, 0xd1, 0xd9, 0xc5, 0x34,
	0x52, 0x35, 0x23, 0xe5, 0x94, 0x9d, 0xa7, 0xef, 0x1b, 0x1b, 0xfd, 0xc2, 0x46, 0xff, 0xbc, 0xb0,
	0x71, 0x70, 0xa8, 0x1c, 0xc9, 0x33, 0xd8, 0x37, 0xf0, 0xb7, 0x80, 0xa0, 0xab, 0x5f, 0xd0, 0x09,
	0x77, 0xd5, 0xcd, 0x9b, 0xe2, 0x42, 0xd5, 0xbb, 0x12, 0x74, 0x52, 0xc6, 0xa6, 0xa3, 0xd2, 0x0e,
	0x61, 0x07, 0x3c, 0xba, 0x73, 0xc0, 0x8a, 0xfa, 0x03, 0x68, 0xe9, 0xf7, 0xad, 0xcd, 0x15, 0x3c,
	0x14, 0xb6, 0xd3, 0x8a, 0x5f, 0x9f, 0x40, 0x47, 0x0d, 0xff, 0x17, 0xeb, 0xce, 0xc1, 0xce, 0x7d,
	0x64, 0xad, 0x90, 0x57, 0x61, 0x51, 0xd8, 0x56, 0xa1, 0x6d, 0xf2, 0xf7, 0xe0, 0x21, 0xc1, 0x33,
	0x2c, 0xa8, 0x30, 0xfa, 0xd6, 0xff, 0xa9, 0x6f, 0xc1, 0xb1, 0x67, 0x38, 0xb6, 0xab, 0x8d, 0xb0,
	0x2d, 0x1b, 0xd2, 0x92, 0xbe, 0x00, 0x80, 0x50, 0x39, 0x12, 0x98, 0x2f, 0x30, 0xef, 0xfd, 0xaf,
	0xbf, 0xe8, 0xc7, 0xe5, 0x2b, 0x28, 0xef, 0x50, 0xd8, 0x24, 0x54, 0xbe, 0xd5, 0x7b, 0xf7, 0x15,
	0xe8, 0x08, 0xc9, 0x78, 0x44, 0xb0, 0x6a, 0x7f, 0x41, 0x27, 0x98, 0xf7, 0x1a, 0xba, 0xf6, 0x69,
	0x39, 0x5d, 0x35, 0x03, 0x85, 0x6d, 0x1b, 0x1a, 0xda, 0xc8, 0x69, 0xfd, 0xf3, 0x17, 0x58, 0x1b,
	0xbc, 0xbe, 0x5e, 0x79, 0xce, 0xcd, 0xca, 0x73, 0x7e, 0xaf, 0x3c, 0xe7, 0x6a, 0xed, 0xd5, 0x6e,
	0xd6, 0x5e, 0xed, 0xc7, 0xda, 0xab, 0xbd, 0x3b, 0xde, 0x7a, 0xb5, 0xc5, 0xdf, 0xa5, 0x58, 0x17,
	0x27, 0xc1, 0xe5, 0xe6, 0xa0, 0x1f, 0xf0, 0xb8, 0xa1, 0xe5, 0x78, 0xfe, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x92, 0xe8, 0x6a, 0xf9, 0x8a, 0x04, 0x00, 0x00,
}

func (m *DistributionProportion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Proportion.Size()
		i -= size
		if _, err := m.Proportion.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PoolProportions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolProportions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolProportions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Team != nil {
		{
			size, err := m.Team.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Ecosystem != nil {
		{
			size, err := m.Ecosystem.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StorageProvider) > 0 {
		i -= len(m.StorageProvider)
		copy(dAtA[i:], m.StorageProvider)
		i = encodeVarintParams(dAtA, i, uint64(len(m.StorageProvider)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.GitServer) > 0 {
		i -= len(m.GitServer)
		copy(dAtA[i:], m.GitServer)
		i = encodeVarintParams(dAtA, i, uint64(len(m.GitServer)))
		i--
		dAtA[i] = 0x2a
	}
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.GenesisTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.GenesisTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	if len(m.TeamProportions) > 0 {
		for iNdEx := len(m.TeamProportions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TeamProportions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.PoolProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n5, err5 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.NextInflationTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.NextInflationTime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintParams(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DistributionProportion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Proportion.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *PoolProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ecosystem != nil {
		l = m.Ecosystem.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Team != nil {
		l = m.Team.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.NextInflationTime)
	n += 1 + l + sovParams(uint64(l))
	l = m.PoolProportions.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.TeamProportions) > 0 {
		for _, e := range m.TeamProportions {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.GenesisTime)
	n += 1 + l + sovParams(uint64(l))
	l = len(m.GitServer)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.StorageProvider)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DistributionProportion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: DistributionProportion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proportion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proportion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *PoolProportions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: PoolProportions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolProportions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ecosystem", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ecosystem == nil {
				m.Ecosystem = &DistributionProportion{}
			}
			if err := m.Ecosystem.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Team", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Team == nil {
				m.Team = &DistributionProportion{}
			}
			if err := m.Team.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextInflationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.NextInflationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeamProportions = append(m.TeamProportions, DistributionProportion{})
			if err := m.TeamProportions[len(m.TeamProportions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.GenesisTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitServer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitServer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageProvider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageProvider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
