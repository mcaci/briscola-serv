package briscolad

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	endp "github.com/mcaci/briscola-serv/endpoint"
	"github.com/mcaci/briscola-serv/pb"
	briscolagrpc "github.com/mcaci/briscola-serv/server/grpc"
	briscolahttp "github.com/mcaci/briscola-serv/server/http"
	serv "github.com/mcaci/briscola-serv/service"
	"google.golang.org/grpc"
)

type Opts struct {
	HTTPAddr string
	GRPCAddr string
}

func Start(o *Opts) error {
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	srv := serv.NewService()
	data := srvData{
		ctx:       context.Background(),
		endpoints: makeEndpoints(srv),
		errChan:   errChan,
	}

	// start HTTP server
	log.Println("listenning to http requests on", o.HTTPAddr)
	data.addr = o.HTTPAddr
	go startHTTPSrv(data)

	// start gRPC server
	log.Println("listenning to grpc requests on", o.GRPCAddr)
	data.addr = o.GRPCAddr
	go startGRPCSrv(data)

	return <-errChan
}

type srvData struct {
	ctx       context.Context
	addr      string
	endpoints endp.Endpoints
	errChan   chan<- error
}

func startHTTPSrv(srv srvData) {
	handler := briscolahttp.NewHTTPServer(srv.ctx, srv.endpoints)
	srv.errChan <- http.ListenAndServe(srv.addr, handler)
}

func startGRPCSrv(srv srvData) {
	listener, err := net.Listen("tcp", srv.addr)
	if err != nil {
		srv.errChan <- err
		return
	}
	handler := briscolagrpc.NewGRPCServer(srv.ctx, srv.endpoints)
	gRPCServer := grpc.NewServer()
	pb.RegisterBriscolaServer(gRPCServer, handler)
	srv.errChan <- gRPCServer.Serve(listener)
}

func makeEndpoints(srv serv.Service) endp.Endpoints {
	pointsEndpoint := endp.MakePointsEndpoint(srv)
	countEndpoint := endp.MakeCountEndpoint(srv)
	compareEndpoint := endp.MakeCompareEndpoint(srv)
	return endp.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
