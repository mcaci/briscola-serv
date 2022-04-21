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
	EpRun    endpointRunner
}

func Start(o *Opts) error {
	ep := o.EpRun
	if ep == nil {
		var conn *grpc.ClientConn
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
		ep.SetEndpoint(conn)
	}
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

func selectEP(cmd string, args []string) (endpointRunner, error) {
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
