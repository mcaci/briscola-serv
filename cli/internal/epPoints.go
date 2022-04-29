package internal

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/briscola-serv/pb"
	"google.golang.org/grpc"
)

type cpEP struct {
	number uint32
}

func Points(args []string) (*cpEP, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("need to provide one argument: current list is %v", args)
	}
	inRange := func(n uint32) bool { return n >= 1 && n <= 10 }
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, err
	}
	if !inRange(uint32(n)) {
		return nil, fmt.Errorf("value %d is not in the correct range: [1-10]", n)
	}
	ep := cpEP{
		number: uint32(n),
	}
	return &ep, nil
}

func (ep cpEP) Run(ctx context.Context, conn *grpc.ClientConn) (any, error) {
	epCall := func(context.Context, interface{}) (interface{}, error) {
		return struct {
			Points uint32 `json:"points"`
			Err    string `json:"err,omitempty"`
		}{}, nil
	}
	if conn != nil {
		epCall = grpctransport.NewClient(
			conn, "pb.Briscola", "CardPoints",
			cpRqDec, cpRsDec,
			pb.CardPointsResponse{},
		).Endpoint()
	}
	req := struct {
		CardNumber uint32 `json:"number"`
	}{CardNumber: ep.number}
	resp, err := epCall(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp, ok := resp.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	if !ok {
		return false, fmt.Errorf(`invalid types at response time: want %T, got %T`, pointsResp, resp)
	}
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}
