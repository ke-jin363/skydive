// Code generated by protoc-gen-go.
// source: flow/request.proto
// DO NOT EDIT!

package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BoolFilterOp int32

const (
	BoolFilterOp_OR  BoolFilterOp = 0
	BoolFilterOp_AND BoolFilterOp = 1
	BoolFilterOp_NOT BoolFilterOp = 2
)

var BoolFilterOp_name = map[int32]string{
	0: "OR",
	1: "AND",
	2: "NOT",
}
var BoolFilterOp_value = map[string]int32{
	"OR":  0,
	"AND": 1,
	"NOT": 2,
}

func (x BoolFilterOp) String() string {
	return proto.EnumName(BoolFilterOp_name, int32(x))
}
func (BoolFilterOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type TermStringFilter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *TermStringFilter) Reset()                    { *m = TermStringFilter{} }
func (m *TermStringFilter) String() string            { return proto.CompactTextString(m) }
func (*TermStringFilter) ProtoMessage()               {}
func (*TermStringFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type TermInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *TermInt64Filter) Reset()                    { *m = TermInt64Filter{} }
func (m *TermInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*TermInt64Filter) ProtoMessage()               {}
func (*TermInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type NeStringFilter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *NeStringFilter) Reset()                    { *m = NeStringFilter{} }
func (m *NeStringFilter) String() string            { return proto.CompactTextString(m) }
func (*NeStringFilter) ProtoMessage()               {}
func (*NeStringFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type NeInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *NeInt64Filter) Reset()                    { *m = NeInt64Filter{} }
func (m *NeInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*NeInt64Filter) ProtoMessage()               {}
func (*NeInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type GtInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *GtInt64Filter) Reset()                    { *m = GtInt64Filter{} }
func (m *GtInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*GtInt64Filter) ProtoMessage()               {}
func (*GtInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type LtInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *LtInt64Filter) Reset()                    { *m = LtInt64Filter{} }
func (m *LtInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*LtInt64Filter) ProtoMessage()               {}
func (*LtInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type GteInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *GteInt64Filter) Reset()                    { *m = GteInt64Filter{} }
func (m *GteInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*GteInt64Filter) ProtoMessage()               {}
func (*GteInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type LteInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *LteInt64Filter) Reset()                    { *m = LteInt64Filter{} }
func (m *LteInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*LteInt64Filter) ProtoMessage()               {}
func (*LteInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type Filter struct {
	TermStringFilter *TermStringFilter `protobuf:"bytes,1,opt,name=TermStringFilter" json:"TermStringFilter,omitempty"`
	TermInt64Filter  *TermInt64Filter  `protobuf:"bytes,2,opt,name=TermInt64Filter" json:"TermInt64Filter,omitempty"`
	GtInt64Filter    *GtInt64Filter    `protobuf:"bytes,3,opt,name=GtInt64Filter" json:"GtInt64Filter,omitempty"`
	LtInt64Filter    *LtInt64Filter    `protobuf:"bytes,4,opt,name=LtInt64Filter" json:"LtInt64Filter,omitempty"`
	GteInt64Filter   *GteInt64Filter   `protobuf:"bytes,5,opt,name=GteInt64Filter" json:"GteInt64Filter,omitempty"`
	LteInt64Filter   *LteInt64Filter   `protobuf:"bytes,6,opt,name=LteInt64Filter" json:"LteInt64Filter,omitempty"`
	BoolFilter       *BoolFilter       `protobuf:"bytes,7,opt,name=BoolFilter" json:"BoolFilter,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *Filter) GetTermStringFilter() *TermStringFilter {
	if m != nil {
		return m.TermStringFilter
	}
	return nil
}

func (m *Filter) GetTermInt64Filter() *TermInt64Filter {
	if m != nil {
		return m.TermInt64Filter
	}
	return nil
}

func (m *Filter) GetGtInt64Filter() *GtInt64Filter {
	if m != nil {
		return m.GtInt64Filter
	}
	return nil
}

func (m *Filter) GetLtInt64Filter() *LtInt64Filter {
	if m != nil {
		return m.LtInt64Filter
	}
	return nil
}

func (m *Filter) GetGteInt64Filter() *GteInt64Filter {
	if m != nil {
		return m.GteInt64Filter
	}
	return nil
}

func (m *Filter) GetLteInt64Filter() *LteInt64Filter {
	if m != nil {
		return m.LteInt64Filter
	}
	return nil
}

func (m *Filter) GetBoolFilter() *BoolFilter {
	if m != nil {
		return m.BoolFilter
	}
	return nil
}

type BoolFilter struct {
	Op      BoolFilterOp `protobuf:"varint,1,opt,name=Op,enum=flow.BoolFilterOp" json:"Op,omitempty"`
	Filters []*Filter    `protobuf:"bytes,2,rep,name=Filters" json:"Filters,omitempty"`
}

func (m *BoolFilter) Reset()                    { *m = BoolFilter{} }
func (m *BoolFilter) String() string            { return proto.CompactTextString(m) }
func (*BoolFilter) ProtoMessage()               {}
func (*BoolFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *BoolFilter) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

type Range struct {
	From int64 `protobuf:"varint,1,opt,name=From" json:"From,omitempty"`
	To   int64 `protobuf:"varint,2,opt,name=To" json:"To,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

type FlowSearchQuery struct {
	Filter *Filter `protobuf:"bytes,1,opt,name=Filter" json:"Filter,omitempty"`
	Range  *Range  `protobuf:"bytes,2,opt,name=Range" json:"Range,omitempty"`
	Sort   bool    `protobuf:"varint,3,opt,name=Sort" json:"Sort,omitempty"`
	Dedup  bool    `protobuf:"varint,4,opt,name=Dedup" json:"Dedup,omitempty"`
}

func (m *FlowSearchQuery) Reset()                    { *m = FlowSearchQuery{} }
func (m *FlowSearchQuery) String() string            { return proto.CompactTextString(m) }
func (*FlowSearchQuery) ProtoMessage()               {}
func (*FlowSearchQuery) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *FlowSearchQuery) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *FlowSearchQuery) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

