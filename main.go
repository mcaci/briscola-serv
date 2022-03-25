package main

import (
	"flag"

	"github.com/mcaci/briscola-serv/cmd/briscolad"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)
	flag.Parse()

	briscolad.Start(&briscolad.Opts{
		HTTPAddr: *httpAddr,
		GRPCAddr: *gRPCAddr,
	})
}
