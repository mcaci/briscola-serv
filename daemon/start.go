package daemon

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	briscolagrpc "github.com/mcaci/briscola-serv/daemon/grpc"
	briscolahttp "github.com/mcaci/briscola-serv/daemon/http"
	httpmdw "github.com/mcaci/briscola-serv/daemon/http/mdw"
	"github.com/mcaci/briscola-serv/pb"
	"google.golang.org/grpc"
)

type Opts struct {
	HTTPAddr string
	GRPCAddr string
}

func Start(o *Opts) error {
	select {
	case err := <-startHTTP(context.Background(), o.HTTPAddr):
		return err
	case err := <-startGRPC(context.Background(), o.GRPCAddr):
		return err
	case err := <-handleSigTerm():
		return err
	}
}

func startHTTP(ctx context.Context, addr string) <-chan error {
	log.Println("listenning to http requests on", addr)
	errChan := make(chan error)
	go func() {
		<-ctx.Done()
		errChan <- ctx.Err()
	}()
	go func() {
		handler := briscolahttp.NewHandler(ctx, httpmdw.Logged)
		errChan <- http.ListenAndServe(addr, handler)
	}()
	return errChan
}

func startGRPC(ctx context.Context, addr string) <-chan error {
	log.Println("listenning to grpc requests on", addr)
	errChan := make(chan error)
	go func() {
		<-ctx.Done()
		errChan <- ctx.Err()
	}()
	go func() {
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			errChan <- err
			return
		}
		defer listener.Close()
		server := briscolagrpc.NewServer(ctx)
		gRPCServer := grpc.NewServer()
		pb.RegisterBriscolaServer(gRPCServer, server)
		errChan <- gRPCServer.Serve(listener)
	}()
	return errChan
}

func handleSigTerm() <-chan error {
	log.Println("Press Ctrl+C to terminate")
	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	return errChan
}
