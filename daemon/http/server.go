package briscolahttp

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	endp "github.com/mcaci/briscola-serv/daemon/endpoint"
)

func NewHTTPServer(ctx context.Context, endpoints endp.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/points", httptransport.NewServer(endpoints.CardPointsEndpoint, RequestDecodeJSON[endp.PointsRequest], ResponseEncodeJSON))
	m.Handle("/count", httptransport.NewServer(endpoints.PointCountEndpoint, RequestDecodeJSON[endp.CountRequest], ResponseEncodeJSON))
	m.Handle("/compare", httptransport.NewServer(endpoints.CardCompareEndpoint, RequestDecodeJSON[endp.CompareRequest], ResponseEncodeJSON))
	return m
}
