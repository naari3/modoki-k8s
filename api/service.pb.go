// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package modoki is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	ServiceSpec
	ServiceState
	ServiceCreateRequest
	ServiceCreateResponse
*/
package modoki

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServiceStatus int32

const (
	ServiceStatus_Deploying  ServiceStatus = 0
	ServiceStatus_Running    ServiceStatus = 1
	ServiceStatus_Terminated ServiceStatus = 2
	ServiceStatus_Error      ServiceStatus = 3
)

var ServiceStatus_name = map[int32]string{
	0: "Deploying",
	1: "Running",
	2: "Terminated",
	3: "Error",
}
var ServiceStatus_value = map[string]int32{
	"Deploying":  0,
	"Running":    1,
	"Terminated": 2,
	"Error":      3,
}

func (x ServiceStatus) String() string {
	return proto.EnumName(ServiceStatus_name, int32(x))
}
func (ServiceStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServiceSpec struct {
	Name    string                          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Image   string                          `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	Command []string                        `protobuf:"bytes,3,rep,name=command" json:"command,omitempty"`
	Args    []string                        `protobuf:"bytes,4,rep,name=args" json:"args,omitempty"`
	Owner   int32                           `protobuf:"varint,5,opt,name=owner" json:"owner,omitempty"`
	Options map[string]*google_protobuf.Any `protobuf:"bytes,10,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ServiceSpec) Reset()                    { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string            { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()               {}
func (*ServiceSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceSpec) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ServiceSpec) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *ServiceSpec) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ServiceSpec) GetOwner() int32 {
	if m != nil {
		return m.Owner
	}
	return 0
}

func (m *ServiceSpec) GetOptions() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Options
	}
	return nil
}

