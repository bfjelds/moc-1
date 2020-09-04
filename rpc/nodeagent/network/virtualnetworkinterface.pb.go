// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualnetworkinterface.proto

package network

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VirtualNetworkInterface_NetworkInterfaceType int32

const (
	VirtualNetworkInterface_Local  VirtualNetworkInterface_NetworkInterfaceType = 0
	VirtualNetworkInterface_Remote VirtualNetworkInterface_NetworkInterfaceType = 1
)

var VirtualNetworkInterface_NetworkInterfaceType_name = map[int32]string{
	0: "Local",
	1: "Remote",
}

var VirtualNetworkInterface_NetworkInterfaceType_value = map[string]int32{
	"Local":  0,
	"Remote": 1,
}

func (x VirtualNetworkInterface_NetworkInterfaceType) String() string {
	return proto.EnumName(VirtualNetworkInterface_NetworkInterfaceType_name, int32(x))
}

func (VirtualNetworkInterface_NetworkInterfaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3db92eb5944c0691, []int{3, 0}
}

type VirtualNetworkInterfaceRequest struct {
	VirtualNetworkInterfaces []*VirtualNetworkInterface `protobuf:"bytes,1,rep,name=VirtualNetworkInterfaces,proto3" json:"VirtualNetworkInterfaces,omitempty"`
	OperationType            common.Operation           `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *VirtualNetworkInterfaceRequest) Reset()         { *m = VirtualNetworkInterfaceRequest{} }
func (m *VirtualNetworkInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkInterfaceRequest) ProtoMessage()    {}
func (*VirtualNetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db92eb5944c0691, []int{0}
}

func (m *VirtualNetworkInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkInterfaceRequest.Unmarshal(m, b)
}
func (m *VirtualNetworkInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkInterfaceRequest.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkInterfaceRequest.Merge(m, src)
}
func (m *VirtualNetworkInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkInterfaceRequest.Size(m)
}
func (m *VirtualNetworkInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkInterfaceRequest proto.InternalMessageInfo

func (m *VirtualNetworkInterfaceRequest) GetVirtualNetworkInterfaces() []*VirtualNetworkInterface {
	if m != nil {
		return m.VirtualNetworkInterfaces
	}
	return nil
}

func (m *VirtualNetworkInterfaceRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualNetworkInterfaceResponse struct {
	VirtualNetworkInterfaces []*VirtualNetworkInterface `protobuf:"bytes,1,rep,name=VirtualNetworkInterfaces,proto3" json:"VirtualNetworkInterfaces,omitempty"`
	Result                   *wrappers.BoolValue        `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                    string                     `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *VirtualNetworkInterfaceResponse) Reset()         { *m = VirtualNetworkInterfaceResponse{} }
func (m *VirtualNetworkInterfaceResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkInterfaceResponse) ProtoMessage()    {}
func (*VirtualNetworkInterfaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db92eb5944c0691, []int{1}
}

func (m *VirtualNetworkInterfaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkInterfaceResponse.Unmarshal(m, b)
}
func (m *VirtualNetworkInterfaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkInterfaceResponse.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkInterfaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkInterfaceResponse.Merge(m, src)
}
func (m *VirtualNetworkInterfaceResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkInterfaceResponse.Size(m)
}
func (m *VirtualNetworkInterfaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkInterfaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkInterfaceResponse proto.InternalMessageInfo

func (m *VirtualNetworkInterfaceResponse) GetVirtualNetworkInterfaces() []*VirtualNetworkInterface {
	if m != nil {
		return m.VirtualNetworkInterfaces
	}
	return nil
}

func (m *VirtualNetworkInterfaceResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualNetworkInterfaceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type IpConfiguration struct {
	Ipaddress            string                    `protobuf:"bytes,1,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	Prefixlength         string                    `protobuf:"bytes,2,opt,name=prefixlength,proto3" json:"prefixlength,omitempty"`
	Subnetid             string                    `protobuf:"bytes,3,opt,name=subnetid,proto3" json:"subnetid,omitempty"`
	Primary              bool                      `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Gateway              string                    `protobuf:"bytes,5,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Allocation           common.IPAllocationMethod `protobuf:"varint,6,opt,name=allocation,proto3,enum=moc.IPAllocationMethod" json:"allocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *IpConfiguration) Reset()         { *m = IpConfiguration{} }
