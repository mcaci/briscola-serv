package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/mcaci/briscola-serv/cli/internal"
	"google.golang.org/grpc"
)

type Opts struct {
	GRPCAddr string
	Cmd      string
	Args     []string
	EpRun    runner
}

type runner interface {
	Run(ctx context.Context, conn *grpc.ClientConn) (any, error)
}

func Start(o *Opts) error {
	var conn *grpc.ClientConn
	var ep = o.EpRun
	if ep == nil {
		var err error
		ep, err = selectEP(o.Cmd, o.Args)
		if err != nil {
			return err
		}
		conn, err = grpc.Dial(o.GRPCAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
		if err != nil {
			return err
		}
		defer conn.Close()
	}
	res, err := ep.Run(context.Background(), conn)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func selectEP(cmd string, args []string) (runner, error) {
	switch cmd {
	case "points":
		return internal.Points(args)
	case "count":
		return internal.Count(args)
	case "compare":
		return internal.Compare(args)
	default:
		return nil, fmt.Errorf("command %q is not recognised", cmd)
	}
}
