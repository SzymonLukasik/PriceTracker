package main

import (
	"fmt"
	"net"

	sv "pricetracker/Users/internal/server"
	pb "pricetracker/pkg/build/pkg/proto"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.WithError(err).Fatal("unable to start Users listener")
	}

	log.Info("creating a new grpc server for Users service")
	s := grpc.NewServer()
	server := sv.Start(UsersDB, Products)
	pb.RegisterUsersServer(s, server)
	log.Info(fmt.Sprintf("server registered, listening on port: %d", port))

	if err := s.Serve(listener); err != nil {
		log.WithError(err).Fatal("grpc server fatal error")
	}
}

const port = 8081
const UsersDB = "host=localhost port=5432 user=postgres password=$SECRET dbname=postgres sslmode=disable"
const Products = "localhost:8083"
