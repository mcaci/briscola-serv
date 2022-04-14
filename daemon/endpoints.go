package daemon

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	briscolahttp "github.com/mcaci/briscola-serv/daemon/http"
)

type Endpoints struct {
	CardPointsEndpoint  endpoint.Endpoint
	PointCountEndpoint  endpoint.Endpoint
	CardCompareEndpoint endpoint.Endpoint
}

func (e Endpoints) Cp() endpoint.Endpoint { return e.CardPointsEndpoint }
func (e Endpoints) Pc() endpoint.Endpoint { return e.PointCountEndpoint }
func (e Endpoints) Cc() endpoint.Endpoint { return e.CardCompareEndpoint }

func (e Endpoints) Cpd() httptransport.DecodeRequestFunc {
	return briscolahttp.RequestDecodeJSON[struct {
		CardNumber uint32 `json:"number"`
	}]
}
func (e Endpoints) Pcd() httptransport.DecodeRequestFunc {
	return briscolahttp.RequestDecodeJSON[struct {
		CardNumbers []uint32 `json:"numbers"`
	}]
}
func (e Endpoints) Ccd() httptransport.DecodeRequestFunc {
	return briscolahttp.RequestDecodeJSON[struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}]
}

func (e Endpoints) CardPoints(ctx context.Context, number uint32) (uint32, error) {
	req := struct {
		CardNumber uint32 `json:"number"`
	}{CardNumber: number}
	resp, err := e.CardPointsEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

func (e Endpoints) PointCount(ctx context.Context, card_numbers []uint32) (uint32, error) {
	req := struct {
		CardNumbers []uint32 `json:"numbers"`
	}{CardNumbers: card_numbers}
	resp, err := e.PointCountEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

func (e Endpoints) CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error) {
	req := struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}{firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed}
	resp, err := e.CardCompareEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	compareResp := resp.(struct {
		SecondCardWins bool   `json:"secondCardWins"`
		Err            string `json:"err,omitempty"`
	})
	if compareResp.Err != "" {
		return false, errors.New(compareResp.Err)
	}
	return compareResp.SecondCardWins, nil
}
