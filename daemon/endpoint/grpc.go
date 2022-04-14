package endp

import (
	"context"

	"github.com/mcaci/briscola-serv/pb"
)

func PointsRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardPointsRequest)
	return PointsRequest{CardNumber: req.CardNumber}, nil
}

func PointsResponseEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(PointsResponse)
	return &pb.CardPointsResponse{Points: res.Points}, nil
}

func CountRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PointCountRequest)
	return CountRequest{CardNumbers: req.CardNumber}, nil
}

func CountResponseEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(CountResponse)
	return &pb.PointCountResponse{Count: res.Points}, nil
}

func CompareRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardCompareRequest)
	return CompareRequest{
		FirstCardNumber:  req.FirstCard.GetNumber(),
		FirstCardSeed:    uint32(req.FirstCard.GetSeed()),
		SecondCardNumber: req.SecondCard.GetNumber(),
		SecondCardSeed:   uint32(req.SecondCard.GetSeed()),
		BriscolaSeed:     uint32(req.Briscola)}, nil
}

func CompareResponseEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(CompareResponse)
	return &pb.CardCompareResponse{SecondCardWinsOverFirstOne: res.SecondCardWins}, nil
}
