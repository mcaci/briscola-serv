package daemon

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	briscolagrpc "github.com/mcaci/briscola-serv/daemon/grpc"
	briscolahttp "github.com/mcaci/briscola-serv/daemon/http"
	httpmdw "github.com/mcaci/briscola-serv/daemon/http/mdw"
)

type Opts struct {
	Ctx      context.Context
	HTTPAddr string
	GRPCAddr string
}

func NewServer(o *Opts) error {
	if o.Ctx == nil {
		o.Ctx = context.Background()
	}
	select {
	case err := <-briscolahttp.NewServer(httpmdw.Logged).Start(o.Ctx, o.HTTPAddr):
		return err
	case err := <-briscolagrpc.NewServer().Start(o.Ctx, o.GRPCAddr):
		return err
	case err := <-handleSigTerm(o.Ctx):
		return err
	case <-o.Ctx.Done():
		return nil
	}
}

func handleSigTerm(ctx context.Context) <-chan error {
	log.Println("Press Ctrl+C to terminate")
	errChan := make(chan error)
	go func() {
		<-ctx.Done()
		errChan <- ctx.Err()
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	return errChan
}
