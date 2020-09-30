// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: random/random.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_tendermint_tendermint_libs_bytes "github.com/tendermint/tendermint/libs/bytes"
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

// MsgRequestRandom defines an sdk.Msg type that supports requesting a random number
type MsgRequestRandom struct {
	BlockInterval uint64                                        `protobuf:"varint,1,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty" yaml:"block_interval"`
	Consumer      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=consumer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"consumer,omitempty"`
	Oracle        bool                                          `protobuf:"varint,3,opt,name=oracle,proto3" json:"oracle,omitempty"`
	ServiceFeeCap github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,4,rep,name=service_fee_cap,json=serviceFeeCap,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"service_fee_cap" yaml:"service_fee_cap"`
}

func (m *MsgRequestRandom) Reset()         { *m = MsgRequestRandom{} }
func (m *MsgRequestRandom) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRandom) ProtoMessage()    {}
func (*MsgRequestRandom) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5da2919a686585f, []int{0}
}
func (m *MsgRequestRandom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRandom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRandom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRandom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRandom.Merge(m, src)
}
func (m *MsgRequestRandom) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRandom) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRandom.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRandom proto.InternalMessageInfo

func (m *MsgRequestRandom) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

func (m *MsgRequestRandom) GetConsumer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Consumer
	}
	return nil
}

func (m *MsgRequestRandom) GetOracle() bool {
	if m != nil {
		return m.Oracle
	}
	return false
}

func (m *MsgRequestRandom) GetServiceFeeCap() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ServiceFeeCap
	}
	return nil
}

// Random defines the feed standard
type Random struct {
	RequestTxHash github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=request_tx_hash,json=requestTxHash,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"request_tx_hash,omitempty" yaml:"request_tx_hash"`
	Height        int64                                                `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Value         string                                               `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Random) Reset()         { *m = Random{} }
func (m *Random) String() string { return proto.CompactTextString(m) }
func (*Random) ProtoMessage()    {}
func (*Random) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5da2919a686585f, []int{1}
}
func (m *Random) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Random) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Random.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Random) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Random.Merge(m, src)
}
func (m *Random) XXX_Size() int {
	return m.Size()
}
func (m *Random) XXX_DiscardUnknown() {
	xxx_messageInfo_Random.DiscardUnknown(m)
}

var xxx_messageInfo_Random proto.InternalMessageInfo

func (m *Random) GetRequestTxHash() github_com_tendermint_tendermint_libs_bytes.HexBytes {
	if m != nil {
		return m.RequestTxHash
	}
	return nil
}

func (m *Random) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Random) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Request defines the random request standard
type Request struct {
	Height           int64                                                `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Consumer         github_com_cosmos_cosmos_sdk_types.AccAddress        `protobuf:"bytes,2,opt,name=consumer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"consumer,omitempty"`
	TxHash           github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,3,opt,name=tx_hash,json=txHash,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"tx_hash,omitempty" yaml:"tx_hash"`
	Oracle           bool                                                 `protobuf:"varint,4,opt,name=oracle,proto3" json:"oracle,omitempty"`
	ServiceFeeCap    github_com_cosmos_cosmos_sdk_types.Coins             `protobuf:"bytes,5,rep,name=service_fee_cap,json=serviceFeeCap,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"service_fee_cap" yaml:"service_fee_cap"`
	ServiceContextID github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,6,opt,name=service_context_id,json=serviceContextId,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"service_context_id,omitempty" yaml:"service_context_id"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5da2919a686585f, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Request) GetConsumer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Consumer
	}
	return nil
}

func (m *Request) GetTxHash() github_com_tendermint_tendermint_libs_bytes.HexBytes {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Request) GetOracle() bool {
	if m != nil {
		return m.Oracle
	}
	return false
}

func (m *Request) GetServiceFeeCap() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ServiceFeeCap
	}
	return nil
}

func (m *Request) GetServiceContextID() github_com_tendermint_tendermint_libs_bytes.HexBytes {
	if m != nil {
		return m.ServiceContextID
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgRequestRandom)(nil), "irismod.random.MsgRequestRandom")
	proto.RegisterType((*Random)(nil), "irismod.random.Random")
	proto.RegisterType((*Request)(nil), "irismod.random.Request")
}

func init() { proto.RegisterFile("random/random.proto", fileDescriptor_e5da2919a686585f) }

