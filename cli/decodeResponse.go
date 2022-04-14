package cli

import (
	"context"

	"github.com/mcaci/briscola-serv/pb"
)

func PointsResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardPointsResponse)
	return struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	}{Points: res.Points, Err: ""}, nil
}

func CountResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.PointCountResponse)
	return struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	}{Points: res.Count, Err: ""}, nil
}

func CompareResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardCompareResponse)
	return struct {
		SecondCardWins bool   `json:"secondCardWins"`
		Err            string `json:"err,omitempty"`
	}{SecondCardWins: res.SecondCardWinsOverFirstOne, Err: ""}, nil
}
