package cli

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type Opts struct {
	GRPCAddr string
	Cmd      string
	Args     []string
	EpRun    endpointRunner
}

func Start(o *Opts) error {
	ep := o.EpRun
	var conn *grpc.ClientConn
	if ep == nil {
		var err error
		ep, err = setup(o.Cmd, o.Args)
		if err != nil {
			return err
		}
		conn, err = grpc.Dial(o.GRPCAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
		if err != nil {
			return err
		}
		defer conn.Close()
	}
	ep.SetEndpoint(conn)
	res, err := ep.Run(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

type endpointRunner interface {
	SetEndpoint(conn *grpc.ClientConn)
	Run(ctx context.Context) (any, error)
}

func setup(cmd string, args []string) (endpointRunner, error) {
	switch cmd {
	case "points":
		var number string
		number, args = pop(args)
		n, _ := strconv.Atoi(number)
		return &cpEP{
			number: uint32(n),
		}, nil
	case "count":
		var numbers []uint32
		for _, arg := range args {
			n, _ := strconv.Atoi(arg)
			numbers = append(numbers, uint32(n))
		}
		return &pcEP{
			cardNumbers: numbers,
		}, nil
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
		return &ccEP{
			firstCardNumber:  uint32(fcnum),
			firstCardSeed:    uint32(fcseed),
			secondCardNumber: uint32(scnum),
			secondCardSeed:   uint32(scseed),
			briscolaSeed:     uint32(brseed),
		}, nil
	default:
		return nil, fmt.Errorf("unknown command: %q", cmd)
	}
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
