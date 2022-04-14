package cli

import (
	"context"
	"errors"
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/briscola-serv/pb"
	"google.golang.org/grpc"
)

type Opts struct {
	GRPCAddr string
}

func Start(o *Opts) {
	ctx := context.Background()
	conn, err := grpc.Dial(o.GRPCAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()
	srv := newGRPCClientService(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "points":
		var number string
		number, args = pop(args)
		n, _ := strconv.Atoi(number)
		res, err := srv.CardPoints(ctx, uint32(n))
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(res)
	case "count":
		var numbers []uint32
		for _, arg := range args {
			n, _ := strconv.Atoi(arg)
			numbers = append(numbers, uint32(n))
		}
		res, err := srv.PointCount(ctx, numbers)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(res)
	case "compare":
		var number string
		number, args = pop(args)
		fcnum, _ := strconv.Atoi(number)
		number, args = pop(args)
		fcseed, _ := strconv.Atoi(number)
		number, args = pop(args)
		scnum, _ := strconv.Atoi(number)
		number, args = pop(args)
		scseed, _ := strconv.Atoi(number)
		number, args = pop(args)
		brseed, _ := strconv.Atoi(number)
		res, err := srv.CardCompare(ctx, uint32(fcnum), uint32(fcseed), uint32(scnum), uint32(scseed), uint32(brseed))
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(res)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

type services interface {
	CardPoints(ctx context.Context, number uint32) (uint32, error)
	PointCount(ctx context.Context, number []uint32) (uint32, error)
	CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error)
}

func newGRPCClientService(conn *grpc.ClientConn) services {
	pointsEndpoint := newClientPointsEndpoint(conn)
	countEndpoint := newClientCountEndpoint(conn)
	compareEndpoint := newClientCompareEndpoint(conn)
	return endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}

func newClientPointsEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "CardPoints",
		PointsRequestEncodeGRPC, PointsResponseDecodeGRPC,
		pb.CardPointsResponse{},
	).Endpoint()
}

func newClientCountEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "PointCount",
		CountRequestEncodeGRPC, CountResponseDecodeGRPC,
		pb.PointCountResponse{},
	).Endpoint()
}

func newClientCompareEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn, "pb.Briscola", "CardCompare",
		CompareRequestEncodeGRPC, CompareResponseDecodeGRPC,
		pb.CardCompareResponse{},
	).Endpoint()
}

type endpoints struct {
	CardPointsEndpoint  endpoint.Endpoint
	PointCountEndpoint  endpoint.Endpoint
	CardCompareEndpoint endpoint.Endpoint
}

func (e endpoints) CardPoints(ctx context.Context, number uint32) (uint32, error) {
	req := struct {
		CardNumber uint32 `json:"number"`
	}{CardNumber: number}
	resp, err := e.CardPointsEndpoint(ctx, req)
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

func (e endpoints) PointCount(ctx context.Context, card_numbers []uint32) (uint32, error) {
	req := struct {
		CardNumbers []uint32 `json:"numbers"`
	}{CardNumbers: card_numbers}
	resp, err := e.PointCountEndpoint(ctx, req)
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

func (e endpoints) CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error) {
	req := struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}{firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed}
	resp, err := e.CardCompareEndpoint(ctx, req)
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

func PointsRequestEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(struct {
		CardNumber uint32 `json:"number"`
	})
	return &pb.CardPointsRequest{CardNumber: req.CardNumber}, nil
}

func PointsResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardPointsResponse)
	return struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	}{Points: res.Points, Err: ""}, nil
}

func CountRequestEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(struct {
		CardNumbers []uint32 `json:"numbers"`
	})
	return &pb.PointCountRequest{CardNumber: req.CardNumbers}, nil
}

func CountResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.PointCountResponse)
	return struct {
		Points uint32 `json:"points"`
		Err    string `json:"err,omitempty"`
	}{Points: res.Count, Err: ""}, nil
}

func CompareRequestEncodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	})
	first := &pb.ItalianCard{Number: req.FirstCardNumber, Seed: pb.Seed(req.FirstCardSeed)}
	second := &pb.ItalianCard{Number: req.SecondCardNumber, Seed: pb.Seed(req.SecondCardSeed)}
	return &pb.CardCompareRequest{FirstCard: first, SecondCard: second, Briscola: pb.Seed(req.BriscolaSeed)}, nil
}

func CompareResponseDecodeGRPC(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardCompareResponse)
	return struct {
		SecondCardWins bool   `json:"secondCardWins"`
		Err            string `json:"err,omitempty"`
	}{SecondCardWins: res.SecondCardWinsOverFirstOne, Err: ""}, nil
}
