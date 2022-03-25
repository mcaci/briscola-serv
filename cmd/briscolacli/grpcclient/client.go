package grpcclient

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	endp "github.com/mcaci/briscola-serv/endpoint"
	"github.com/mcaci/briscola-serv/pb"
	serv "github.com/mcaci/briscola-serv/service"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) serv.Service {
	var (
		pointsEndpoint = grpctransport.NewClient(
			conn, "pb.Briscola", "CardPoints",
			EncodeGRPCPointsRequest,
			DecodeGRPCPointsResponse,
			pb.CardPointsResponse{},
		).Endpoint()
		countEndpoint = grpctransport.NewClient(
			conn, "pb.Briscola", "PointCount",
			EncodeGRPCCountRequest,
			DecodeGRPCCountResponse,
			pb.PointCountResponse{},
		).Endpoint()
		compareEndpoint = grpctransport.NewClient(
			conn, "pb.Briscola", "CardCompare",
			EncodeGRPCCompareRequest,
			DecodeGRPCCompareResponse,
			pb.CardCompareResponse{},
		).Endpoint()
	)

	return endp.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}

func EncodeGRPCCountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(endp.CountRequest)
	return &pb.PointCountRequest{CardNumber: req.CardNumbers}, nil
}

func DecodeGRPCCountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.PointCountResponse)
	return endp.CountResponse{Points: res.Count, Err: ""}, nil
}

func EncodeGRPCPointsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(endp.PointsRequest)
	return &pb.CardPointsRequest{CardNumber: req.CardNumber}, nil
}

func DecodeGRPCPointsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardPointsResponse)
	return endp.PointsResponse{Points: res.Points, Err: ""}, nil
}

func EncodeGRPCCompareRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(endp.CompareRequest)
	first := &pb.ItalianCard{Number: req.FirstCardNumber, Seed: pb.Seed(req.FirstCardSeed)}
	second := &pb.ItalianCard{Number: req.SecondCardNumber, Seed: pb.Seed(req.SecondCardSeed)}
	return &pb.CardCompareRequest{FirstCard: first, SecondCard: second, Briscola: pb.Seed(req.BriscolaSeed)}, nil
}

func DecodeGRPCCompareResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardCompareResponse)
	return endp.CompareResponse{SecondCardWins: res.SecondCardWinsOverFirstOne, Err: ""}, nil
}