var fileDescriptor_e5da2919a686585f = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xe3, 0x26, 0x4d, 0xcb, 0xd1, 0xa4, 0x91, 0x29, 0x25, 0xe9, 0x60, 0x47, 0x9e, 0xb2,
	0xd4, 0x56, 0x80, 0x89, 0x89, 0x3a, 0x08, 0x35, 0x95, 0xba, 0x18, 0x26, 0x16, 0xeb, 0x7c, 0x7e,
	0xc4, 0xa7, 0xda, 0xbe, 0xe0, 0xbb, 0x44, 0x09, 0xdf, 0x01, 0x89, 0xad, 0xdf, 0x81, 0x91, 0x4f,
	0x51, 0x89, 0xa5, 0x23, 0x93, 0x41, 0xc9, 0x37, 0xc8, 0xd8, 0x09, 0xc5, 0x77, 0x6d, 0x93, 0x2e,
	0x20, 0x40, 0xea, 0x74, 0xf7, 0xee, 0x7f, 0xef, 0xff, 0xee, 0x7e, 0x77, 0x7a, 0xe8, 0x51, 0x86,
	0xd3, 0x90, 0x25, 0x8e, 0x1c, 0xec, 0x61, 0xc6, 0x04, 0xd3, 0xeb, 0x34, 0xa3, 0x3c, 0x61, 0xa1,
	0x2d, 0x57, 0x0f, 0xf6, 0x06, 0x6c, 0xc0, 0x0a, 0xc9, 0x59, 0xce, 0xe4, 0xae, 0x83, 0x27, 0x84,
	0xf1, 0x84, 0x71, 0x5f, 0x0a, 0x84, 0xd1, 0x54, 0x0a, 0xd6, 0xb7, 0x0d, 0xd4, 0x38, 0xe5, 0x03,
	0x0f, 0x3e, 0x8c, 0x80, 0x0b, 0xaf, 0xf0, 0xd0, 0x5f, 0xa2, 0x7a, 0x10, 0x33, 0x72, 0xe6, 0xd3,
	0x54, 0x40, 0x36, 0xc6, 0x71, 0x53, 0x6b, 0x6b, 0x9d, 0x8a, 0xdb, 0x5a, 0xe4, 0xe6, 0xe3, 0x29,
	0x4e, 0xe2, 0x17, 0xd6, 0xba, 0x6e, 0x79, 0xb5, 0x62, 0xa1, 0xaf, 0x62, 0xfd, 0x14, 0x6d, 0x13,
	0x96, 0xf2, 0x51, 0x02, 0x59, 0x73, 0xa3, 0xad, 0x75, 0x76, 0xdc, 0xee, 0x55, 0x6e, 0x1e, 0x0e,
	0xa8, 0x88, 0x46, 0x81, 0x4d, 0x58, 0xe2, 0xc8, 0x03, 0xa9, 0xe1, 0x90, 0x87, 0x67, 0x8e, 0x98,
	0x0e, 0x81, 0xdb, 0x47, 0x84, 0x1c, 0x85, 0x61, 0x06, 0x9c, 0x7b, 0x37, 0x16, 0xfa, 0x3e, 0xaa,
	0xb2, 0x0c, 0x93, 0x18, 0x9a, 0xe5, 0xb6, 0xd6, 0xd9, 0xf6, 0x54, 0xa4, 0x7f, 0xd2, 0xd0, 0x2e,
	0x87, 0x6c, 0x4c, 0x09, 0xf8, 0xef, 0x01, 0x7c, 0x82, 0x87, 0xcd, 0x4a, 0xbb, 0xdc, 0x79, 0xf8,
	0xb4, 0x65, 0x4b, 0x67, 0x3b, 0xc0, 0x1c, 0xec, 0x71, 0x37, 0x00, 0x81, 0xbb, 0x76, 0x8f, 0xd1,
	0xd4, 0x3d, 0xb9, 0xc8, 0xcd, 0xd2, 0x22, 0x37, 0xf7, 0xe5, 0x4d, 0xee, 0xe4, 0x5b, 0x5f, 0x7e,
	0x98, 0x9d, 0x3f, 0x38, 0xe7, 0xd2, 0x8a, 0x7b, 0x35, 0x95, 0xfd, 0x1a, 0xa0, 0x87, 0x87, 0xd6,
	0x57, 0x0d, 0x55, 0x15, 0xc3, 0x8f, 0x68, 0x37, 0x93, 0x50, 0x7d, 0x31, 0xf1, 0x23, 0xcc, 0xa3,
	0x02, 0xe2, 0x8e, 0xeb, 0xdd, 0x96, 0xbe, 0xb3, 0xc1, 0xba, 0xca, 0xcd, 0xe7, 0x2b, 0xa5, 0x05,
	0xa4, 0x21, 0x64, 0x09, 0x4d, 0xc5, 0xea, 0x34, 0xa6, 0x01, 0x77, 0x82, 0xa9, 0x00, 0x6e, 0x1f,
	0xc3, 0xc4, 0x5d, 0x4e, 0xbc, 0x9a, 0x72, 0x7a, 0x3b, 0x39, 0xc6, 0x3c, 0x5a, 0xe2, 0x8a, 0x80,
	0x0e, 0x22, 0x51, 0xb0, 0x2f, 0x7b, 0x2a, 0xd2, 0xf7, 0xd0, 0xe6, 0x18, 0xc7, 0x23, 0x49, 0xf1,
	0x81, 0x27, 0x03, 0xeb, 0xbc, 0x82, 0xb6, 0xd4, 0xfb, 0xaf, 0x64, 0x6a, 0x6b, 0x99, 0xff, 0xf9,
	0x3d, 0x09, 0xda, 0xba, 0x86, 0x52, 0x2e, 0xdc, 0x4e, 0x16, 0xb9, 0x59, 0x97, 0x50, 0xfe, 0x19,
	0x46, 0x55, 0xdc, 0x50, 0x50, 0x9f, 0xa6, 0xf2, 0xdb, 0x4f, 0xb3, 0x79, 0x7f, 0x9f, 0x46, 0x3f,
	0xd7, 0x90, 0x7e, 0xed, 0x47, 0x58, 0x2a, 0x60, 0x22, 0x7c, 0x1a, 0x36, 0xab, 0x05, 0x18, 0x3a,
	0xcb, 0xcd, 0xc6, 0x1b, 0xa9, 0xf6, 0xa4, 0xd8, 0x7f, 0xb5, 0xc8, 0xcd, 0xd6, 0xfa, 0x39, 0x6e,
	0xf3, 0xfe, 0x9e, 0x5b, 0x83, 0xaf, 0x97, 0x09, 0xdd, 0xfe, 0xc5, 0xcc, 0xd0, 0x2e, 0x67, 0x86,
	0xf6, 0x73, 0x66, 0x68, 0x9f, 0xe7, 0x46, 0xe9, 0x72, 0x6e, 0x94, 0xbe, 0xcf, 0x8d, 0xd2, 0x3b,
	0x67, 0xa5, 0xc2, 0xb2, 0x01, 0xa5, 0x20, 0x1c, 0xd5, 0x88, 0x9c, 0x84, 0x85, 0xa3, 0x18, 0xb8,
	0x6a, 0x53, 0xf2, 0xe6, 0x41, 0xb5, 0x68, 0x37, 0xcf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x97,
	0x91, 0xbf, 0x0f, 0xc4, 0x04, 0x00, 0x00,
}

