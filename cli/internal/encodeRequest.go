package internal

import (
	"context"

	"github.com/mcaci/briscola-serv/pb"
)

func cpRqDec(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(struct {
		CardNumber uint32 `json:"number"`
	})
	return &pb.CardPointsRequest{CardNumber: req.CardNumber}, nil
}

func pcRqDec(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(struct {
		CardNumbers []uint32 `json:"numbers"`
	})
	return &pb.PointCountRequest{CardNumber: req.CardNumbers}, nil
}

func ccRqDec(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	})
	first := &pb.ItalianCard{Number: req.FirstCardNumber, Seed: pb.Seed(req.FirstCardSeed)}
	second := &pb.ItalianCard{Number: req.SecondCardNumber, Seed: pb.Seed(req.SecondCardSeed)}
	return &pb.CardCompareRequest{FirstCard: first, SecondCard: second, Briscola: pb.Seed(req.BriscolaSeed)}, nil
}
