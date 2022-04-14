package briscolahttp

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type endpointInt interface {
	Cp() endpoint.Endpoint
	Pc() endpoint.Endpoint
	Cc() endpoint.Endpoint
}

type jsonDecoder interface {
	Cpd() httptransport.DecodeRequestFunc
	Pcd() httptransport.DecodeRequestFunc
	Ccd() httptransport.DecodeRequestFunc
}

func NewHTTPServer(ctx context.Context, ep endpointInt, jd jsonDecoder) http.Handler {
	m := http.NewServeMux()
	m.Handle("/points", httptransport.NewServer(ep.Cp(), jd.Cpd(), ResponseEncodeJSON))
	m.Handle("/count", httptransport.NewServer(ep.Pc(), jd.Pcd(), ResponseEncodeJSON))
	m.Handle("/compare", httptransport.NewServer(ep.Cc(), jd.Ccd(), ResponseEncodeJSON))
	return m
}
