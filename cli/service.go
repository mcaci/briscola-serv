package cli

import (
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/briscola-serv/pb"
	"google.golang.org/grpc"
)

func newGRPCClientService(conn *grpc.ClientConn) endpoints {
	pointsEndpoint := newClientPointsEndpoint(conn)
	countEndpoint := newClientCountEndpoint(conn)
	compareEndpoint := newClientCompareEndpoint(conn)
	return endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}

func newClientPointsEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "CardPoints",
		PointsRequestEncodeGRPC, PointsResponseDecodeGRPC,
		pb.CardPointsResponse{},
	).Endpoint()
}

func newClientCountEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "PointCount",
		CountRequestEncodeGRPC, CountResponseDecodeGRPC,
		pb.PointCountResponse{},
	).Endpoint()
}

func newClientCompareEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "CardCompare",
		CompareRequestEncodeGRPC, CompareResponseDecodeGRPC,
		pb.CardCompareResponse{},
	).Endpoint()
}
