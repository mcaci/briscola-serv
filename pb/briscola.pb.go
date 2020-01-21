// Code generated by protoc-gen-go. DO NOT EDIT.
// source: briscola/pb/briscola.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

func init() { proto.RegisterFile("briscola/pb/briscola.proto", fileDescriptor_9f8c8ee6503118d5) }

var fileDescriptor_9f8c8ee6503118d5 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2a, 0xca, 0x2c,
	0x4e, 0xce, 0xcf, 0x49, 0xd4, 0x2f, 0x48, 0xd2, 0x87, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x98, 0x0a, 0x92, 0xa4, 0x24, 0x91, 0xe5, 0x93, 0xf3, 0x73, 0x0b, 0x12, 0x8b, 0x52, 0x21,
	0xd2, 0x52, 0xe2, 0xa8, 0x52, 0xa5, 0x79, 0x25, 0x50, 0x09, 0x09, 0x64, 0x89, 0x82, 0xfc, 0xcc,
	0xbc, 0x92, 0x62, 0x88, 0x8c, 0xd1, 0x29, 0x46, 0x2e, 0x0e, 0x27, 0xa8, 0xa4, 0x90, 0x2d, 0x17,
	0x97, 0x73, 0x62, 0x51, 0x4a, 0x00, 0x58, 0x81, 0x90, 0xa8, 0x5e, 0x41, 0x92, 0x1e, 0x82, 0x1f,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x25, 0x86, 0x2e, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0xaa, 0xc4, 0x00, 0xd2, 0x0e, 0x16, 0x73, 0x06, 0xd9, 0x0c, 0xd1, 0x8e, 0xe0, 0xa3, 0x68, 0x47,
	0x16, 0x86, 0x6b, 0x77, 0xe0, 0xe2, 0x06, 0x19, 0xeb, 0x0c, 0xf1, 0x92, 0x10, 0xdc, 0x1e, 0xa8,
	0x00, 0xcc, 0x00, 0x71, 0x0c, 0x71, 0x98, 0x09, 0x49, 0x6c, 0x60, 0x3f, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xb0, 0x12, 0x33, 0x1b, 0x43, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BriscolaClient is the client API for Briscola service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BriscolaClient interface {
	CardPoints(ctx context.Context, in *CardPointsRequest, opts ...grpc.CallOption) (*CardPointsResponse, error)
	PointCount(ctx context.Context, in *PointCountRequest, opts ...grpc.CallOption) (*PointCountResponse, error)
	CardCompare(ctx context.Context, in *CardCompareRequest, opts ...grpc.CallOption) (*CardCompareResponse, error)
}

type briscolaClient struct {
	cc *grpc.ClientConn
}

func NewBriscolaClient(cc *grpc.ClientConn) BriscolaClient {
	return &briscolaClient{cc}
}

func (c *briscolaClient) CardPoints(ctx context.Context, in *CardPointsRequest, opts ...grpc.CallOption) (*CardPointsResponse, error) {
	out := new(CardPointsResponse)
	err := c.cc.Invoke(ctx, "/pb.Briscola/CardPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briscolaClient) PointCount(ctx context.Context, in *PointCountRequest, opts ...grpc.CallOption) (*PointCountResponse, error) {
	out := new(PointCountResponse)
	err := c.cc.Invoke(ctx, "/pb.Briscola/PointCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briscolaClient) CardCompare(ctx context.Context, in *CardCompareRequest, opts ...grpc.CallOption) (*CardCompareResponse, error) {
	out := new(CardCompareResponse)
	err := c.cc.Invoke(ctx, "/pb.Briscola/CardCompare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BriscolaServer is the server API for Briscola service.
type BriscolaServer interface {
	CardPoints(context.Context, *CardPointsRequest) (*CardPointsResponse, error)
	PointCount(context.Context, *PointCountRequest) (*PointCountResponse, error)
	CardCompare(context.Context, *CardCompareRequest) (*CardCompareResponse, error)
}

// UnimplementedBriscolaServer can be embedded to have forward compatible implementations.
type UnimplementedBriscolaServer struct {
}

func (*UnimplementedBriscolaServer) CardPoints(ctx context.Context, req *CardPointsRequest) (*CardPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardPoints not implemented")
}
func (*UnimplementedBriscolaServer) PointCount(ctx context.Context, req *PointCountRequest) (*PointCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PointCount not implemented")
}
func (*UnimplementedBriscolaServer) CardCompare(ctx context.Context, req *CardCompareRequest) (*CardCompareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardCompare not implemented")
}

func RegisterBriscolaServer(s *grpc.Server, srv BriscolaServer) {
	s.RegisterService(&_Briscola_serviceDesc, srv)
}

func _Briscola_CardPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriscolaServer).CardPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Briscola/CardPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriscolaServer).CardPoints(ctx, req.(*CardPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Briscola_PointCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriscolaServer).PointCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Briscola/PointCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriscolaServer).PointCount(ctx, req.(*PointCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Briscola_CardCompare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardCompareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriscolaServer).CardCompare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Briscola/CardCompare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriscolaServer).CardCompare(ctx, req.(*CardCompareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Briscola_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Briscola",
	HandlerType: (*BriscolaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CardPoints",
			Handler:    _Briscola_CardPoints_Handler,
		},
		{
			MethodName: "PointCount",
			Handler:    _Briscola_PointCount_Handler,
		},
		{
			MethodName: "CardCompare",
			Handler:    _Briscola_CardCompare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "briscola/pb/briscola.proto",
}