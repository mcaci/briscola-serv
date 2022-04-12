package briscolahttp

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	endp "github.com/mcaci/briscola-serv/daemon/endpoint"
)

func NewHTTPServer(ctx context.Context, endpoints endp.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/points", httptransport.NewServer(endpoints.CardPointsEndpoint, endp.RequestDecodeJSON[endp.PointsRequest], endp.ResponseEncodeJSON))
	m.Handle("/count", httptransport.NewServer(endpoints.PointCountEndpoint, endp.RequestDecodeJSON[endp.CountRequest], endp.ResponseEncodeJSON))
	m.Handle("/compare", httptransport.NewServer(endpoints.CardCompareEndpoint, endp.RequestDecodeJSON[endp.CompareRequest], endp.ResponseEncodeJSON))
	return m
}
