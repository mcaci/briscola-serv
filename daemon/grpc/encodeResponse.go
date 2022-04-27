package srvgrpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/mcaci/briscola-serv/pb"
)

func pointsResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	switch res := r.(type) {
	case struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	}:
		return &pb.CardPointsResponse{Points: res.Points}, errors.New(res.Err)
	default:
		return nil, fmt.Errorf("cannot encode response for points; got %#v", res)
	}
}

func countResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	switch res := r.(type) {
	case struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	}:
		return &pb.PointCountResponse{Count: res.Points}, errors.New(res.Err)
	default:
		return nil, fmt.Errorf("cannot encode response for count; got %#v", res)
	}
}

func compareResponseEncode(ctx context.Context, r interface{}) (interface{}, error) {
	switch res := r.(type) {
	case struct {
		SecondCardWins bool   `json:"secondCardWins"`
		Err            string `json:"err,omitempty"`
	}:
		return &pb.CardCompareResponse{SecondCardWinsOverFirstOne: res.SecondCardWins}, errors.New(res.Err)
	default:
		return nil, fmt.Errorf("cannot encode response for compare; got %#v", res)
	}
}
