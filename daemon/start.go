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
	briscola "github.com/mcaci/briscola-serv/daemon/lib"
	"github.com/mcaci/briscola-serv/pb"
	"google.golang.org/grpc"
)

type Opts struct {
	HTTPAddr string
	GRPCAddr string
}

func Start(o *Opts) error {
	select {
	case err := <-startHTTPSrv(newSrvData(o.HTTPAddr)):
		return err
	case err := <-startGRPCSrv(newSrvData(o.GRPCAddr)):
		return err
	case err := <-handleSigTerm():
		return err
	}
}

type srvData struct {
	ctx       context.Context
	addr      string
	endpoints Endpoints
}

func newSrvData(addr string) *srvData {
	srv := briscola.NewService()
	data := srvData{
		ctx:       context.Background(),
		addr:      addr,
		endpoints: newServerEndpoints(srv),
	}
	return &data
}

func startHTTPSrv(srv *srvData) <-chan error {
	log.Println("listenning to http requests on", srv.addr)
	handler := briscolahttp.NewHTTPServer(srv.ctx, srv.endpoints, srv.endpoints)
	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(srv.addr, handler)
	}()
	return errChan
}

func startGRPCSrv(srv *srvData) <-chan error {
	log.Println("listenning to grpc requests on", srv.addr)
	errChan := make(chan error)
	go func() {
		listener, err := net.Listen("tcp", srv.addr)
		if err != nil {
			errChan <- err
			return
		}
		handler := briscolagrpc.NewGRPCServer(srv.ctx, srv.endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterBriscolaServer(gRPCServer, handler)
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

type services interface {
	CardPoints(ctx context.Context, number uint32) (uint32, error)
	PointCount(ctx context.Context, number []uint32) (uint32, error)
	CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error)
}

func newServerEndpoints(srv services) Endpoints {
	pointsEndpoint := newPointsEndpoint(srv)
	countEndpoint := newCountEndpoint(srv)
	compareEndpoint := newCompareEndpoint(srv)
	return Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
