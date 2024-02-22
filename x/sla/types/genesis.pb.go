// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/sla/v1/genesis.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/skip-mev/slinky/pkg/types"
	_ "github.com/skip-mev/slinky/x/oracle/types"
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

// GenesisState defines the sla module's genesis state.
type GenesisState struct {
	// SLAs are the SLAs that are currently active.
	SLAs []PriceFeedSLA `protobuf:"bytes,1,rep,name=slas,proto3" json:"slas"`
	// PrceFeeds are the price feeds that are currently active.
	PriceFeeds []PriceFeed `protobuf:"bytes,2,rep,name=price_feeds,json=priceFeeds,proto3" json:"price_feeds"`
	// Params are the parameters for the sla module.
	Params Params `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_017e50c7677a1cf4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetSLAs() []PriceFeedSLA {
	if m != nil {
		return m.SLAs
	}
	return nil
}

func (m *GenesisState) GetPriceFeeds() []PriceFeed {
	if m != nil {
		return m.PriceFeeds
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// Params defines the parameters for the sla module.
type Params struct {
	// Enabled is a flag to enable or disable the sla module.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_017e50c7677a1cf4, []int{1}
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

func (m *Params) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// PriceFeedSLA defines the the desired SLA for a given set of price feeds. A
// price feed is defined to be a set of price prices for the same (currency
// pair, validator).
type PriceFeedSLA struct {
	// MaximumViableWindow is the maximum time window that we are interested
	// for the SLA. This is used to determine the moving window of blocks that
	// we are interested in.
	MaximumViableWindow uint64 `protobuf:"varint,1,opt,name=maximum_viable_window,json=maximumViableWindow,proto3" json:"maximum_viable_window,omitempty"`
	// ExpectedUptime is the expected uptime for the given validator and price
	// feed.
	ExpectedUptime cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=expected_uptime,json=expectedUptime,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"expected_uptime"`
	// SlashConstant is the constant by which we will multiply the deviation from
	// the expected uptime.
	SlashConstant cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=slash_constant,json=slashConstant,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_constant"`
	// MinimumBlockUpdates is the minimum number of blocks that the
	// validator had to have voted on in the maximum viable window
	// in order to be considered for the SLA.
	MinimumBlockUpdates uint64 `protobuf:"varint,4,opt,name=minimum_block_updates,json=minimumBlockUpdates,proto3" json:"minimum_block_updates,omitempty"`
	// Frequency is the frequency at which we will check the SLA.
	Frequency uint64 `protobuf:"varint,5,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// ID is the unique identifier for the SLA.
	ID string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *PriceFeedSLA) Reset()         { *m = PriceFeedSLA{} }
func (m *PriceFeedSLA) String() string { return proto.CompactTextString(m) }
func (*PriceFeedSLA) ProtoMessage()    {}
func (*PriceFeedSLA) Descriptor() ([]byte, []int) {
	return fileDescriptor_017e50c7677a1cf4, []int{2}
}
func (m *PriceFeedSLA) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceFeedSLA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceFeedSLA.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceFeedSLA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceFeedSLA.Merge(m, src)
}
func (m *PriceFeedSLA) XXX_Size() int {
	return m.Size()
}
func (m *PriceFeedSLA) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceFeedSLA.DiscardUnknown(m)
}

var xxx_messageInfo_PriceFeedSLA proto.InternalMessageInfo

func (m *PriceFeedSLA) GetMaximumViableWindow() uint64 {
	if m != nil {
		return m.MaximumViableWindow
	}
	return 0
}

func (m *PriceFeedSLA) GetMinimumBlockUpdates() uint64 {
	if m != nil {
		return m.MinimumBlockUpdates
	}
	return 0
}

func (m *PriceFeedSLA) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *PriceFeedSLA) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

// PriceFeed defines the object type that will be utilized to monitor how
// frequently validators are voting with price updates across the network.
type PriceFeed struct {
	// UpdateMap represents the relevant moving window of price feed updates.
	UpdateMap []byte `protobuf:"bytes,1,opt,name=update_map,json=updateMap,proto3" json:"update_map,omitempty"`
	// InclusionMap represents the relevant moving window of blocks that the
	// validator has voted on.
	InclusionMap []byte `protobuf:"bytes,2,opt,name=inclusion_map,json=inclusionMap,proto3" json:"inclusion_map,omitempty"`
	// Index corresponds to the current index into the bitmap.
	Index uint64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	// Validator represents the validator that this SLA corresponds to.
	Validator []byte `protobuf:"bytes,4,opt,name=validator,proto3" json:"validator,omitempty"`
	// CurrencyPair represents the currency pair that this SLA corresponds to.
	CurrencyPair types.CurrencyPair `protobuf:"bytes,5,opt,name=currency_pair,json=currencyPair,proto3" json:"currency_pair"`
	// MaximumViableWindow represents the maximum number of blocks that can be
	// represented by the bit map.
	MaximumViableWindow uint64 `protobuf:"varint,6,opt,name=maximum_viable_window,json=maximumViableWindow,proto3" json:"maximum_viable_window,omitempty"`
	// ID corresponds to the SLA ID that this price feed corresponds to.
	ID string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *PriceFeed) Reset()         { *m = PriceFeed{} }
func (m *PriceFeed) String() string { return proto.CompactTextString(m) }
func (*PriceFeed) ProtoMessage()    {}
func (*PriceFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_017e50c7677a1cf4, []int{3}
}
func (m *PriceFeed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceFeed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceFeed.Merge(m, src)
}
func (m *PriceFeed) XXX_Size() int {
	return m.Size()
}
func (m *PriceFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceFeed.DiscardUnknown(m)
}

var xxx_messageInfo_PriceFeed proto.InternalMessageInfo

func (m *PriceFeed) GetUpdateMap() []byte {
	if m != nil {
		return m.UpdateMap
	}
	return nil
}

func (m *PriceFeed) GetInclusionMap() []byte {
	if m != nil {
		return m.InclusionMap
	}
	return nil
}

func (m *PriceFeed) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PriceFeed) GetValidator() []byte {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *PriceFeed) GetCurrencyPair() types.CurrencyPair {
	if m != nil {
		return m.CurrencyPair
	}
	return types.CurrencyPair{}
}

func (m *PriceFeed) GetMaximumViableWindow() uint64 {
	if m != nil {
		return m.MaximumViableWindow
	}
	return 0
}

func (m *PriceFeed) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "slinky.sla.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "slinky.sla.v1.Params")
	proto.RegisterType((*PriceFeedSLA)(nil), "slinky.sla.v1.PriceFeedSLA")
	proto.RegisterType((*PriceFeed)(nil), "slinky.sla.v1.PriceFeed")
}

func init() { proto.RegisterFile("slinky/sla/v1/genesis.proto", fileDescriptor_017e50c7677a1cf4) }

var fileDescriptor_017e50c7677a1cf4 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x4e, 0xdb, 0x40,
	0x14, 0x8d, 0x4d, 0x08, 0xcd, 0xc4, 0xa1, 0xd2, 0x14, 0x2a, 0x17, 0x8a, 0x13, 0x05, 0xa9, 0xca,
	0x06, 0x5b, 0x84, 0x75, 0xd5, 0x12, 0x50, 0x1f, 0x12, 0x95, 0x90, 0x11, 0x6d, 0xc5, 0xc6, 0x9a,
	0x8c, 0x87, 0x30, 0x8a, 0xed, 0x71, 0x3d, 0x76, 0x48, 0xfe, 0xa2, 0xdb, 0xfe, 0x47, 0xff, 0xa0,
	0x1b, 0x96, 0xa8, 0xab, 0xaa, 0x8b, 0xa8, 0x32, 0x3f, 0x52, 0xf9, 0x8e, 0xc3, 0xab, 0x62, 0xd3,
	0x5d, 0xe6, 0x9e, 0x73, 0x4f, 0xee, 0x3d, 0x73, 0xc6, 0x68, 0x5d, 0x06, 0x3c, 0x1a, 0x4d, 0x1d,
	0x19, 0x10, 0x67, 0xbc, 0xed, 0x0c, 0x59, 0xc4, 0x24, 0x97, 0x76, 0x9c, 0x88, 0x54, 0xe0, 0xa6,
	0x02, 0x6d, 0x19, 0x10, 0x7b, 0xbc, 0xbd, 0xb6, 0x32, 0x14, 0x43, 0x01, 0x88, 0x53, 0xfc, 0x52,
	0xa4, 0xb5, 0x67, 0x54, 0xc8, 0x50, 0x48, 0x4f, 0x01, 0xea, 0x50, 0x42, 0x56, 0x29, 0x2e, 0x12,
	0x42, 0x03, 0xf6, 0x8f, 0xfe, 0xda, 0x66, 0x89, 0xa7, 0xd3, 0x98, 0xc9, 0x02, 0xa6, 0x59, 0x92,
	0xb0, 0x88, 0x4e, 0xbd, 0x98, 0xf0, 0x44, 0x91, 0x3a, 0x3f, 0x34, 0x64, 0xbc, 0x55, 0x6d, 0x47,
	0x29, 0x49, 0x19, 0x7e, 0x89, 0xaa, 0x32, 0x20, 0xd2, 0xd4, 0xda, 0x0b, 0xdd, 0x46, 0x6f, 0xdd,
	0xbe, 0x33, 0xa4, 0x7d, 0x98, 0x70, 0xca, 0xde, 0x30, 0xe6, 0x1f, 0x1d, 0xec, 0xf6, 0x8d, 0x8b,
	0x59, 0xab, 0x92, 0xcf, 0x5a, 0xd5, 0xa3, 0x83, 0x5d, 0xe9, 0x42, 0x1b, 0x7e, 0x85, 0x1a, 0x71,
	0xc1, 0xf1, 0x4e, 0x19, 0xf3, 0xa5, 0xa9, 0x83, 0x8a, 0xf9, 0x90, 0x4a, 0xbf, 0x5a, 0x48, 0xb8,
	0x28, 0x9e, 0x17, 0x24, 0xde, 0x41, 0xb5, 0x98, 0x24, 0x24, 0x94, 0xe6, 0x42, 0x5b, 0xeb, 0x36,
	0x7a, 0xab, 0xf7, 0x7b, 0x01, 0x2c, 0x1b, 0x4b, 0x6a, 0xa7, 0x83, 0x6a, 0xaa, 0x8e, 0x4d, 0xb4,
	0xc4, 0x22, 0x32, 0x08, 0x98, 0x6f, 0x6a, 0x6d, 0xad, 0xfb, 0xc8, 0x9d, 0x1f, 0x3b, 0xb9, 0x8e,
	0x8c, 0xdb, 0xe3, 0xe3, 0x1e, 0x5a, 0x0d, 0xc9, 0x84, 0x87, 0x59, 0xe8, 0x8d, 0x79, 0xc1, 0xf1,
	0xce, 0x79, 0xe4, 0x8b, 0x73, 0x68, 0xac, 0xba, 0x4f, 0x4a, 0xf0, 0x23, 0x60, 0x9f, 0x00, 0xc2,
	0x27, 0xe8, 0x31, 0x9b, 0xc4, 0x8c, 0xa6, 0xcc, 0xf7, 0xb2, 0x38, 0xe5, 0x21, 0x33, 0xf5, 0xb6,
	0xd6, 0xad, 0xf7, 0xb7, 0x8b, 0x79, 0x7e, 0xcf, 0x5a, 0xeb, 0xea, 0x8a, 0xa4, 0x3f, 0xb2, 0xb9,
	0x70, 0x42, 0x92, 0x9e, 0xd9, 0x07, 0x6c, 0x48, 0xe8, 0x74, 0x9f, 0xd1, 0x9f, 0xdf, 0xb7, 0x50,
	0x79, 0x83, 0xfb, 0x8c, 0xba, 0xcb, 0x73, 0xa5, 0x63, 0x10, 0xc2, 0x9f, 0xd1, 0x72, 0x61, 0xe1,
	0x99, 0x47, 0x45, 0x24, 0x53, 0x12, 0xa5, 0xe0, 0xc0, 0x7f, 0x49, 0x37, 0x41, 0x68, 0xaf, 0xd4,
	0x81, 0x4d, 0x79, 0x04, 0x9b, 0x0e, 0x02, 0x41, 0x47, 0x5e, 0x16, 0xfb, 0x24, 0x65, 0xd2, 0xac,
	0x96, 0x9b, 0x2a, 0xb0, 0x5f, 0x60, 0xc7, 0x0a, 0xc2, 0xcf, 0x51, 0xfd, 0x34, 0x61, 0x5f, 0xb2,
	0x22, 0x30, 0xe6, 0x22, 0xf0, 0x6e, 0x0a, 0xf8, 0x29, 0xd2, 0xb9, 0x6f, 0xd6, 0x60, 0xbe, 0x5a,
	0x3e, 0x6b, 0xe9, 0xef, 0xf7, 0x5d, 0x9d, 0xfb, 0x9d, 0x6f, 0x3a, 0xaa, 0x5f, 0x9b, 0x8c, 0x37,
	0x10, 0x52, 0xff, 0xe4, 0x85, 0x24, 0x06, 0x5b, 0x0d, 0xb7, 0xae, 0x2a, 0x1f, 0x48, 0x8c, 0x37,
	0x51, 0x93, 0x47, 0x34, 0xc8, 0x24, 0x17, 0x11, 0x30, 0x74, 0x60, 0x18, 0xd7, 0xc5, 0x82, 0xb4,
	0x82, 0x16, 0x79, 0xe4, 0xb3, 0x09, 0x98, 0x51, 0x75, 0xd5, 0xa1, 0x98, 0x6e, 0x4c, 0x02, 0xee,
	0x93, 0x54, 0x24, 0xb0, 0x85, 0xe1, 0xde, 0x14, 0xf0, 0x3b, 0xd4, 0xbc, 0x93, 0x75, 0x98, 0xbf,
	0xd1, 0xdb, 0x98, 0x47, 0x09, 0x5e, 0x44, 0x11, 0xa6, 0xbd, 0x92, 0x75, 0x48, 0x78, 0x52, 0x46,
	0xca, 0xa0, 0xb7, 0x6a, 0x0f, 0x67, 0xa4, 0xf6, 0x70, 0x46, 0x94, 0x37, 0x4b, 0xf7, 0xbd, 0xe9,
	0xbf, 0xbe, 0xc8, 0x2d, 0xed, 0x32, 0xb7, 0xb4, 0x3f, 0xb9, 0xa5, 0x7d, 0xbd, 0xb2, 0x2a, 0x97,
	0x57, 0x56, 0xe5, 0xd7, 0x95, 0x55, 0x39, 0x79, 0x31, 0xe4, 0xe9, 0x59, 0x36, 0xb0, 0xa9, 0x08,
	0x1d, 0x39, 0xe2, 0xf1, 0x56, 0xc8, 0xc6, 0x4e, 0xf9, 0x7a, 0x27, 0xf0, 0xf1, 0x80, 0x89, 0x07,
	0x35, 0x78, 0xb3, 0x3b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x17, 0xc9, 0x21, 0x24, 0x57, 0x04,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.PriceFeeds) > 0 {
		for iNdEx := len(m.PriceFeeds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PriceFeeds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.SLAs) > 0 {
		for iNdEx := len(m.SLAs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SLAs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PriceFeedSLA) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceFeedSLA) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceFeedSLA) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x32
	}
	if m.Frequency != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Frequency))
		i--
		dAtA[i] = 0x28
	}
	if m.MinimumBlockUpdates != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MinimumBlockUpdates))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.SlashConstant.Size()
		i -= size
		if _, err := m.SlashConstant.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.ExpectedUptime.Size()
		i -= size
		if _, err := m.ExpectedUptime.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.MaximumViableWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaximumViableWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PriceFeed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceFeed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceFeed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x3a
	}
	if m.MaximumViableWindow != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaximumViableWindow))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.CurrencyPair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x22
	}
	if m.Index != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if len(m.InclusionMap) > 0 {
		i -= len(m.InclusionMap)
		copy(dAtA[i:], m.InclusionMap)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.InclusionMap)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UpdateMap) > 0 {
		i -= len(m.UpdateMap)
		copy(dAtA[i:], m.UpdateMap)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.UpdateMap)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SLAs) > 0 {
		for _, e := range m.SLAs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PriceFeeds) > 0 {
		for _, e := range m.PriceFeeds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *PriceFeedSLA) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaximumViableWindow != 0 {
		n += 1 + sovGenesis(uint64(m.MaximumViableWindow))
	}
	l = m.ExpectedUptime.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SlashConstant.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.MinimumBlockUpdates != 0 {
		n += 1 + sovGenesis(uint64(m.MinimumBlockUpdates))
	}
	if m.Frequency != 0 {
		n += 1 + sovGenesis(uint64(m.Frequency))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *PriceFeed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UpdateMap)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.InclusionMap)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovGenesis(uint64(m.Index))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.CurrencyPair.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.MaximumViableWindow != 0 {
		n += 1 + sovGenesis(uint64(m.MaximumViableWindow))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SLAs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SLAs = append(m.SLAs, PriceFeedSLA{})
			if err := m.SLAs[len(m.SLAs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceFeeds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceFeeds = append(m.PriceFeeds, PriceFeed{})
			if err := m.PriceFeeds[len(m.PriceFeeds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
				return ErrIntOverflowGenesis
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *PriceFeedSLA) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: PriceFeedSLA: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceFeedSLA: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumViableWindow", wireType)
			}
			m.MaximumViableWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaximumViableWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedUptime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExpectedUptime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashConstant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashConstant.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumBlockUpdates", wireType)
			}
			m.MinimumBlockUpdates = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinimumBlockUpdates |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frequency", wireType)
			}
			m.Frequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frequency |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *PriceFeed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: PriceFeed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceFeed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateMap", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdateMap = append(m.UpdateMap[:0], dAtA[iNdEx:postIndex]...)
			if m.UpdateMap == nil {
				m.UpdateMap = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InclusionMap", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InclusionMap = append(m.InclusionMap[:0], dAtA[iNdEx:postIndex]...)
			if m.InclusionMap == nil {
				m.InclusionMap = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = append(m.Validator[:0], dAtA[iNdEx:postIndex]...)
			if m.Validator == nil {
				m.Validator = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrencyPair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumViableWindow", wireType)
			}
			m.MaximumViableWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaximumViableWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
