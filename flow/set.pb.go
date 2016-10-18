// Code generated by protoc-gen-go.
// source: flow/set.proto
// DO NOT EDIT!

package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FlowSet struct {
	Flows []*Flow `protobuf:"bytes,1,rep,name=Flows" json:"Flows,omitempty"`
	Start int64   `protobuf:"varint,2,opt,name=Start" json:"Start,omitempty"`
	End   int64   `protobuf:"varint,3,opt,name=End" json:"End,omitempty"`
}

func (m *FlowSet) Reset()                    { *m = FlowSet{} }
func (m *FlowSet) String() string            { return proto.CompactTextString(m) }
func (*FlowSet) ProtoMessage()               {}
func (*FlowSet) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *FlowSet) GetFlows() []*Flow {
	if m != nil {
		return m.Flows
	}
	return nil
}

func init() {
	proto.RegisterType((*FlowSet)(nil), "flow.FlowSet")
}

var fileDescriptor1 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xc9, 0x2f,
	0xd7, 0x2f, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xf1, 0xa5, 0xf8,
	0xc1, 0xa2, 0x20, 0x02, 0x22, 0xac, 0x14, 0xcc, 0xc5, 0xee, 0x06, 0xe4, 0x05, 0xa7, 0x96, 0x08,
	0x29, 0x70, 0xb1, 0x82, 0x98, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x5c, 0x7a, 0x60,
	0x65, 0x20, 0xa1, 0x20, 0x88, 0x84, 0x90, 0x08, 0x17, 0x6b, 0x70, 0x49, 0x62, 0x51, 0x89, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x84, 0x23, 0x24, 0xc0, 0xc5, 0xec, 0x9a, 0x97, 0x22, 0xc1,
	0x0c, 0x16, 0x03, 0x31, 0x93, 0xd8, 0xc0, 0x66, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e,
	0x21, 0x3b, 0x84, 0x84, 0x00, 0x00, 0x00,
}
