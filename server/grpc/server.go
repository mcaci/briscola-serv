package briscola

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endp "github.com/mcaci/briscola-serv/endpoint"
	"github.com/mcaci/briscola-serv/pb"
	"golang.org/x/net/context"
)

type grpcServer struct {
	points  grpctransport.Handler
	count   grpctransport.Handler
	compare grpctransport.Handler
}

func (s *grpcServer) CardPoints(ctx context.Context, r *pb.CardPointsRequest) (*pb.CardPointsResponse, error) {
	_, resp, err := s.points.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardPointsResponse), nil
}

func (s *grpcServer) PointCount(ctx context.Context, r *pb.PointCountRequest) (*pb.PointCountResponse, error) {
	_, resp, err := s.count.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PointCountResponse), nil
}

func (s *grpcServer) CardCompare(ctx context.Context, r *pb.CardCompareRequest) (*pb.CardCompareResponse, error) {
	_, resp, err := s.compare.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardCompareResponse), nil
}

func NewGRPCServer(ctx context.Context, endpoints endp.Endpoints) pb.BriscolaServer {
	return &grpcServer{
		points: grpctransport.NewServer(
			endpoints.CardPointsEndpoint,
			DecodeGRPCPointsRequest,
			EncodeGRPCPointsResponse),
		count: grpctransport.NewServer(
			endpoints.PointCountEndpoint,
			DecodeGRPCCountRequest,
			EncodeGRPCCountResponse),
		compare: grpctransport.NewServer(
			endpoints.CardCompareEndpoint,
			DecodeGRPCCompareRequest,
			EncodeGRPCCompareResponse),
	}
}

func DecodeGRPCCountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PointCountRequest)
	return endp.CountRequest{CardNumbers: req.CardNumber}, nil
}

func EncodeGRPCCountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(endp.CountResponse)
	return &pb.PointCountResponse{Count: res.Points}, nil
}

func DecodeGRPCPointsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardPointsRequest)
	return endp.PointsRequest{CardNumber: req.CardNumber}, nil
}

func EncodeGRPCPointsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(endp.PointsResponse)
	return &pb.CardPointsResponse{Points: res.Points}, nil
}

func DecodeGRPCCompareRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardCompareRequest)
	return endp.CompareRequest{
		FirstCardNumber:  req.FirstCard.GetNumber(),
		FirstCardSeed:    uint32(req.FirstCard.GetSeed()),
		SecondCardNumber: req.SecondCard.GetNumber(),
		SecondCardSeed:   uint32(req.SecondCard.GetSeed()),
		BriscolaSeed:     uint32(req.Briscola)}, nil
}

func EncodeGRPCCompareResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(endp.CompareResponse)
	return &pb.CardCompareResponse{SecondCardWinsOverFirstOne: res.SecondCardWins}, nil
}