type FlowSearchReply struct {
	FlowSet *FlowSet `protobuf:"bytes,1,opt,name=FlowSet" json:"FlowSet,omitempty"`
}

func (m *FlowSearchReply) Reset()                    { *m = FlowSearchReply{} }
func (m *FlowSearchReply) String() string            { return proto.CompactTextString(m) }
func (*FlowSearchReply) ProtoMessage()               {}
func (*FlowSearchReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *FlowSearchReply) GetFlowSet() *FlowSet {
	if m != nil {
		return m.FlowSet
	}
	return nil
}

func init() {
	proto.RegisterType((*TermStringFilter)(nil), "flow.TermStringFilter")
	proto.RegisterType((*TermInt64Filter)(nil), "flow.TermInt64Filter")
	proto.RegisterType((*NeStringFilter)(nil), "flow.NeStringFilter")
	proto.RegisterType((*NeInt64Filter)(nil), "flow.NeInt64Filter")
	proto.RegisterType((*GtInt64Filter)(nil), "flow.GtInt64Filter")
	proto.RegisterType((*LtInt64Filter)(nil), "flow.LtInt64Filter")
	proto.RegisterType((*GteInt64Filter)(nil), "flow.GteInt64Filter")
	proto.RegisterType((*LteInt64Filter)(nil), "flow.LteInt64Filter")
	proto.RegisterType((*Filter)(nil), "flow.Filter")
	proto.RegisterType((*BoolFilter)(nil), "flow.BoolFilter")
	proto.RegisterType((*Range)(nil), "flow.Range")
	proto.RegisterType((*FlowSearchQuery)(nil), "flow.FlowSearchQuery")
	proto.RegisterType((*FlowSearchReply)(nil), "flow.FlowSearchReply")
	proto.RegisterEnum("flow.BoolFilterOp", BoolFilterOp_name, BoolFilterOp_value)
}

