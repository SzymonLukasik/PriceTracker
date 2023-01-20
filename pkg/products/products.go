package products

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "pricetracker/pkg/build/pkg/proto"

	log "github.com/sirupsen/logrus"
)

func NewProductsClient(addrPort string) pb.ProductsClient {
	conn, err := grpc.Dial(addrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithError(err).Fatal("unable to connect to products")
	}
	return pb.NewProductsClient(conn)
}
