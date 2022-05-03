package srvgrpc_test

import (
	"context"
	"errors"
	"testing"

	// srvgrpc "github.com/mcaci/briscola-serv/daemon/grpc"

	srvgrpc "github.com/mcaci/briscola-serv/daemon/grpc"
	"github.com/mcaci/briscola-serv/pb"
)

type erroredContext struct{ context.Context }

func (erroredContext) Err() error { return errors.New("this context always errs") }

type mockPointsGRPCHandler struct {
	ctx context.Context
	rs  interface{}
	err error
}

func (m mockPointsGRPCHandler) ServeGRPC(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	if m.ctx != nil {
		return m.ctx, m.rs, m.err
	}
	return ctx, m.rs, m.err
}

func TestNewServerPoints(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		rs        interface{}
		err       error
		expectErr bool
	}{
		{name: "no errors"},
		{name: "wrong rs type", rs: struct{}{}, expectErr: true},
		{name: "response is error", err: errors.New("an error occurred"), expectErr: true},
		{name: "error from context", ctx: erroredContext{}, expectErr: true},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ts := srvgrpc.NewServer()
			ts.Start(context.TODO(), "")
			if tc.expectErr {
				ts.SetPointsHandler(mockPointsGRPCHandler{ctx: tc.ctx, rs: tc.rs, err: tc.err})
			}
			res, err := ts.CardPoints(context.TODO(), &pb.CardPointsRequest{CardNumber: 1})
			if !tc.expectErr && err != nil {
				t.Fatal(err)
			}
			if tc.expectErr && err == nil {
				t.Fatal("I was expecting an error here but got nil")
			}
			t.Log(res)
		})
	}
}

type mockCountGRPCHandler struct{ err error }

func (m mockCountGRPCHandler) ServeGRPC(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	return ctx, &pb.PointCountResponse{}, m.err
}

func TestNewServerCount(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		rs        interface{}
		err       error
		expectErr bool
	}{
		{name: "no errors"},
		{name: "wrong rs type", rs: struct{}{}, expectErr: true},
		{name: "response is error", err: errors.New("an error occurred"), expectErr: true},
		{name: "error from context", ctx: erroredContext{}, expectErr: true},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ts := srvgrpc.NewServer()
			ts.Start(context.TODO(), "")
			if tc.expectErr {
				ts.SetCountHandler(mockPointsGRPCHandler{ctx: tc.ctx, rs: tc.rs, err: tc.err})
			}
			res, err := ts.PointCount(context.TODO(), &pb.PointCountRequest{CardNumber: []uint32{1, 8, 10, 4}})
			if !tc.expectErr && err != nil {
				t.Fatal(err)
			}
			if tc.expectErr && err == nil {
				t.Fatal("I was expecting an error here but got nil")
			}
			t.Log(res)
		})
	}
}

type mockCompareGRPCHandler struct{ err error }

func (m mockCompareGRPCHandler) ServeGRPC(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	return ctx, &pb.CardCompareResponse{}, m.err
}

func TestNewServerCompare(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		rs        interface{}
		err       error
		expectErr bool
	}{
		{name: "no errors"},
		{name: "wrong rs type", rs: struct{}{}, expectErr: true},
		{name: "response is error", err: errors.New("an error occurred"), expectErr: true},
		{name: "error from context", ctx: erroredContext{}, expectErr: true},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ts := srvgrpc.NewServer()
			ts.Start(context.TODO(), "")
			if tc.expectErr {
				ts.SetCompareHandler(mockPointsGRPCHandler{ctx: tc.ctx, rs: tc.rs, err: tc.err})
			}
			res, err := ts.CardCompare(context.TODO(), &pb.CardCompareRequest{FirstCard: &pb.ItalianCard{Number: 1, Seed: pb.Seed_CUP}, SecondCard: &pb.ItalianCard{Number: 3, Seed: pb.Seed_COIN}, Briscola: pb.Seed_COIN})
			if !tc.expectErr && err != nil {
				t.Fatal(err)
			}
			if tc.expectErr && err == nil {
				t.Fatal("I was expecting an error here but got nil")
			}
			t.Log(res)
		})
	}
}
