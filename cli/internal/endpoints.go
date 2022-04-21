package internal

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/briscola-serv/pb"
	"google.golang.org/grpc"
)

type cpEP struct {
	number uint32
	call   endpoint.Endpoint
}

func (ep *cpEP) SetEndpoint(conn *grpc.ClientConn) {
	ep.call = grpctransport.NewClient(
		conn, "pb.Briscola", "CardPoints",
		cpRqDec, cpRsDec,
		pb.CardPointsResponse{},
	).Endpoint()
}

func (ep *cpEP) Run(ctx context.Context) (any, error) {
	req := struct {
		CardNumber uint32 `json:"number"`
	}{CardNumber: ep.number}
	resp, err := ep.call(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

type pcEP struct {
	cardNumbers []uint32
	call        endpoint.Endpoint
}

func (ep *pcEP) SetEndpoint(conn *grpc.ClientConn) {
	ep.call = grpctransport.NewClient(
		conn, "pb.Briscola", "PointCount",
		pcRqDec, pcRsDec,
		pb.PointCountResponse{},
	).Endpoint()
}

func (ep pcEP) Run(ctx context.Context) (any, error) {
	req := struct {
		CardNumbers []uint32 `json:"numbers"`
	}{CardNumbers: ep.cardNumbers}
	resp, err := ep.call(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	})
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

type ccEP struct {
	firstCardNumber  uint32
	firstCardSeed    uint32
	secondCardNumber uint32
	secondCardSeed   uint32
	briscolaSeed     uint32
	call             endpoint.Endpoint
}

func (ep *ccEP) SetEndpoint(conn *grpc.ClientConn) {
	ep.call = grpctransport.NewClient(
		conn, "pb.Briscola", "CardCompare",
		ccRqDec, ccRsDec,
		pb.CardCompareResponse{},
	).Endpoint()
}

func (ep ccEP) Run(ctx context.Context) (any, error) {
	req := struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}{ep.firstCardNumber, ep.firstCardSeed, ep.secondCardNumber, ep.secondCardSeed, ep.briscolaSeed}
	resp, err := ep.call(ctx, req)
	if err != nil {
		return false, err
	}
	compareResp := resp.(struct {
		SecondCardWins bool   `json:"secondCardWins"`
		Err            string `json:"err,omitempty"`
	})
	if compareResp.Err != "" {
		return false, errors.New(compareResp.Err)
	}
	return compareResp.SecondCardWins, nil
}