func (m *IpConfiguration) String() string { return proto.CompactTextString(m) }
func (*IpConfiguration) ProtoMessage()    {}
func (*IpConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db92eb5944c0691, []int{2}
}

func (m *IpConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpConfiguration.Unmarshal(m, b)
}
func (m *IpConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpConfiguration.Marshal(b, m, deterministic)
}
func (m *IpConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpConfiguration.Merge(m, src)
}
func (m *IpConfiguration) XXX_Size() int {
	return xxx_messageInfo_IpConfiguration.Size(m)
}
func (m *IpConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_IpConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_IpConfiguration proto.InternalMessageInfo

func (m *IpConfiguration) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func (m *IpConfiguration) GetPrefixlength() string {
	if m != nil {
		return m.Prefixlength
	}
	return ""
}

func (m *IpConfiguration) GetSubnetid() string {
	if m != nil {
		return m.Subnetid
	}
	return ""
}

func (m *IpConfiguration) GetPrimary() bool {
	if m != nil {
		return m.Primary
	}
	return false
}

func (m *IpConfiguration) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *IpConfiguration) GetAllocation() common.IPAllocationMethod {
	if m != nil {
		return m.Allocation
	}
	return common.IPAllocationMethod_Invalid
}

type VirtualNetworkInterface struct {
	Name                 string                                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                                       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 VirtualNetworkInterface_NetworkInterfaceType `protobuf:"varint,3,opt,name=type,proto3,enum=moc.nodeagent.network.VirtualNetworkInterface_NetworkInterfaceType" json:"type,omitempty"`
	Ipconfigs            []*IpConfiguration                           `protobuf:"bytes,4,rep,name=ipconfigs,proto3" json:"ipconfigs,omitempty"`
	Macaddress           string                                       `protobuf:"bytes,5,opt,name=macaddress,proto3" json:"macaddress,omitempty"`
	DnsSettings          *common.Dns                                  `protobuf:"bytes,6,opt,name=dnsSettings,proto3" json:"dnsSettings,omitempty"`
	VirtualMachineName   string                                       `protobuf:"bytes,7,opt,name=virtualMachineName,proto3" json:"virtualMachineName,omitempty"`
	Status               *common.Status                               `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity                               `protobuf:"bytes,9,opt,name=entity,proto3" json:"entity,omitempty"`
	IovWeight            uint32                                       `protobuf:"varint,10,opt,name=iovWeight,proto3" json:"iovWeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *VirtualNetworkInterface) Reset()         { *m = VirtualNetworkInterface{} }
func (m *VirtualNetworkInterface) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkInterface) ProtoMessage()    {}
func (*VirtualNetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db92eb5944c0691, []int{3}
}

func (m *VirtualNetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkInterface.Unmarshal(m, b)
}
func (m *VirtualNetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkInterface.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkInterface.Merge(m, src)
}
func (m *VirtualNetworkInterface) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkInterface.Size(m)
}
func (m *VirtualNetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkInterface proto.InternalMessageInfo

func (m *VirtualNetworkInterface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualNetworkInterface) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualNetworkInterface) GetType() VirtualNetworkInterface_NetworkInterfaceType {
	if m != nil {
		return m.Type
	}
	return VirtualNetworkInterface_Local
}

func (m *VirtualNetworkInterface) GetIpconfigs() []*IpConfiguration {
	if m != nil {
		return m.Ipconfigs
	}
	return nil
}

func (m *VirtualNetworkInterface) GetMacaddress() string {
	if m != nil {
		return m.Macaddress
	}
	return ""
}

func (m *VirtualNetworkInterface) GetDnsSettings() *common.Dns {
	if m != nil {
		return m.DnsSettings
	}
	return nil
}

func (m *VirtualNetworkInterface) GetVirtualMachineName() string {
	if m != nil {
		return m.VirtualMachineName
	}
	return ""
}

func (m *VirtualNetworkInterface) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualNetworkInterface) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *VirtualNetworkInterface) GetIovWeight() uint32 {
	if m != nil {
		return m.IovWeight
	}
	return 0
}

