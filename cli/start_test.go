package cli_test

import (
	"context"
	"errors"
	"testing"

	"github.com/mcaci/briscola-serv/cli"
	"google.golang.org/grpc"
)

type mockRunner struct{ err error }

func (mockRunner) SetEndpoint(*grpc.ClientConn)       {}
func (r mockRunner) Run(context.Context) (any, error) { return 0, r.err }

func TestStartOK(t *testing.T) {
	err := cli.Start(&cli.Opts{EpRun: mockRunner{}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartKO(t *testing.T) {
	testcases := []struct {
		name string
		o    *cli.Opts
	}{
		{name: "mock runner with error", o: &cli.Opts{EpRun: mockRunner{err: errors.New("ko at start")}}},
		{name: "no address provided", o: &cli.Opts{}},
		{name: "no command provided", o: &cli.Opts{GRPCAddr: ":8081"}},
		{name: "mock runner with error", o: &cli.Opts{GRPCAddr: ":8081", Cmd: "points"}},
		{name: "mock runner with error", o: &cli.Opts{GRPCAddr: ":8081", Cmd: "count"}},
		{name: "mock runner with error", o: &cli.Opts{GRPCAddr: ":8081", Cmd: "compare"}},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := cli.Start(tc.o)
			if err == nil {
				t.Fatal(err)
			}
		})
	}
}
