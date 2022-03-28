package briscolahttp

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	endp "github.com/mcaci/briscola-serv/endpoint"
)

func NewHTTPServer(ctx context.Context, endpoints endp.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/points", httptransport.NewServer(endpoints.CardPointsEndpoint, endp.PointsRequestDecodeJSON, endp.PointsResponseEncodeJSON))
	m.Handle("/count", httptransport.NewServer(endpoints.PointCountEndpoint, endp.CountRequestDecodeJSON, endp.CountResponseEncodeJSON))
	m.Handle("/compare", httptransport.NewServer(endpoints.CardCompareEndpoint, endp.CompareRequestDecodeJSON, endp.CompareResponseEncodeJSON))
	return m
}