type ServiceState struct {
	State        *ServiceState                   `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	StateMessage string                          `protobuf:"bytes,2,opt,name=stateMessage" json:"stateMessage,omitempty"`
	Options      map[string]*google_protobuf.Any `protobuf:"bytes,10,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CreatedAt    *google_protobuf1.Timestamp     `protobuf:"bytes,11,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt    *google_protobuf1.Timestamp     `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *ServiceState) Reset()                    { *m = ServiceState{} }
func (m *ServiceState) String() string            { return proto.CompactTextString(m) }
func (*ServiceState) ProtoMessage()               {}
func (*ServiceState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceState) GetState() *ServiceState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ServiceState) GetStateMessage() string {
	if m != nil {
		return m.StateMessage
	}
	return ""
}

func (m *ServiceState) GetOptions() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ServiceState) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ServiceState) GetUpdatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ServiceCreateRequest struct {
	Spec *ServiceSpec `protobuf:"bytes,1,opt,name=spec" json:"spec,omitempty"`
}

func (m *ServiceCreateRequest) Reset()                    { *m = ServiceCreateRequest{} }
func (m *ServiceCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceCreateRequest) ProtoMessage()               {}
func (*ServiceCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceCreateRequest) GetSpec() *ServiceSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ServiceCreateResponse struct {
	Id   string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec *ServiceSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
}

func (m *ServiceCreateResponse) Reset()                    { *m = ServiceCreateResponse{} }
func (m *ServiceCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceCreateResponse) ProtoMessage()               {}
func (*ServiceCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServiceCreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceCreateResponse) GetSpec() *ServiceSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "modoki.ServiceSpec")
	proto.RegisterType((*ServiceState)(nil), "modoki.ServiceState")
	proto.RegisterType((*ServiceCreateRequest)(nil), "modoki.ServiceCreateRequest")
	proto.RegisterType((*ServiceCreateResponse)(nil), "modoki.ServiceCreateResponse")
	proto.RegisterEnum("modoki.ServiceStatus", ServiceStatus_name, ServiceStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	Create(ctx context.Context, in *ServiceCreateRequest, opts ...grpc.CallOption) (*ServiceCreateResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Create(ctx context.Context, in *ServiceCreateRequest, opts ...grpc.CallOption) (*ServiceCreateResponse, error) {
	out := new(ServiceCreateResponse)
	err := grpc.Invoke(ctx, "/modoki.Service/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	Create(context.Context, *ServiceCreateRequest) (*ServiceCreateResponse, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Service/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Create(ctx, req.(*ServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modoki.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Service_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x18, 0x24, 0x49, 0xd3, 0x92, 0x2f, 0xed, 0xaa, 0x32, 0x45, 0x0a, 0x11, 0x88, 0xd0, 0x0b, 0xd5,
	0x1e, 0xb2, 0x52, 0xb9, 0xc0, 0x72, 0x40, 0x05, 0xca, 0x0d, 0x51, 0x79, 0xf7, 0xc4, 0x05, 0x79,
	0x13, 0x13, 0x45, 0xdb, 0xd8, 0xc1, 0x76, 0x16, 0xe5, 0x11, 0x78, 0x4e, 0x5e, 0x04, 0xd9, 0x4e,
	0xa0, 0x1b, 0x8a, 0xf6, 0xc4, 0xed, 0xfb, 0x99, 0x6f, 0x32, 0x33, 0x6e, 0x61, 0x26, 0xa9, 0xb8,
	0x29, 0x33, 0x9a, 0xd6, 0x82, 0x2b, 0x8e, 0xc6, 0x15, 0xcf, 0xf9, 0x75, 0x19, 0x3f, 0x2a, 0x38,
	0x2f, 0xf6, 0xf4, 0xcc, 0x4c, 0xaf, 0x9a, 0xaf, 0x67, 0x84, 0xb5, 0x16, 0x12, 0x3f, 0x1d, 0xae,
	0x54, 0x59, 0x51, 0xa9, 0x48, 0x55, 0x5b, 0xc0, 0xf2, 0x87, 0x0b, 0xe1, 0x85, 0x65, 0xbd, 0xa8,
	0x69, 0x86, 0x10, 0x8c, 0x18, 0xa9, 0x68, 0xe4, 0x24, 0xce, 0x2a, 0xc0, 0xa6, 0x46, 0x0b, 0xf0,
	0xcb, 0x8a, 0x14, 0x34, 0x72, 0xcd, 0xd0, 0x36, 0x28, 0x82, 0x49, 0xc6, 0xab, 0x8a, 0xb0, 0x3c,
	0xf2, 0x12, 0x6f, 0x15, 0xe0, 0xbe, 0xd5, 0x1c, 0x44, 0x14, 0x32, 0x1a, 0x99, 0xb1, 0xa9, 0x35,
	0x07, 0xff, 0xce, 0xa8, 0x88, 0xfc, 0xc4, 0x59, 0xf9, 0xd8, 0x36, 0xe8, 0x1c, 0x26, 0xbc, 0x56,
	0x25, 0x67, 0x32, 0x82, 0xc4, 0x5b, 0x85, 0xeb, 0x24, 0xb5, 0x9e, 0xd2, 0x03, 0x4d, 0xe9, 0x27,
	0x0b, 0xd9, 0x32, 0x25, 0x5a, 0xdc, 0x1f, 0xc4, 0x3b, 0x98, 0x1e, 0x2e, 0xd0, 0x1c, 0xbc, 0x6b,
	0xda, 0x76, 0xc2, 0x75, 0x89, 0x4e, 0xc1, 0xbf, 0x21, 0xfb, 0xc6, 0xea, 0x0e, 0xd7, 0x8b, 0xd4,
	0x86, 0x91, 0xf6, 0x61, 0xa4, 0x1b, 0xd6, 0x62, 0x0b, 0x39, 0x77, 0x5f, 0x3a, 0xcb, 0x9f, 0x2e,
	0x4c, 0xfb, 0xef, 0x2a, 0xa2, 0xa8, 0x26, 0x90, 0xba, 0x30, 0xa4, 0x9a, 0x60, 0x20, 0x4e, 0xef,
	0xb0, 0x85, 0xa0, 0x25, 0x4c, 0x4d, 0xf1, 0x91, 0x4a, 0xf9, 0x27, 0xab, 0x5b, 0x33, 0xf4, 0x7a,
	0x68, 0xf7, 0xd9, 0x31, 0xc6, 0xe3, 0x7e, 0xd1, 0x2b, 0x80, 0x4c, 0x50, 0xa2, 0x68, 0xfe, 0x85,
	0xa8, 0x28, 0x34, 0x8a, 0xe2, 0xbf, 0x2c, 0x5d, 0xf6, 0xef, 0x8b, 0x83, 0x0e, 0xbd, 0x51, 0xfa,
	0xb4, 0xa9, 0xf3, 0xfe, 0x74, 0x7a, 0xf7, 0x69, 0x87, 0xde, 0xa8, 0xff, 0x90, 0xf2, 0x1b, 0x58,
	0x74, 0x6e, 0xdf, 0x19, 0x81, 0x98, 0x7e, 0x6b, 0xa8, 0x54, 0xe8, 0x39, 0x8c, 0x64, 0x4d, 0xb3,
	0x2e, 0xeb, 0x07, 0x47, 0x7e, 0x08, 0xd8, 0x00, 0x96, 0x3b, 0x78, 0x38, 0x20, 0x90, 0x35, 0x67,
	0x92, 0xa2, 0x13, 0x70, 0xcb, 0xbc, 0x93, 0xe6, 0x96, 0xf9, 0x6f, 0x46, 0xf7, 0x0e, 0xc6, 0xd3,
	0x0f, 0x30, 0x3b, 0x78, 0x80, 0x46, 0xa2, 0x19, 0x04, 0xef, 0x69, 0xbd, 0xe7, 0x6d, 0xc9, 0x8a,
	0xf9, 0x3d, 0x14, 0xc2, 0x04, 0x37, 0x8c, 0xe9, 0xc6, 0x41, 0x27, 0x00, 0x97, 0x54, 0x54, 0x25,
	0xd3, 0x09, 0xcd, 0x5d, 0x14, 0x80, 0xbf, 0x15, 0x82, 0x8b, 0xb9, 0xb7, 0xde, 0xc1, 0xa4, 0xe3,
	0x41, 0x5b, 0x18, 0x5b, 0x75, 0xe8, 0xf1, 0xe0, 0xbb, 0xb7, 0x5c, 0xc7, 0x4f, 0xfe, 0xb1, 0xb5,
	0x96, 0xde, 0xde, 0xff, 0xdc, 0xfd, 0xc9, 0xaf, 0xc6, 0x26, 0xcf, 0x17, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x3e, 0xfe, 0xa4, 0x04, 0x04, 0x00, 0x00,
}