func (m *MsgRequestRandom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRandom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRandom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceFeeCap) > 0 {
		for iNdEx := len(m.ServiceFeeCap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceFeeCap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRandom(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Oracle {
		i--
		if m.Oracle {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x12
	}
	if m.BlockInterval != 0 {
		i = encodeVarintRandom(dAtA, i, uint64(m.BlockInterval))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Random) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Random) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Random) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Height != 0 {
		i = encodeVarintRandom(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RequestTxHash) > 0 {
		i -= len(m.RequestTxHash)
		copy(dAtA[i:], m.RequestTxHash)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.RequestTxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceContextID) > 0 {
		i -= len(m.ServiceContextID)
		copy(dAtA[i:], m.ServiceContextID)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.ServiceContextID)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ServiceFeeCap) > 0 {
		for iNdEx := len(m.ServiceFeeCap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceFeeCap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRandom(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Oracle {
		i--
		if m.Oracle {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintRandom(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRandom(dAtA []byte, offset int, v uint64) int {
	offset -= sovRandom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRequestRandom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockInterval != 0 {
		n += 1 + sovRandom(uint64(m.BlockInterval))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	if m.Oracle {
		n += 2
	}
	if len(m.ServiceFeeCap) > 0 {
		for _, e := range m.ServiceFeeCap {
			l = e.Size()
			n += 1 + l + sovRandom(uint64(l))
		}
	}
	return n
}

func (m *Random) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestTxHash)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovRandom(uint64(m.Height))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovRandom(uint64(m.Height))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	if m.Oracle {
		n += 2
	}
	if len(m.ServiceFeeCap) > 0 {
		for _, e := range m.ServiceFeeCap {
			l = e.Size()
			n += 1 + l + sovRandom(uint64(l))
		}
	}
	l = len(m.ServiceContextID)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	return n
}

func sovRandom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRandom(x uint64) (n int) {
	return sovRandom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRequestRandom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandom
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
			return fmt.Errorf("proto: MsgRequestRandom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRandom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockInterval", wireType)
			}
			m.BlockInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Consumer = append(m.Consumer[:0], dAtA[iNdEx:postIndex]...)
			if m.Consumer == nil {
				m.Consumer = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Oracle = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceFeeCap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceFeeCap = append(m.ServiceFeeCap, types.Coin{})
			if err := m.ServiceFeeCap[len(m.ServiceFeeCap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRandom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRandom
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRandom
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
func (m *Random) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandom
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
			return fmt.Errorf("proto: Random: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Random: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestTxHash = append(m.RequestTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestTxHash == nil {
				m.RequestTxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRandom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRandom
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRandom
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
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandom
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Consumer = append(m.Consumer[:0], dAtA[iNdEx:postIndex]...)
			if m.Consumer == nil {
				m.Consumer = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Oracle = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceFeeCap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceFeeCap = append(m.ServiceFeeCap, types.Coin{})
			if err := m.ServiceFeeCap[len(m.ServiceFeeCap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceContextID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceContextID = append(m.ServiceContextID[:0], dAtA[iNdEx:postIndex]...)
			if m.ServiceContextID == nil {
				m.ServiceContextID = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRandom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRandom
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRandom
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
func skipRandom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRandom
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
					return 0, ErrIntOverflowRandom
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
					return 0, ErrIntOverflowRandom
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
				return 0, ErrInvalidLengthRandom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRandom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRandom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRandom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRandom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRandom = fmt.Errorf("proto: unexpected end of group")
)