package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	httpmdw "github.com/mcaci/briscola-serv/daemon/http/mdw"
	briscola "github.com/mcaci/briscola-serv/daemon/lib"
)

func NewHandler(ctx context.Context, mdws ...func(http.Handler) http.Handler) http.Handler {
	var pointsHnd http.Handler = httptransport.NewServer(
		briscola.PointsEP,
		requestDecode[struct {
			CardNumber uint32 `json:"number"`
		}],
		responseEncode,
	)
	var countHnd http.Handler = httptransport.NewServer(
		briscola.CountEP,
		requestDecode[struct {
			CardNumbers []uint32 `json:"numbers"`
		}],
		responseEncode,
	)
	var compareHnd http.Handler = httptransport.NewServer(
		briscola.CompareEP,
		requestDecode[struct {
			FirstCardNumber  uint32 `json:"firstCardNumber"`
			FirstCardSeed    uint32 `json:"firstCardSeed"`
			SecondCardNumber uint32 `json:"secondCardNumber"`
			SecondCardSeed   uint32 `json:"secondCardSeed"`
			BriscolaSeed     uint32 `json:"briscolaSeed"`
		}],
		responseEncode,
	)

	if len(mdws) == 0 {
		mdws = append(mdws, httpmdw.Logged)
	}
	for _, mdw := range mdws {
		pointsHnd = mdw(pointsHnd)
		countHnd = mdw(countHnd)
		compareHnd = mdw(compareHnd)
	}

	m := http.NewServeMux()
	m.Handle("/points", pointsHnd)
	m.Handle("/count", countHnd)
	m.Handle("/compare", compareHnd)
	return m
}

func requestDecode[T any](ctx context.Context, r *http.Request) (interface{}, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func responseEncode(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
