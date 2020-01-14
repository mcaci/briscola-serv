package endp

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	serv "github.com/mcaci/briscola-serv/service"
)

type PointsRequest struct {
	CardNumber uint32 `json:"number"`
}

type CountRequest struct {
	CardNumbers []uint32 `json:"numbers"`
}

type CompareRequest struct {
	FirstCardNumber  uint32 `json:"firstCardNumber"`
	FirstCardSeed    uint32 `json:"firstCardSeed"`
	SecondCardNumber uint32 `json:"secondCardNumber"`
	SecondCardSeed   uint32 `json:"secondCardSeed"`
	BriscolaSeed     uint32 `json:"briscolaSeed"`
}

type PointsResponse struct {
	Points uint32 `json:"points"`
	Err    string `json:"err,omitempty"`
}

type CountResponse PointsResponse

type CompareResponse struct {
	SecondCardWins bool   `json:"secondCardWins"`
	Err            string `json:"err,omitempty"`
}

func MakePointsEndpoint(srv serv.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PointsRequest)
		v, err := srv.CardPoints(ctx, req.CardNumber)
		if err != nil {
			return PointsResponse{v, err.Error()}, nil
		}
		return PointsResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(srv serv.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v, err := srv.PointCount(ctx, req.CardNumbers)
		if err != nil {
			return CountResponse{v, err.Error()}, nil
		}
		return CountResponse{v, ""}, nil
	}
}

func MakeCompareEndpoint(srv serv.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CompareRequest)
		v, err := srv.CardCompare(ctx, req.FirstCardNumber, req.FirstCardSeed, req.SecondCardNumber, req.SecondCardSeed, req.BriscolaSeed)
		if err != nil {
			return CompareResponse{v, err.Error()}, nil
		}
		return CompareResponse{v, ""}, nil
	}
}

type Endpoints struct {
	CardPointsEndpoint  endpoint.Endpoint
	PointCountEndpoint  endpoint.Endpoint
	CardCompareEndpoint endpoint.Endpoint
}

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
