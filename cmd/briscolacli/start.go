package cli

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	endp "github.com/mcaci/briscola-serv/endpoint"
	serv "github.com/mcaci/briscola-serv/service"
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
	cli := endp.NewGRPCClient(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "points":
		var number string
		number, args = pop(args)
		n, _ := strconv.Atoi(number)
		points(ctx, cli, uint32(n))
	case "count":
		var numbers []uint32
		for _, arg := range args {
			n, _ := strconv.Atoi(arg)
			numbers = append(numbers, uint32(n))
		}
		count(ctx, cli, numbers)
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
		compare(ctx, cli, uint32(fcnum), uint32(fcseed), uint32(scnum), uint32(scseed), uint32(brseed))
	default:
		log.Fatalln("unknown command", cmd)
	}
}

func points(ctx context.Context, service serv.Service, number uint32) {
	res, err := service.CardPoints(ctx, number)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(res)
}

func count(ctx context.Context, service serv.Service, numbers []uint32) {
	res, err := service.PointCount(ctx, numbers)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(res)
}

func compare(ctx context.Context, service serv.Service, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) {
	res, err := service.CardCompare(ctx, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(res)
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
