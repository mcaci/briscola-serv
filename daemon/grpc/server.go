package srvgrpc

import (
	"context"
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	briscola "github.com/mcaci/briscola-serv/daemon/lib"
	"github.com/mcaci/briscola-serv/pb"
)

type srv struct {
	points  grpctransport.Handler
	count   grpctransport.Handler
	compare grpctransport.Handler
}

func NewServer(ctx context.Context) *srv {
	return &srv{
		points:  grpctransport.NewServer(briscola.PointsEP, pointsRequestDecode, pointsResponseEncode),
		count:   grpctransport.NewServer(briscola.CountEP, countRequestDecode, countResponseEncode),
		compare: grpctransport.NewServer(briscola.CompareEP, compareRequestDecode, compareResponseEncode),
	}
}

func (s *srv) SetPointsHandler(h grpctransport.Handler)  { s.points = h }
func (s *srv) SetCountHandler(h grpctransport.Handler)   { s.count = h }
func (s *srv) SetCompareHandler(h grpctransport.Handler) { s.compare = h }

func (s *srv) CardPoints(pctx context.Context, r *pb.CardPointsRequest) (*pb.CardPointsResponse, error) {
	ctx, resp, err := s.points.ServeGRPC(pctx, r)
	if err != nil {
		return nil, err
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	rs, ok := resp.(*pb.CardPointsResponse)
	if !ok {
		return nil, fmt.Errorf("expecting response of type *pb.CardPointsResponse but got %#v", rs)
	}
	return rs, nil
}

func (s *srv) PointCount(ctx context.Context, r *pb.PointCountRequest) (*pb.PointCountResponse, error) {
	ctx, resp, err := s.count.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	rs, ok := resp.(*pb.PointCountResponse)
	if !ok {
		return nil, fmt.Errorf("expecting response of type *pb.PointCountResponse but got %#v", rs)
	}
	return rs, nil
}

func (s *srv) CardCompare(ctx context.Context, r *pb.CardCompareRequest) (*pb.CardCompareResponse, error) {
	ctx, resp, err := s.compare.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	rs, ok := resp.(*pb.CardCompareResponse)
	if !ok {
		return nil, fmt.Errorf("expecting response of type *pb.CardCompareResponse but got %#v", rs)
	}
	return rs, nil
}
