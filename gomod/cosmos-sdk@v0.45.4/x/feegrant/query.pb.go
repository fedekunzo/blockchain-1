// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/feegrant/v1beta1/query.proto

package feegrant

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryAllowanceRequest is the request type for the Query/Allowance RPC method.
type QueryAllowanceRequest struct {
	// granter is the address of the user granting an allowance of their funds.
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
}

func (m *QueryAllowanceRequest) Reset()         { *m = QueryAllowanceRequest{} }
func (m *QueryAllowanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllowanceRequest) ProtoMessage()    {}
func (*QueryAllowanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59efc303945de53f, []int{0}
}
func (m *QueryAllowanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllowanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllowanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllowanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllowanceRequest.Merge(m, src)
}
func (m *QueryAllowanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllowanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllowanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllowanceRequest proto.InternalMessageInfo

func (m *QueryAllowanceRequest) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *QueryAllowanceRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

// QueryAllowanceResponse is the response type for the Query/Allowance RPC method.
type QueryAllowanceResponse struct {
	// allowance is a allowance granted for grantee by granter.
	Allowance *Grant `protobuf:"bytes,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (m *QueryAllowanceResponse) Reset()         { *m = QueryAllowanceResponse{} }
func (m *QueryAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllowanceResponse) ProtoMessage()    {}
func (*QueryAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59efc303945de53f, []int{1}
}
func (m *QueryAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllowanceResponse.Merge(m, src)
}
func (m *QueryAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllowanceResponse proto.InternalMessageInfo

func (m *QueryAllowanceResponse) GetAllowance() *Grant {
	if m != nil {
		return m.Allowance
	}
	return nil
}

// QueryAllowancesRequest is the request type for the Query/Allowances RPC method.
type QueryAllowancesRequest struct {
	Grantee string `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllowancesRequest) Reset()         { *m = QueryAllowancesRequest{} }
func (m *QueryAllowancesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllowancesRequest) ProtoMessage()    {}
func (*QueryAllowancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59efc303945de53f, []int{2}
}
func (m *QueryAllowancesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllowancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllowancesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllowancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllowancesRequest.Merge(m, src)
}
func (m *QueryAllowancesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllowancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllowancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllowancesRequest proto.InternalMessageInfo

func (m *QueryAllowancesRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *QueryAllowancesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAllowancesResponse is the response type for the Query/Allowances RPC method.
type QueryAllowancesResponse struct {
	// allowances are allowance's granted for grantee by granter.
	Allowances []*Grant `protobuf:"bytes,1,rep,name=allowances,proto3" json:"allowances,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllowancesResponse) Reset()         { *m = QueryAllowancesResponse{} }
func (m *QueryAllowancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllowancesResponse) ProtoMessage()    {}
func (*QueryAllowancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59efc303945de53f, []int{3}
}
func (m *QueryAllowancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllowancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllowancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllowancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllowancesResponse.Merge(m, src)
}
func (m *QueryAllowancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllowancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllowancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllowancesResponse proto.InternalMessageInfo

func (m *QueryAllowancesResponse) GetAllowances() []*Grant {
	if m != nil {
		return m.Allowances
	}
	return nil
}

func (m *QueryAllowancesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllowanceRequest)(nil), "cosmos.feegrant.v1beta1.QueryAllowanceRequest")
	proto.RegisterType((*QueryAllowanceResponse)(nil), "cosmos.feegrant.v1beta1.QueryAllowanceResponse")
	proto.RegisterType((*QueryAllowancesRequest)(nil), "cosmos.feegrant.v1beta1.QueryAllowancesRequest")
	proto.RegisterType((*QueryAllowancesResponse)(nil), "cosmos.feegrant.v1beta1.QueryAllowancesResponse")
}

func init() {
	proto.RegisterFile("cosmos/feegrant/v1beta1/query.proto", fileDescriptor_59efc303945de53f)
}

var fileDescriptor_59efc303945de53f = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4b, 0x4d, 0x4d, 0x2f, 0x4a, 0xcc, 0x2b, 0xd1, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x87, 0x28, 0xd2, 0x83, 0x29, 0xd2, 0x83, 0x2a, 0x92, 0x52, 0xc3, 0xa5, 0x1b, 0xae, 0x12,
	0x6c, 0x80, 0x94, 0x16, 0x54, 0x5d, 0x52, 0x62, 0x71, 0x2a, 0xc4, 0x64, 0xb8, 0xca, 0x82, 0xc4,
	0xf4, 0xcc, 0xbc, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0xa8, 0x5a, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c,
	0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x12, 0xb0, 0x64, 0x31, 0x44, 0x56,
	0xc9, 0x9b, 0x4b, 0x34, 0x10, 0xa4, 0xdf, 0x31, 0x27, 0x27, 0xbf, 0x3c, 0x31, 0x2f, 0x39, 0x35,
	0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x82, 0x8b, 0x1d, 0x6c, 0x63, 0x6a, 0x91, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x8b, 0x90, 0x49, 0x95, 0x60, 0x42, 0x96, 0x49, 0x55,
	0x0a, 0xe3, 0x12, 0x43, 0x37, 0xac, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x86, 0x8b, 0x33,
	0x11, 0x26, 0x08, 0x36, 0x8f, 0xdb, 0x48, 0x4e, 0x0f, 0x47, 0x28, 0xe8, 0xb9, 0x83, 0x78, 0x41,
	0x08, 0x0d, 0x4a, 0x55, 0xe8, 0xe6, 0x16, 0x63, 0xb8, 0x32, 0x15, 0xd5, 0x95, 0xa9, 0x42, 0x6e,
	0x5c, 0x5c, 0x88, 0xa0, 0x00, 0x3b, 0x94, 0xdb, 0x48, 0x0d, 0x66, 0x25, 0x28, 0xdc, 0xf4, 0x20,
	0x31, 0x02, 0xb3, 0x34, 0x20, 0x31, 0x1d, 0xe6, 0xf7, 0x20, 0x24, 0x9d, 0x4a, 0x8b, 0x18, 0xb9,
	0xc4, 0x31, 0x2c, 0x87, 0xfa, 0xca, 0x8e, 0x8b, 0x0b, 0xee, 0xc8, 0x62, 0x09, 0x46, 0x05, 0x66,
	0x22, 0xbc, 0x85, 0xa4, 0x43, 0xc8, 0x1d, 0x8b, 0x1b, 0xd5, 0x09, 0xba, 0x11, 0x62, 0x39, 0xb2,
	0x23, 0x8d, 0xee, 0x33, 0x71, 0xb1, 0x82, 0x1d, 0x29, 0xb4, 0x86, 0x91, 0x8b, 0x13, 0xee, 0x52,
	0x21, 0x3d, 0x9c, 0x8e, 0xc1, 0x1a, 0xe9, 0x52, 0xfa, 0x44, 0xab, 0x87, 0x38, 0x42, 0xc9, 0xae,
	0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x16, 0x42, 0x66, 0xfa, 0xb8, 0x52, 0x2e, 0xdc, 0xbb, 0xfa, 0xd5,
	0xd0, 0x04, 0x54, 0x0b, 0x63, 0xa5, 0xd6, 0x0a, 0x2d, 0x63, 0xe4, 0xe2, 0x42, 0x04, 0xac, 0x10,
	0xb1, 0xf6, 0xc3, 0xe2, 0x5f, 0xca, 0x80, 0x78, 0x0d, 0x50, 0x17, 0x9b, 0x82, 0x5d, 0xac, 0x2f,
	0xa4, 0x4b, 0xd8, 0xc5, 0xc5, 0x08, 0x87, 0x3a, 0x39, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0xcc,
	0x48, 0x08, 0xa5, 0x5b, 0x9c, 0x92, 0xad, 0x5f, 0x01, 0x37, 0x3f, 0x89, 0x0d, 0x9c, 0xe3, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x08, 0xf0, 0xae, 0x23, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.

// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Allowance returns fee granted to the grantee by the granter.
	Allowance(ctx context.Context, in *QueryAllowanceRequest, opts ...grpc.CallOption) (*QueryAllowanceResponse, error)
	// Allowances returns all the grants for address.
	Allowances(ctx context.Context, in *QueryAllowancesRequest, opts ...grpc.CallOption) (*QueryAllowancesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Allowance(ctx context.Context, in *QueryAllowanceRequest, opts ...grpc.CallOption) (*QueryAllowanceResponse, error) {
	out := new(QueryAllowanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Query/Allowance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Allowances(ctx context.Context, in *QueryAllowancesRequest, opts ...grpc.CallOption) (*QueryAllowancesResponse, error) {
	out := new(QueryAllowancesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Query/Allowances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Allowance returns fee granted to the grantee by the granter.
	Allowance(context.Context, *QueryAllowanceRequest) (*QueryAllowanceResponse, error)
	// Allowances returns all the grants for address.
	Allowances(context.Context, *QueryAllowancesRequest) (*QueryAllowancesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Allowance(ctx context.Context, req *QueryAllowanceRequest) (*QueryAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allowance not implemented")
}
func (*UnimplementedQueryServer) Allowances(ctx context.Context, req *QueryAllowancesRequest) (*QueryAllowancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allowances not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Allowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllowanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Allowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Query/Allowance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Allowance(ctx, req.(*QueryAllowanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Allowances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllowancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Allowances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Query/Allowances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Allowances(ctx, req.(*QueryAllowancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.feegrant.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allowance",
			Handler:    _Query_Allowance_Handler,
		},
		{
			MethodName: "Allowances",
			Handler:    _Query_Allowances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/feegrant/v1beta1/query.proto",
}

func (m *QueryAllowanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllowanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllowanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllowancesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllowancesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllowancesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllowancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllowancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllowancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Allowances) > 0 {
		for iNdEx := len(m.Allowances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allowances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAllowanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllowancesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllowancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Allowances) > 0 {
		for _, e := range m.Allowances {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllowanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllowanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllowanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllowanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &Grant{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllowancesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllowancesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllowancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllowancesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllowancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllowancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allowances = append(m.Allowances, &Grant{})
			if err := m.Allowances[len(m.Allowances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
