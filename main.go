package main

import (
	"flag"

	"github.com/mcaci/briscola-serv/cli"
	"github.com/mcaci/briscola-serv/daemon"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
		isDaemon = flag.Bool("d", false, "start briscola-serv's daemon")
		isCli    = flag.Bool("cli", false, "start briscola-serv's cli")
	)
	flag.Parse()

	switch {
	case *isDaemon:
		daemon.Start(&daemon.Opts{
			HTTPAddr: *httpAddr,
			GRPCAddr: *gRPCAddr,
		})
	case *isCli:
		cli.Start(&cli.Opts{
			GRPCAddr: *gRPCAddr,
			Cmd:      flag.Args()[0],
			Args:     flag.Args()[1:],
		})
	}
}
