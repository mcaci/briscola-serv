package briscola

import (
	"context"
	"fmt"
)

func PointsEP(ctx context.Context, request interface{}) (response interface{}, err error) {
	switch req := request.(type) {
	case struct {
		CardNumber uint32 `json:"number"`
	}:
		v, err := Points(req.CardNumber), error(nil)
		if err != nil {
			return struct {
				Points uint32 `json:"points"`
				Err    string `json:"err,omitempty"`
			}{Err: err.Error()}, err
		}
		return struct {
			Points uint32 `json:"points"`
			Err    string `json:"err,omitempty"`
		}{Points: v}, nil
	default:
		return nil, fmt.Errorf("req %#v is not supported.", req)
	}
}

func CountEP(ctx context.Context, request interface{}) (response interface{}, err error) {
	switch req := request.(type) {
	case struct {
		CardNumbers []uint32 `json:"numbers"`
	}:
		v, err := Count(req.CardNumbers), error(nil)
		if err != nil {
			return struct {
				Points uint32 `json:"points"`
				Err    string `json:"err,omitempty"`
			}{Err: err.Error()}, err
		}
		return struct {
			Points uint32 `json:"points"`
			Err    string `json:"err,omitempty"`
		}{Points: v}, nil
	default:
		return nil, fmt.Errorf("req %#v is not supported.", req)
	}
}

func CompareEP(ctx context.Context, request interface{}) (response interface{}, err error) {
	switch req := request.(type) {
	case struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}:
		v, err := IsOtherWinning(req.FirstCardNumber, req.FirstCardSeed, req.SecondCardNumber, req.SecondCardSeed, req.BriscolaSeed), error(nil)
		if err != nil {
			return struct {
				SecondCardWins bool   `json:"secondCardWins"`
				Err            string `json:"err,omitempty"`
			}{Err: err.Error()}, err
		}
		return struct {
			SecondCardWins bool   `json:"secondCardWins"`
			Err            string `json:"err,omitempty"`
		}{SecondCardWins: v}, nil
	default:
		return nil, fmt.Errorf("req %#v is not supported.", req)
	}
}
