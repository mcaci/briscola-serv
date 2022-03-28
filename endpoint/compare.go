package endp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mcaci/briscola-serv/pb"
)

func (e Endpoints) CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error) {
	req := CompareRequest{firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed}
	resp, err := e.CardCompareEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	compareResp := resp.(CompareResponse)
	if compareResp.Err != "" {
		return false, errors.New(compareResp.Err)
	}
	return compareResp.SecondCardWins, nil
}

type CompareRequest struct {
	FirstCardNumber  uint32 `json:"firstCardNumber"`
	FirstCardSeed    uint32 `json:"firstCardSeed"`
	SecondCardNumber uint32 `json:"secondCardNumber"`
	SecondCardSeed   uint32 `json:"secondCardSeed"`
	BriscolaSeed     uint32 `json:"briscolaSeed"`
}

func CompareRequestDecodeJSON(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CompareRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
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

type CompareResponse struct {
	SecondCardWins bool   `json:"secondCardWins"`
	Err            string `json:"err,omitempty"`
}

func CompareResponseEncodeJSON(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func CompareResponseEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(CompareResponse)
	return &pb.CardCompareResponse{SecondCardWinsOverFirstOne: res.SecondCardWins}, nil
}

func CompareResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardCompareResponse)
	return CompareResponse{SecondCardWins: res.SecondCardWinsOverFirstOne, Err: ""}, nil
}
