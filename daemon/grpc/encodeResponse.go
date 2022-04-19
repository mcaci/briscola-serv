package srvgrpc

import (
	"context"

	"github.com/mcaci/briscola-serv/pb"
)

func pointsResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	return &pb.CardPointsResponse{Points: res.Points}, nil
}

func countResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	return &pb.PointCountResponse{Count: res.Points}, nil
}

func compareResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(struct {
		SecondCardWins bool   `json:"secondCardWins"`
		Err            string `json:"err,omitempty"`
	})
	return &pb.CardCompareResponse{SecondCardWinsOverFirstOne: res.SecondCardWins}, nil
}
