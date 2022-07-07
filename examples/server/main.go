package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sreeram-narayanan/gosip"
	"github.com/sreeram-narayanan/gosip/log"
	"github.com/sreeram-narayanan/gosip/transport"
)

var (
	logger log.Logger
)

func init() {
	logger = log.NewDefaultLogrusLogger().WithPrefix("Server")
}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	srvConf := gosip.ServerConfig{}
	srv := gosip.NewServer(srvConf, nil, nil, logger)
	srv.Listen("wss", "0.0.0.0:5081", &transport.TLSConfig{Cert: "certs/cert.pem", Key: "certs/key.pem"})

	<-stop

	srv.Shutdown()
}
