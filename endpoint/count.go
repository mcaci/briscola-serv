package endp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mcaci/briscola-serv/pb"
)

func (e Endpoints) PointCount(ctx context.Context, card_numbers []uint32) (uint32, error) {
	req := CountRequest{CardNumbers: card_numbers}
	resp, err := e.PointCountEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(CountResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

type CountRequest struct {
	CardNumbers []uint32 `json:"numbers"`
}

func CountRequestDecodeJSON(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func CountRequestEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(CountRequest)
	return &pb.PointCountRequest{CardNumber: req.CardNumbers}, nil
}

func CountRequestDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PointCountRequest)
	return CountRequest{CardNumbers: req.CardNumber}, nil
}

type CountResponse PointsResponse

func CountResponseEncodeJSON(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func CountResponseEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(CountResponse)
	return &pb.PointCountResponse{Count: res.Points}, nil
}

func CountResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.PointCountResponse)
	return CountResponse{Points: res.Count, Err: ""}, nil
}
