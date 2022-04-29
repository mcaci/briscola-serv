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

type pcEP struct {
	cardNumbers []uint32
}

func Count(args []string) (*pcEP, error) {
	inRange := func(n uint32) bool { return n >= 1 && n <= 10 }
	var numbers []uint32
	for i := range args {
		n, err := strconv.Atoi(args[i])
		if err != nil {
			return nil, err
		}
		if !inRange(uint32(n)) {
			return nil, fmt.Errorf("value %d in position %d is not in the correct range: [1-10] for position 0 and 2, [1-4] for the rest", n, i)
		}
		numbers = append(numbers, uint32(n))
	}
	return &pcEP{
		cardNumbers: numbers,
	}, nil
}

func (ep pcEP) Run(ctx context.Context, conn *grpc.ClientConn) (any, error) {
	epCall := func(context.Context, interface{}) (interface{}, error) {
		return struct {
			Points uint32 `json:"points"`
			Err    string `json:"err,omitempty"`
		}{}, nil
	}
	if conn != nil {
		epCall = grpctransport.NewClient(
			conn, "pb.Briscola", "PointCount",
			pcRqDec, pcRsDec,
			pb.PointCountResponse{},
		).Endpoint()
	}
	req := struct {
		CardNumbers []uint32 `json:"numbers"`
	}{CardNumbers: ep.cardNumbers}
	resp, err := epCall(ctx, req)
	if err != nil {
		return 0, err
	}
	countResp, ok := resp.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	if !ok {
		return false, fmt.Errorf(`invalid types at response time: want %T, got %T`, countResp, resp)
	}
	if countResp.Err != "" {
		return 0, errors.New(countResp.Err)
	}
	return countResp.Points, nil
}
