package briscola

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/briscola-serv/pb"
)

type grpcServer struct {
	points  grpctransport.Handler
	count   grpctransport.Handler
	compare grpctransport.Handler
}

type endpointInt interface {
	Cp() endpoint.Endpoint
	Pc() endpoint.Endpoint
	Cc() endpoint.Endpoint
}

func NewGRPCServer(ctx context.Context, ep endpointInt) pb.BriscolaServer {
	return &grpcServer{
		points:  grpctransport.NewServer(ep.Cp(), PointsRequestDecodeGRPC, PointsResponseEncodeGRPC),
		count:   grpctransport.NewServer(ep.Pc(), CountRequestDecodeGRPC, CountResponseEncodeGRPC),
		compare: grpctransport.NewServer(ep.Cc(), CompareRequestDecodeGRPC, CompareResponseEncodeGRPC),
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
