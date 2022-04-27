package srvgrpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	briscola "github.com/mcaci/briscola-serv/daemon/lib"
	"github.com/mcaci/briscola-serv/pb"
)

type srv struct {
	points  grpctransport.Handler
	count   grpctransport.Handler
	compare grpctransport.Handler
}

func NewServer(ctx context.Context) pb.BriscolaServer {
	return &srv{
		points:  grpctransport.NewServer(briscola.PointsEP, pointsRequestDecode, pointsResponseEncode),
		count:   grpctransport.NewServer(briscola.CountEP, countRequestDecode, countResponseEncode),
		compare: grpctransport.NewServer(briscola.CompareEP, compareRequestDecode, compareResponseEncode),
	}
}

func (s *srv) CardPoints(ctx context.Context, r *pb.CardPointsRequest) (*pb.CardPointsResponse, error) {
	_, resp, err := s.points.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardPointsResponse), nil
}

func (s *srv) PointCount(ctx context.Context, r *pb.PointCountRequest) (*pb.PointCountResponse, error) {
	_, resp, err := s.count.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PointCountResponse), nil
}

func (s *srv) CardCompare(ctx context.Context, r *pb.CardCompareRequest) (*pb.CardCompareResponse, error) {
	_, resp, err := s.compare.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardCompareResponse), nil
}