var fileDescriptor2 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0x51, 0x6f, 0xd3, 0x30,
	0x14, 0x85, 0x69, 0xd3, 0xa6, 0xe3, 0x76, 0xcb, 0x22, 0x33, 0x50, 0xc4, 0x13, 0x44, 0x08, 0x26,
	0x90, 0x0a, 0x2a, 0x08, 0x28, 0x9a, 0x84, 0x98, 0xa6, 0x4e, 0x88, 0x28, 0x15, 0x6e, 0x35, 0xf1,
	0x5a, 0xc0, 0x8c, 0x49, 0x59, 0x1d, 0x5c, 0x07, 0xd4, 0x77, 0xfe, 0x10, 0xff, 0x10, 0xe7, 0xda,
	0xa1, 0x8e, 0xf3, 0x50, 0x94, 0x37, 0xfb, 0x5e, 0x7f, 0x3e, 0xce, 0x3d, 0x47, 0x01, 0xf2, 0x2d,
	0xe3, 0xbf, 0x9e, 0x0a, 0xf6, 0xa3, 0x60, 0x6b, 0x39, 0xca, 0x05, 0x97, 0x9c, 0xf4, 0xca, 0xda,
	0xdd, 0x00, 0x3b, 0x6b, 0x66, 0xaa, 0xf1, 0x1b, 0x08, 0x17, 0x4c, 0x5c, 0xcf, 0xa5, 0xb8, 0x5a,
	0x5d, 0x4e, 0xaf, 0x32, 0xc9, 0x04, 0x09, 0xc1, 0xfb, 0xc0, 0x36, 0x51, 0xe7, 0x5e, 0xe7, 0xf8,
	0x26, 0x2d, 0x97, 0xe4, 0x08, 0xfa, 0x3f, 0x97, 0x59, 0xc1, 0xa2, 0x2e, 0xd6, 0xf4, 0x26, 0x9e,
	0xc0, 0x61, 0xc9, 0xbe, 0x5f, 0xc9, 0x97, 0x2f, 0xfe, 0x0f, 0xf5, 0x2a, 0xf4, 0x35, 0x04, 0x29,
	0xdb, 0x2d, 0x7a, 0x61, 0x8b, 0xe2, 0x26, 0x7e, 0x05, 0x07, 0x29, 0xdb, 0x29, 0x79, 0x61, 0x4b,
	0xfe, 0x03, 0xcf, 0x65, 0x4b, 0x30, 0x69, 0x05, 0xaa, 0x8f, 0x3c, 0x97, 0xac, 0x25, 0x99, 0xb4,
	0x23, 0xff, 0x78, 0xe0, 0x1b, 0xe4, 0xb4, 0x69, 0x2d, 0xf2, 0xc3, 0xf1, 0x9d, 0x51, 0x99, 0x82,
	0x91, 0xdb, 0xa5, 0xcd, 0x28, 0xbc, 0x6d, 0x58, 0x8c, 0x72, 0xc3, 0xf1, 0xed, 0xed, 0x15, 0x56,
	0x93, 0x36, 0x02, 0x31, 0x71, 0xa6, 0x1e, 0x79, 0x88, 0xdf, 0xd2, 0x78, 0xad, 0x45, 0x1d, 0x7f,
	0x26, 0xce, 0xdc, 0xa3, 0x9e, 0x8d, 0x26, 0x75, 0xb4, 0xee, 0xd0, 0x89, 0x3b, 0xf9, 0xa8, 0x8f,
	0xec, 0x51, 0x25, 0x6b, 0xf7, 0xa8, 0xeb, 0xd2, 0x89, 0x3b, 0xfd, 0xc8, 0xb7, 0xe9, 0xc4, 0xa1,
	0x1d, 0xa7, 0x9e, 0x01, 0x9c, 0x72, 0x9e, 0x19, 0x72, 0x80, 0x64, 0xa8, 0xc9, 0x6d, 0x9d, 0x5a,
	0x67, 0xe2, 0x4f, 0x36, 0x41, 0x62, 0xe8, 0xce, 0x72, 0x34, 0x2a, 0x18, 0x13, 0x97, 0x9b, 0xe5,
	0x54, 0x75, 0xc9, 0x43, 0x18, 0xe8, 0xfd, 0x5a, 0xd9, 0xe1, 0x29, 0x81, 0x7d, 0x7d, 0xd0, 0x5c,
	0x5e, 0x35, 0xe3, 0x27, 0xd0, 0xa7, 0xcb, 0xd5, 0x25, 0x23, 0x04, 0x7a, 0x53, 0xc1, 0xaf, 0xf1,
	0x5a, 0x8f, 0xe2, 0x9a, 0x04, 0xd0, 0x5d, 0x70, 0x93, 0x1e, 0xb5, 0x8a, 0x7f, 0x77, 0xe0, 0x70,
	0xaa, 0x6e, 0x99, 0xb3, 0xa5, 0xf8, 0xf2, 0xfd, 0x63, 0xc1, 0xc4, 0x86, 0x3c, 0xa8, 0xd2, 0x64,
	0x92, 0x53, 0xd7, 0xa9, 0x92, 0x76, 0xdf, 0xc8, 0x98, 0x6c, 0x0c, 0xf5, 0x21, 0x2c, 0xd1, 0xed,
	0x03, 0xe6, 0x5c, 0x48, 0xb4, 0x7f, 0x8f, 0xe2, 0xba, 0x4c, 0xf0, 0x19, 0xfb, 0x5a, 0xe4, 0x68,
	0xec, 0x1e, 0xd5, 0x1b, 0xf5, 0x47, 0xb2, 0x5e, 0x41, 0x59, 0x9e, 0x6d, 0xc8, 0x23, 0xf5, 0xb9,
	0x58, 0x92, 0xe6, 0x19, 0x07, 0xe6, 0x19, 0xba, 0x48, 0xab, 0xee, 0xe3, 0x63, 0xd8, 0xb7, 0x67,
	0x45, 0x7c, 0x35, 0x4b, 0x1a, 0xde, 0x20, 0x03, 0xf0, 0xde, 0xa5, 0x67, 0x61, 0xa7, 0x5c, 0xa4,
	0xb3, 0x45, 0xd8, 0xfd, 0xec, 0xe3, 0xef, 0xef, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0x1d, 0xf2, 0x17, 0x2a, 0x05, 0x00, 0x00,
}
