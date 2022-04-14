package daemon

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

func newPointsEndpoint(srv interface {
	CardPoints(ctx context.Context, number uint32) (uint32, error)
}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		switch req := request.(type) {
		case struct {
			CardNumber uint32 `json:"number"`
		}:
			v, err := srv.CardPoints(ctx, req.CardNumber)
			if err != nil {
				return struct {
					Points uint32 `json:"points"`
					Err    string `json:"err,omitempty"`
				}{v, err.Error()}, err
			}
			return struct {
				Points uint32 `json:"points"`
				Err    string `json:"err,omitempty"`
			}{v, ""}, nil
		default:
			return nil, fmt.Errorf("req of type %t is not supported.", req)
		}
	}
}

func newCountEndpoint(srv interface {
	PointCount(ctx context.Context, number []uint32) (uint32, error)
}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		switch req := request.(type) {
		case struct {
			CardNumbers []uint32 `json:"numbers"`
		}:
			v, err := srv.PointCount(ctx, req.CardNumbers)
			if err != nil {
				return struct {
					Points uint32 `json:"points"`
					Err    string `json:"err,omitempty"`
				}{v, err.Error()}, err
			}
			return struct {
				Points uint32 `json:"points"`
				Err    string `json:"err,omitempty"`
			}{v, ""}, nil
		default:
			return nil, fmt.Errorf("req of type %t is not supported.", req)
		}
	}
}

func newCompareEndpoint(srv interface {
	CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error)
}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		switch req := request.(type) {
		case struct {
			FirstCardNumber  uint32 `json:"firstCardNumber"`
			FirstCardSeed    uint32 `json:"firstCardSeed"`
			SecondCardNumber uint32 `json:"secondCardNumber"`
			SecondCardSeed   uint32 `json:"secondCardSeed"`
			BriscolaSeed     uint32 `json:"briscolaSeed"`
		}:
			v, err := srv.CardCompare(ctx, req.FirstCardNumber, req.FirstCardSeed, req.SecondCardNumber, req.SecondCardSeed, req.BriscolaSeed)
			if err != nil {
				return struct {
					SecondCardWins bool   `json:"secondCardWins"`
					Err            string `json:"err,omitempty"`
				}{v, err.Error()}, err
			}
			return struct {
				SecondCardWins bool   `json:"secondCardWins"`
				Err            string `json:"err,omitempty"`
			}{v, ""}, nil
		default:
			return nil, fmt.Errorf("req of type %t is not supported.", req)
		}
	}
}
