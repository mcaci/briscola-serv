package endp

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/briscola-serv/pb"
	serv "github.com/mcaci/briscola-serv/service"
	"google.golang.org/grpc"
)

type Endpoints struct {
	CardPointsEndpoint  endpoint.Endpoint
	PointCountEndpoint  endpoint.Endpoint
	CardCompareEndpoint endpoint.Endpoint
}

func NewEndpoints(srv serv.Service) Endpoints {
	pointsEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		switch req := request.(type) {
		case PointsRequest:
			v, err := srv.CardPoints(ctx, req.CardNumber)
			if err != nil {
				return PointsResponse{v, err.Error()}, err
			}
			return PointsResponse{v, ""}, nil
		default:
			return nil, fmt.Errorf("req of type %t is not supported.", req)
		}
	}
	countEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		switch req := request.(type) {
		case CountRequest:
			v, err := srv.PointCount(ctx, req.CardNumbers)
			if err != nil {
				return CountResponse{v, err.Error()}, err
			}
			return CountResponse{v, ""}, nil
		default:
			return nil, fmt.Errorf("req of type %t is not supported.", req)
		}
	}
	compareEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		switch req := request.(type) {
		case CompareRequest:
			v, err := srv.CardCompare(ctx, req.FirstCardNumber, req.FirstCardSeed, req.SecondCardNumber, req.SecondCardSeed, req.BriscolaSeed)
			if err != nil {
				return CompareResponse{v, err.Error()}, err
			}
			return CompareResponse{v, ""}, nil
		default:
			return nil, fmt.Errorf("req of type %t is not supported.", req)
		}
	}
	return Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}

func NewGRPCClient(conn *grpc.ClientConn) Endpoints {
	pointsEndpoint := grpctransport.NewClient(
		conn, "pb.Briscola", "CardPoints",
		PointsRequestEncodeGRPC, PointsResponseDecodeGRPC,
		pb.CardPointsResponse{},
	).Endpoint()
	countEndpoint := grpctransport.NewClient(
		conn, "pb.Briscola", "PointCount",
		CountRequestEncodeGRPC, CountResponseDecodeGRPC,
		pb.PointCountResponse{},
	).Endpoint()
	compareEndpoint := grpctransport.NewClient(
		conn, "pb.Briscola", "CardCompare",
		CompareRequestEncodeGRPC, CompareResponseDecodeGRPC,
		pb.CardCompareResponse{},
	).Endpoint()
	return Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
