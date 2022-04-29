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

type ccEP struct {
	firstCardNumber  uint32
	firstCardSeed    uint32
	secondCardNumber uint32
	secondCardSeed   uint32
	briscolaSeed     uint32
}

func Compare(args []string) (*ccEP, error) {
	if len(args) != 5 {
		return nil, fmt.Errorf("need to provide five arguments: current list is %v", args)
	}
	inRange := func(min, max uint32) func(uint32) bool { return func(n uint32) bool { return n >= min && n <= max } }
	inRangeArgs := []func(uint32) bool{inRange(1, 10), inRange(1, 4), inRange(1, 10), inRange(1, 4), inRange(1, 4)}
	var nArgs []uint32
	for i := range args {
		n, err := strconv.Atoi(args[i])
		if err != nil {
			return nil, err
		}
		if !inRangeArgs[i](uint32(n)) {
			return nil, fmt.Errorf("value %d in position %d is not in the correct range: [1-10] for position 0 and 2, [1-4] for the rest", n, i)
		}
		nArgs = append(nArgs, uint32(n))
	}
	return &ccEP{
		firstCardNumber:  nArgs[0],
		firstCardSeed:    nArgs[1],
		secondCardNumber: nArgs[2],
		secondCardSeed:   nArgs[3],
		briscolaSeed:     nArgs[4],
	}, nil
}

func (ep ccEP) Run(ctx context.Context, conn *grpc.ClientConn) (any, error) {
	epCall := func(context.Context, interface{}) (interface{}, error) {
		return struct {
			SecondCardWins bool   `json:"secondCardWins"`
			Err            string `json:"err,omitempty"`
		}{}, nil
	}
	if conn != nil {
		epCall = grpctransport.NewClient(
			conn, "pb.Briscola", "CardCompare",
			ccRqDec, ccRsDec,
			pb.CardCompareResponse{},
		).Endpoint()
	}
	req := struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}{ep.firstCardNumber, ep.firstCardSeed, ep.secondCardNumber, ep.secondCardSeed, ep.briscolaSeed}
	resp, err := epCall(ctx, req)
	if err != nil {
		return false, err
	}
	compareResp, ok := resp.(struct {
		SecondCardWins bool   `json:"secondCardWins"`
		Err            string `json:"err,omitempty"`
	})
	if !ok {
		return false, fmt.Errorf(`invalid types at response time: want %T, got %T`, compareResp, resp)
	}
	if compareResp.Err != "" {
		return false, errors.New(compareResp.Err)
	}
	return compareResp.SecondCardWins, nil
}
