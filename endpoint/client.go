package endp

import (
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/briscola-serv/pb"
	"google.golang.org/grpc"
)

func NewClientPointsEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "CardPoints",
		PointsRequestEncodeGRPC, PointsResponseDecodeGRPC,
		pb.CardPointsResponse{},
	).Endpoint()
}

func NewClientCountEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "PointCount",
		CountRequestEncodeGRPC, CountResponseDecodeGRPC,
		pb.PointCountResponse{},
	).Endpoint()
}

func NewClientCompareEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "CardCompare",
		CompareRequestEncodeGRPC, CompareResponseDecodeGRPC,
		pb.CardCompareResponse{},
	).Endpoint()
}
