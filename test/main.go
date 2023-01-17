package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "pricetracker/pkg/build/pkg/proto"

	log "github.com/sirupsen/logrus"
)

func main() {
	sdconn, err := grpc.Dial(users, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithError(err).Fatal("unable to connect to service dispatcher")
		// FIXME add retry policy
	}
	defer sdconn.Close()
	user := pb.NewUsersClient(sdconn)
	list, err := user.AddProduct(context.Background(), &pb.UserProduct{
		User: &pb.User{
			Name: "marcin",
		},
		Product: &pb.Product{
			Shop: "euro",
			Name: "laptop",
			Url:  "https://www.euro.com.pl/laptopy-i-netbooki/asus-laptop-fx506-i5-16gb-512ssd-3060-w11.bhtml",
		},
	})
	log.WithError(err).WithField("list", list).Info("inserted new product")
}

const users = "localhost:8081"
