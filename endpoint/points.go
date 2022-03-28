package endp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mcaci/briscola-serv/pb"
)

func (e Endpoints) CardPoints(ctx context.Context, number uint32) (uint32, error) {
	req := PointsRequest{CardNumber: number}
	resp, err := e.CardPointsEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(PointsResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

type PointsRequest struct {
	CardNumber uint32 `json:"number"`
}

func PointsRequestDecodeJSON(ctx context.Context, r *http.Request) (interface{}, error) {
	var req PointsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func PointsRequestEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(PointsRequest)
	return &pb.CardPointsRequest{CardNumber: req.CardNumber}, nil
}

func PointsRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardPointsRequest)
	return PointsRequest{CardNumber: req.CardNumber}, nil
}

type PointsResponse struct {
	Points uint32 `json:"points"`
	Err    string `json:"err,omitempty"`
}

func PointsResponseEncodeJSON(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func PointsResponseEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(PointsResponse)
	return &pb.CardPointsResponse{Points: res.Points}, nil
}

func PointsResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardPointsResponse)
	return PointsResponse{Points: res.Points, Err: ""}, nil
}
