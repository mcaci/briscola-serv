package srvhttp

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	briscola "github.com/mcaci/briscola-serv/daemon/lib"
)

func NewHandler(ctx context.Context) http.Handler {
	m := http.NewServeMux()
	m.Handle(points.pattern, points.handler)
	m.Handle(count.pattern, count.handler)
	m.Handle(compare.pattern, compare.handler)
	return m
}

type handlerData struct {
	pattern string
	handler http.Handler
}

var points = handlerData{pattern: "/points", handler: httptransport.NewServer(
	briscola.PointsEP,
	requestDecode[struct {
		CardNumber uint32 `json:"number"`
	}],
	responseEncode,
)}

var count = handlerData{pattern: "/count", handler: httptransport.NewServer(
	briscola.CountEP,
	requestDecode[struct {
		CardNumbers []uint32 `json:"numbers"`
	}],
	responseEncode,
)}

var compare = handlerData{pattern: "/compare", handler: httptransport.NewServer(
	briscola.CompareEP,
	requestDecode[struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}],
	responseEncode,
)}

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
