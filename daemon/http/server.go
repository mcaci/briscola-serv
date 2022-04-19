package srvhttp

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type httpEP interface {
	Cp() endpoint.Endpoint
	Pc() endpoint.Endpoint
	Cc() endpoint.Endpoint
}

type jsonDec interface {
	CpRqDec() httptransport.DecodeRequestFunc
	PcRqDec() httptransport.DecodeRequestFunc
	CcRqDec() httptransport.DecodeRequestFunc
}

func NewHTTPServer(ctx context.Context, ep httpEP, jd jsonDec) http.Handler {
	m := http.NewServeMux()
	m.Handle("/points", httptransport.NewServer(ep.Cp(), jd.CpRqDec(), responseEncode))
	m.Handle("/count", httptransport.NewServer(ep.Pc(), jd.PcRqDec(), responseEncode))
	m.Handle("/compare", httptransport.NewServer(ep.Cc(), jd.CcRqDec(), responseEncode))
	return m
}



type srv struct {
	h    http.Handler
	addr string
	errC chan<- (error)
}

func (s srv) Start() {
	// srvx.handler := NewHTTPServer(s.ctx, s.endpoints, s.endpoints)
	s.errC <- http.ListenAndServe(s.addr, s.h)
}
