package briscola

import (
	endp "github.com/mcaci/briscola-serv/endpoint"
	"github.com/mcaci/briscola-serv/pb"
	"golang.org/x/net/context"
)

func (s *grpcServer) PointCount(ctx context.Context, r *pb.PointCountRequest) (*pb.PointCountResponse, error) {
	_, resp, err := s.count.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PointCountResponse), nil
}

func EncodeGRPCCountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(endp.CountRequest)
	return &pb.PointCountRequest{CardNumber: req.CardNumbers}, nil
}

func DecodeGRPCCountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PointCountRequest)
	return endp.CountRequest{CardNumbers: req.CardNumber}, nil
}

func EncodeGRPCCountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(endp.CountResponse)
	return &pb.PointCountResponse{Count: res.Points}, nil
}

func DecodeGRPCCountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.PointCountResponse)
	return endp.CountResponse{Points: res.Count, Err: ""}, nil
}
