package grpcmdw

import (
	"context"
	"log"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// look at
// endpoint.Chain() from go-kit

func Logged(handler grpctransport.Handler) grpctransport.Handler {
	return &logged{handler: handler}
}

type logged struct {
	handler grpctransport.Handler
}

func (l *logged) ServeGRPC(ctx context.Context, request interface{}) (context.Context, interface{}, error) {
	log.Printf("handling request <%v> from client", request)
	return l.handler.ServeGRPC(ctx, request)
}
