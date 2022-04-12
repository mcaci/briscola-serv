package briscola

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	endp "github.com/mcaci/briscola-serv/daemon/endpoint"
	"github.com/mcaci/briscola-serv/pb"
)

type grpcServer struct {
	points  grpctransport.Handler
	count   grpctransport.Handler
	compare grpctransport.Handler
}

func NewGRPCServer(ctx context.Context, endpoints endp.Endpoints) pb.BriscolaServer {
	return &grpcServer{
		points:  grpctransport.NewServer(endpoints.CardPointsEndpoint, endp.PointsRequestDecodeGRPC, endp.PointsResponseEncodeGRPC),
		count:   grpctransport.NewServer(endpoints.PointCountEndpoint, endp.CountRequestDecodeGRPC, endp.CountResponseEncodeGRPC),
		compare: grpctransport.NewServer(endpoints.CardCompareEndpoint, endp.CompareRequestDecodeGRPC, endp.CompareResponseEncodeGRPC),
	}
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