func init() {
	proto.RegisterEnum("moc.nodeagent.network.VirtualNetworkInterface_NetworkInterfaceType", VirtualNetworkInterface_NetworkInterfaceType_name, VirtualNetworkInterface_NetworkInterfaceType_value)
	proto.RegisterType((*VirtualNetworkInterfaceRequest)(nil), "moc.nodeagent.network.VirtualNetworkInterfaceRequest")
	proto.RegisterType((*VirtualNetworkInterfaceResponse)(nil), "moc.nodeagent.network.VirtualNetworkInterfaceResponse")
	proto.RegisterType((*IpConfiguration)(nil), "moc.nodeagent.network.IpConfiguration")
	proto.RegisterType((*VirtualNetworkInterface)(nil), "moc.nodeagent.network.VirtualNetworkInterface")
}

func init() { proto.RegisterFile("virtualnetworkinterface.proto", fileDescriptor_3db92eb5944c0691) }

var fileDescriptor_3db92eb5944c0691 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x9b, 0x34, 0x4d, 0x26, 0x6d, 0x29, 0x4b, 0xa1, 0x26, 0x94, 0x12, 0x05, 0x09, 0x45,
	0x48, 0x75, 0xa4, 0x50, 0xca, 0xb9, 0x7f, 0x12, 0x91, 0x68, 0x41, 0x5b, 0xd4, 0x4a, 0xdc, 0x36,
	0xf6, 0xc4, 0x59, 0x6a, 0xef, 0x9a, 0xdd, 0x75, 0x4b, 0x5e, 0x8f, 0x97, 0xe0, 0xc2, 0x8d, 0x07,
	0xe0, 0x15, 0x90, 0xd7, 0xce, 0x4f, 0x4b, 0x73, 0xe8, 0x85, 0x9b, 0x67, 0xbe, 0x6f, 0xe7, 0xe7,
	0x9b, 0xf1, 0xc0, 0xf3, 0x2b, 0xae, 0x4c, 0xca, 0x22, 0x81, 0xe6, 0x5a, 0xaa, 0x4b, 0x2e, 0x0c,
	0xaa, 0x01, 0xf3, 0xd1, 0x4b, 0x94, 0x34, 0x92, 0x3c, 0x8e, 0xa5, 0xef, 0x09, 0x19, 0x20, 0x0b,
	0x51, 0x18, 0xaf, 0x60, 0x35, 0xb6, 0x43, 0x29, 0xc3, 0x08, 0x3b, 0x96, 0xd4, 0x4f, 0x07, 0x9d,
	0x6b, 0xc5, 0x92, 0x04, 0x95, 0xce, 0x9f, 0x35, 0x9e, 0xdd, 0xc6, 0x31, 0x4e, 0xcc, 0xa8, 0x00,
	0x57, 0x7c, 0x19, 0xc7, 0x52, 0x14, 0x16, 0x11, 0xd2, 0xf0, 0x01, 0xf7, 0x99, 0xe1, 0x13, 0xdf,
	0xa3, 0x22, 0xcf, 0x2c, 0xb1, 0xf5, 0xc3, 0x81, 0xed, 0xf3, 0xbc, 0xd8, 0xd3, 0x1c, 0xee, 0x8d,
	0x8b, 0xa5, 0xf8, 0x2d, 0x45, 0x6d, 0xc8, 0x57, 0x70, 0xe7, 0x30, 0xb4, 0xeb, 0x34, 0x4b, 0xed,
	0x7a, 0xd7, 0xf3, 0xee, 0x6c, 0xc8, 0x9b, 0x17, 0x78, 0x6e, 0x3c, 0xb2, 0x0b, 0xab, 0x1f, 0x13,
	0x54, 0xb6, 0xec, 0xcf, 0xa3, 0x04, 0xdd, 0xc5, 0xa6, 0xd3, 0x5e, 0xeb, 0xae, 0xd9, 0x04, 0x13,
	0x84, 0xde, 0x24, 0xb5, 0x7e, 0x3a, 0xf0, 0x62, 0x6e, 0x13, 0x3a, 0x91, 0x42, 0xe3, 0x7f, 0xed,
	0xa2, 0x0b, 0x15, 0x8a, 0x3a, 0x8d, 0x8c, 0x2d, 0xbf, 0xde, 0x6d, 0x78, 0xf9, 0xe4, 0xbc, 0xf1,
	0xe4, 0xbc, 0x03, 0x29, 0xa3, 0x73, 0x16, 0xa5, 0x48, 0x0b, 0x26, 0xd9, 0x80, 0xa5, 0x63, 0xa5,
	0xa4, 0x72, 0x4b, 0x4d, 0xa7, 0x5d, 0xa3, 0xb9, 0xd1, 0xfa, 0xe5, 0xc0, 0x83, 0x5e, 0x72, 0x28,
	0xc5, 0x80, 0x87, 0x69, 0xde, 0x31, 0xd9, 0x82, 0x1a, 0x4f, 0x58, 0x10, 0x28, 0xd4, 0x59, 0xe9,
	0x19, 0x7b, 0xea, 0x20, 0x2d, 0x58, 0x49, 0x14, 0x0e, 0xf8, 0xf7, 0x08, 0x45, 0x68, 0x86, 0xb6,
	0x82, 0x1a, 0xbd, 0xe1, 0x23, 0x0d, 0xa8, 0xea, 0xb4, 0x2f, 0xd0, 0xf0, 0xa0, 0x48, 0x37, 0xb1,
	0x89, 0x0b, 0xcb, 0x89, 0xe2, 0x31, 0x53, 0x23, 0xb7, 0xdc, 0x74, 0xda, 0x55, 0x3a, 0x36, 0x33,
	0x24, 0x64, 0x06, 0xaf, 0xd9, 0xc8, 0x5d, 0xb2, 0x8f, 0xc6, 0x26, 0x79, 0x07, 0xc0, 0xa2, 0x48,
	0xe6, 0xdb, 0xe6, 0x56, 0xec, 0xc8, 0x36, 0xad, 0x9a, 0xbd, 0x4f, 0xfb, 0x13, 0xe0, 0x04, 0xcd,
	0x50, 0x06, 0x74, 0x86, 0xda, 0xfa, 0x53, 0x82, 0xcd, 0x39, 0x2a, 0x12, 0x02, 0x65, 0xc1, 0x62,
	0x2c, 0x3a, 0xb4, 0xdf, 0x64, 0x0d, 0x16, 0x79, 0x50, 0xb4, 0xb4, 0xc8, 0x03, 0x72, 0x01, 0x65,
	0x93, 0x6d, 0x49, 0xc9, 0xa6, 0x3c, 0xbc, 0xdf, 0x00, 0xbd, 0xdb, 0x8e, 0x6c, 0x97, 0xa8, 0x0d,
	0x48, 0x8e, 0x32, 0x8d, 0x7d, 0x2b, 0xbb, 0x76, 0xcb, 0x76, 0x3d, 0x5e, 0xcd, 0x89, 0x7e, 0x6b,
	0x3c, 0x74, 0xfa, 0x90, 0x6c, 0x03, 0xc4, 0xcc, 0x1f, 0x8f, 0x2a, 0x17, 0x6d, 0xc6, 0x43, 0x5e,
	0x43, 0x3d, 0x10, 0xfa, 0x0c, 0x8d, 0xe1, 0x22, 0xd4, 0x56, 0xb8, 0x7a, 0xb7, 0x6a, 0xf3, 0x1c,
	0x09, 0x4d, 0x67, 0x41, 0xe2, 0x01, 0x29, 0x8e, 0xca, 0x09, 0xf3, 0x87, 0x5c, 0xe0, 0x69, 0x26,
	0xce, 0xb2, 0x8d, 0x79, 0x07, 0x42, 0x5e, 0x42, 0x45, 0x1b, 0x66, 0x52, 0xed, 0x56, 0x6d, 0xd8,
	0xba, 0x0d, 0x7b, 0x66, 0x5d, 0xb4, 0x80, 0x32, 0x12, 0x0a, 0xc3, 0xcd, 0xc8, 0xad, 0xcd, 0x90,
	0x8e, 0xad, 0x8b, 0x16, 0x90, 0xdd, 0x37, 0x79, 0x75, 0x81, 0x3c, 0x1c, 0x1a, 0x17, 0x9a, 0x4e,
	0x7b, 0x95, 0x4e, 0x1d, 0xad, 0x1d, 0xd8, 0xb8, 0x4b, 0x47, 0x52, 0x83, 0xa5, 0x0f, 0xd2, 0x67,
	0xd1, 0xfa, 0x02, 0x81, 0xec, 0x77, 0x88, 0xa5, 0xc1, 0x75, 0xa7, 0xfb, 0xdb, 0x81, 0xad, 0x39,
	0xf3, 0xd8, 0xcf, 0x64, 0x25, 0x23, 0xa8, 0xf4, 0xc4, 0x95, 0xbc, 0x44, 0xf2, 0xf6, 0x9e, 0xff,
	0x63, 0x7e, 0xae, 0x1a, 0x7b, 0xf7, 0x7d, 0x96, 0x1f, 0x88, 0xd6, 0x02, 0x79, 0x0f, 0x0f, 0x0f,
	0x87, 0xe8, 0x5f, 0x9e, 0xce, 0xdc, 0x4e, 0xf2, 0xe4, 0x9f, 0x7f, 0xf7, 0x38, 0xbb, 0xba, 0x8d,
	0xa7, 0x36, 0xcd, 0x2c, 0x75, 0x1a, 0xe9, 0x60, 0xef, 0xcb, 0x6e, 0xc8, 0xcd, 0x30, 0xed, 0x7b,
	0xbe, 0x8c, 0x3b, 0x31, 0xf7, 0x95, 0xd4, 0x72, 0x60, 0x3a, 0xb1, 0xf4, 0x77, 0x6c, 0x9c, 0x8e,
	0x4a, 0xfc, 0xce, 0xa4, 0xc6, 0x4e, 0x51, 0x63, 0xbf, 0x62, 0xc1, 0x37, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x60, 0x7e, 0xfe, 0x43, 0x40, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualNetworkInterfaceAgentClient is the client API for VirtualNetworkInterfaceAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualNetworkInterfaceAgentClient interface {
	Invoke(ctx context.Context, in *VirtualNetworkInterfaceRequest, opts ...grpc.CallOption) (*VirtualNetworkInterfaceResponse, error)
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
}

type virtualNetworkInterfaceAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualNetworkInterfaceAgentClient(cc *grpc.ClientConn) VirtualNetworkInterfaceAgentClient {
	return &virtualNetworkInterfaceAgentClient{cc}
}

func (c *virtualNetworkInterfaceAgentClient) Invoke(ctx context.Context, in *VirtualNetworkInterfaceRequest, opts ...grpc.CallOption) (*VirtualNetworkInterfaceResponse, error) {
	out := new(VirtualNetworkInterfaceResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.network.VirtualNetworkInterfaceAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualNetworkInterfaceAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.network.VirtualNetworkInterfaceAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualNetworkInterfaceAgentServer is the server API for VirtualNetworkInterfaceAgent service.
type VirtualNetworkInterfaceAgentServer interface {
	Invoke(context.Context, *VirtualNetworkInterfaceRequest) (*VirtualNetworkInterfaceResponse, error)
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
}

// UnimplementedVirtualNetworkInterfaceAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualNetworkInterfaceAgentServer struct {
}

func (*UnimplementedVirtualNetworkInterfaceAgentServer) Invoke(ctx context.Context, req *VirtualNetworkInterfaceRequest) (*VirtualNetworkInterfaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedVirtualNetworkInterfaceAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}

func RegisterVirtualNetworkInterfaceAgentServer(s *grpc.Server, srv VirtualNetworkInterfaceAgentServer) {
	s.RegisterService(&_VirtualNetworkInterfaceAgent_serviceDesc, srv)
}

func _VirtualNetworkInterfaceAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualNetworkInterfaceAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.network.VirtualNetworkInterfaceAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualNetworkInterfaceAgentServer).Invoke(ctx, req.(*VirtualNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualNetworkInterfaceAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualNetworkInterfaceAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.network.VirtualNetworkInterfaceAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualNetworkInterfaceAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualNetworkInterfaceAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.network.VirtualNetworkInterfaceAgent",
	HandlerType: (*VirtualNetworkInterfaceAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualNetworkInterfaceAgent_Invoke_Handler,
		},
		{
			MethodName: "CheckNotification",
			Handler:    _VirtualNetworkInterfaceAgent_CheckNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualnetworkinterface.proto",
}