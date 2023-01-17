package main

import (
	"fmt"
	"net"

	sv "pricetracker/Products/internal/server"
	pb "pricetracker/pkg/build/pkg/proto"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

// FIXME JUST A LAYOUT
func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.WithError(err).Fatal("unable to start Products listener")
	}

	log.Info("creating a new grpc server for Products service")
	s := grpc.NewServer()
	server := sv.Start(ProductsDBs)
	pb.RegisterProductsServer(s, server)
	log.Info(fmt.Sprintf("server registered, listening on port: %d", port))

	if err := s.Serve(listener); err != nil {
		log.WithError(err).Fatal("grpc server fatal error")
	}
}

const port = 8083

var ProductsDBs = []string{
	"host=localhost port=5432 user=postgres password=$SECRET dbname=postgres sslmode=disable",
	"host=localhost port=5432 user=postgres password=$SECRET dbname=postgres sslmode=disable",
}
