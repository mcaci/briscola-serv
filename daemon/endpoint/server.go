package endp

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

func NewPointsEndpoint(srv interface {
	CardPoints(ctx context.Context, number uint32) (uint32, error)
}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
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
}

func NewCountEndpoint(srv interface {
	PointCount(ctx context.Context, number []uint32) (uint32, error)
}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
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
}

func NewCompareEndpoint(srv interface {
	CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error)
}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
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
}
