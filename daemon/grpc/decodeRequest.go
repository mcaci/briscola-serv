package srvgrpc

import (
	"context"
	"fmt"

	"github.com/mcaci/briscola-serv/pb"
)

func pointsRequestDecode(ctx context.Context, r interface{}) (interface{}, error) {
	switch req := r.(type) {
	case *pb.CardPointsRequest:
		return struct {
				CardNumber uint32 `json:"number"`
			}{CardNumber: req.CardNumber},
			nil
	default:
		return nil, fmt.Errorf("cannot decode request for points; got %#v", req)
	}
}

func countRequestDecode(ctx context.Context, r interface{}) (interface{}, error) {
	switch req := r.(type) {
	case *pb.PointCountRequest:
		return struct {
				CardNumber []uint32 `json:"numbers"`
			}{CardNumber: req.CardNumber},
			nil
	default:
		return nil, fmt.Errorf("cannot decode request for count; got %#v", req)
	}
}

func compareRequestDecode(ctx context.Context, r interface{}) (interface{}, error) {
	switch req := r.(type) {
	case *pb.CardCompareRequest:
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
				BriscolaSeed:     uint32(req.Briscola)},
			nil
	default:
		return nil, fmt.Errorf("cannot decode request for compare; got %#v", req)
	}
}
