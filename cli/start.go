package cli

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	endp "github.com/mcaci/briscola-serv/endpoint"
	"google.golang.org/grpc"
)

func Start() {
	const grpcAddr = ":8081"
	ctx := context.Background()
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
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
		fmt.Println(res)
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
		fmt.Println(res)
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
		fmt.Println(res)
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

func newGRPCClientService(conn *grpc.ClientConn) endp.Endpoints {
	pointsEndpoint := endp.NewClientPointsEndpoint(conn)
	countEndpoint := endp.NewClientCountEndpoint(conn)
	compareEndpoint := endp.NewClientCompareEndpoint(conn)
	return endp.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
