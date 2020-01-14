package grpcclient

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endp "github.com/mcaci/briscola-serv/endpoint"
	"github.com/mcaci/briscola-serv/pb"
	grpcserv "github.com/mcaci/briscola-serv/server/grpc"
	serv "github.com/mcaci/briscola-serv/service"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) serv.Service {
	var pointsEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "CardPoints",
		grpcserv.EncodeGRPCPointsRequest,
		grpcserv.DecodeGRPCPointsResponse,
		pb.CardPointsResponse{},
	).Endpoint()
	var countEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "PointCount",
		grpcserv.EncodeGRPCCountRequest,
		grpcserv.DecodeGRPCCountResponse,
		pb.PointCountResponse{},
	).Endpoint()
	var compareEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "CardCompare",
		grpcserv.EncodeGRPCCompareRequest,
		grpcserv.DecodeGRPCCompareResponse,
		pb.CardCompareResponse{},
	).Endpoint()

	return endp.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
