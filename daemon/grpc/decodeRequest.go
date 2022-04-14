package briscola

import (
	"context"

	"github.com/mcaci/briscola-serv/pb"
)

func PointsRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardPointsRequest)
	return struct {
		CardNumber uint32 `json:"number"`
	}{CardNumber: req.CardNumber}, nil
}

func CountRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PointCountRequest)
	return struct {
		CardNumbers []uint32 `json:"numbers"`
	}{CardNumbers: req.CardNumber}, nil
}

func CompareRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardCompareRequest)
	return struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}{
		FirstCardNumber:  req.FirstCard.GetNumber(),
		FirstCardSeed:    uint32(req.FirstCard.GetSeed()),
		SecondCardNumber: req.SecondCard.GetNumber(),
		SecondCardSeed:   uint32(req.SecondCard.GetSeed()),
		BriscolaSeed:     uint32(req.Briscola)}, nil
}
