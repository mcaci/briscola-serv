package srvgrpc

import (
	"context"
	"errors"

	"github.com/mcaci/briscola-serv/pb"
)

func pointsResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	return &pb.CardPointsResponse{Points: res.Points}, errors.New(res.Err)
}

func countResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	return &pb.PointCountResponse{Count: res.Points}, errors.New(res.Err)
}

func compareResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(struct {
		SecondCardWins bool   `json:"secondCardWins"`
		Err            string `json:"err,omitempty"`
	})
	return &pb.CardCompareResponse{SecondCardWinsOverFirstOne: res.SecondCardWins}, errors.New(res.Err)
}
