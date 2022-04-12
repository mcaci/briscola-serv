package endp

import (
	"context"

	"github.com/mcaci/briscola-serv/pb"
)

func PointsRequestEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(PointsRequest)
	return &pb.CardPointsRequest{CardNumber: req.CardNumber}, nil
}

func PointsRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardPointsRequest)
	return PointsRequest{CardNumber: req.CardNumber}, nil
}

func PointsResponseEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(PointsResponse)
	return &pb.CardPointsResponse{Points: res.Points}, nil
}

func PointsResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardPointsResponse)
	return PointsResponse{Points: res.Points, Err: ""}, nil
}

func CountRequestEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(CountRequest)
	return &pb.PointCountRequest{CardNumber: req.CardNumbers}, nil
}

func CountRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PointCountRequest)
	return CountRequest{CardNumbers: req.CardNumber}, nil
}

func CountResponseEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(CountResponse)
	return &pb.PointCountResponse{Count: res.Points}, nil
}

func CountResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.PointCountResponse)
	return CountResponse{Points: res.Count, Err: ""}, nil
}

func CompareRequestEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(CompareRequest)
	first := &pb.ItalianCard{Number: req.FirstCardNumber, Seed: pb.Seed(req.FirstCardSeed)}
	second := &pb.ItalianCard{Number: req.SecondCardNumber, Seed: pb.Seed(req.SecondCardSeed)}
	return &pb.CardCompareRequest{FirstCard: first, SecondCard: second, Briscola: pb.Seed(req.BriscolaSeed)}, nil
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

func CompareResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardCompareResponse)
	return CompareResponse{SecondCardWins: res.SecondCardWinsOverFirstOne, Err: ""}, nil
}
