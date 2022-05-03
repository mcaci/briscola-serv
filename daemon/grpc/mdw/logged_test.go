package grpcmdw_test

import (
	"context"
	"testing"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	grpcmdw "github.com/mcaci/briscola-serv/daemon/grpc/mdw"
)

func TestLoggedMdw(t *testing.T) {
	ts := grpctransport.NewServer(endpoint.Nop, endpoint.Nop, endpoint.Nop)
	logHndl := grpcmdw.Logged(ts)
	ctx, res, err := logHndl.ServeGRPC(context.TODO(), struct{}{})
	if ctx.Err() != nil {
		t.Fatal(ctx.Err())
	}
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
