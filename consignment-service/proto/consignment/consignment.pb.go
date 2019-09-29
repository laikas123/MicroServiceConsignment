// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

package consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//messages are custom types
//here we define our messages
//MESSAGES ARE HANDLED BY PROTOBUF
type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	Created              bool         `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*Response)(nil), "consignment.Response")
}

func init() { proto.RegisterFile("consignment.proto", fileDescriptor_3804bf87090b51a9) }

var fileDescriptor_3804bf87090b51a9 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x49, 0xff, 0xe7, 0x2c, 0x81, 0x6a, 0x89, 0x62, 0xc1, 0x40, 0x94, 0xa9, 0x53, 0x87,
	0x22, 0x31, 0xb0, 0x56, 0x42, 0xea, 0xea, 0x4e, 0x4c, 0x50, 0xec, 0x53, 0x7a, 0x12, 0xb5, 0x23,
	0xdb, 0x2d, 0x5f, 0x8a, 0x0f, 0x89, 0xe2, 0x34, 0x60, 0x40, 0x6c, 0x79, 0xef, 0xdd, 0x9f, 0x5f,
	0x7c, 0x30, 0x55, 0xd6, 0x78, 0xaa, 0xcc, 0x1e, 0x4d, 0x58, 0xd4, 0xce, 0x06, 0xcb, 0x59, 0x62,
	0x95, 0x1f, 0x19, 0xb0, 0xd5, 0xb7, 0xe6, 0xe7, 0xd0, 0x23, 0x2d, 0xb2, 0x22, 0x9b, 0xe7, 0xb2,
	0x47, 0x9a, 0x17, 0xc0, 0x34, 0x7a, 0xe5, 0xa8, 0x0e, 0x64, 0x8d, 0xe8, 0xc5, 0x20, 0xb5, 0xf8,
	0x0c, 0x46, 0xef, 0x48, 0xd5, 0x2e, 0x88, 0x7e, 0x91, 0xcd, 0x87, 0xf2, 0xa4, 0xf8, 0x3d, 0x80,
	0xb2, 0x26, 0x6c, 0xc9, 0xa0, 0xf3, 0x62, 0x50, 0xf4, 0xe7, 0x6c, 0x39, 0x5b, 0xa4, 0x38, 0xab,
	0x2e, 0x96, 0x49, 0x25, 0xbf, 0x81, 0xfc, 0x88, 0xde, 0xe3, 0xdb, 0x33, 0x69, 0x31, 0x8c, 0xfb,
	0x26, 0xad, 0xb1, 0xd6, 0xe5, 0x1e, 0xf2, 0xaf, 0xae, 0x3f, 0xac, 0xb7, 0xc0, 0xd4, 0xc1, 0x07,
	0xbb, 0x47, 0xd7, 0xf4, 0xb6, 0xac, 0xd0, 0x59, 0x6b, 0xdd, 0xa0, 0x5a, 0x47, 0x15, 0x99, 0x88,
	0x9a, 0xcb, 0x93, 0xe2, 0x57, 0x30, 0x3e, 0xf8, 0xb6, 0x69, 0xd0, 0x06, 0x8d, 0x5c, 0xeb, 0xf2,
	0x05, 0x26, 0x12, 0x7d, 0x6d, 0x8d, 0x47, 0x2e, 0x60, 0xac, 0x1c, 0x6e, 0x03, 0xb6, 0x2b, 0x27,
	0xb2, 0x93, 0xfc, 0x01, 0xd2, 0x27, 0x8d, 0x7b, 0xd9, 0x52, 0xfc, 0xfe, 0xd5, 0xee, 0x5b, 0xa6,
	0xc5, 0xcb, 0x27, 0xb8, 0xd8, 0xec, 0xa8, 0xae, 0xc9, 0x54, 0x1b, 0x74, 0x47, 0x52, 0xc8, 0x1f,
	0x61, 0xba, 0x8a, 0x93, 0xd3, 0xbb, 0xfc, 0x3b, 0xee, 0xfa, 0xf2, 0x47, 0xd2, 0xe1, 0x96, 0x67,
	0xaf, 0xa3, 0x78, 0xee, 0xbb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x54, 0x6b, 0x12, 0x03,
	0x02, 0x00, 0x00,
}
