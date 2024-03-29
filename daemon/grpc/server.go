package srvgrpc

import (
	"context"
	"fmt"
	"log"
	"net"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	grpcmdw "github.com/mcaci/briscola-serv/daemon/grpc/mdw"
	briscola "github.com/mcaci/briscola-serv/daemon/lib"
	"github.com/mcaci/briscola-serv/pb"
	"google.golang.org/grpc"
)

type srv struct {
	points  grpctransport.Handler
	count   grpctransport.Handler
	compare grpctransport.Handler
}

func NewServer(mdws ...func(grpctransport.Handler) grpctransport.Handler) *srv {
	s := &srv{
		points:  grpctransport.NewServer(briscola.PointsEP, pointsRequestDecode, pointsResponseEncode),
		count:   grpctransport.NewServer(briscola.CountEP, countRequestDecode, countResponseEncode),
		compare: grpctransport.NewServer(briscola.CompareEP, compareRequestDecode, compareResponseEncode),
	}
	if len(mdws) == 0 {
		mdws = append(mdws, grpcmdw.Logged)
	}
	for _, mdw := range mdws {
		s.points = mdw(s.points)
		s.count = mdw(s.count)
		s.compare = mdw(s.compare)
	}
	return s
}

func (s *srv) Start(ctx context.Context, addr string) <-chan error {
	if addr == "" {
		addr = ":8081"
	}
	errChan := make(chan error)
	go func() {
		<-ctx.Done()
		errChan <- ctx.Err()
	}()
	go func() {
		log.Println("listenning to grpc requests on", addr)
		srv := grpc.NewServer()
		pb.RegisterBriscolaServer(srv, s)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			errChan <- err
			return
		}
		defer listener.Close()
		errChan <- srv.Serve(listener)
	}()
	return errChan
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
