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
	err := cli.Start(&cli.Opts{GRPCAddr: ":8081", EpRun: mockRunner{}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartKO(t *testing.T) {
	err := cli.Start(&cli.Opts{GRPCAddr: ":8081", EpRun: mockRunner{err: errors.New("Start KO")}})
	if err == nil {
		t.Fatal(err)
	}
}
