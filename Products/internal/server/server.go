package server

import (
	"context"

	psql "pricetracker/Products/internal/db"
	pb "pricetracker/pkg/build/pkg/proto"

	log "github.com/sirupsen/logrus"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type ProductsServer struct {
	pb.UnimplementedProductsServer
	psql psql.Postgres
}

func Start(connStrings []string) pb.ProductsServer {
	p := psql.Start(connStrings)
	return &ProductsServer{psql: p}
}

func (s *ProductsServer) GetProductPrices(ctx context.Context, p *pb.Product) (*pb.ProductPrices, error) {
	result, err := s.psql.GetProductPrices(ctx, p)
	if err != nil {
		log.WithError(err).WithField("product", *p).Error("unable to get product's prices")
		return nil, err
	}
	log.WithField("product", *p).Info("product's prices queried successfully")
	return result, nil
}

func (s *ProductsServer) AddNewPrice(ctx context.Context, newP *pb.ProductNewPrice) (*emptypb.Empty, error) {
	_, err := s.psql.AddNewPrice(ctx, newP)
	if err != nil {
		log.WithError(err).Error("unable to add new price")
	}
	return nil, err
}

func (s *ProductsServer) AddNewProduct(ctx context.Context, p *pb.Product) (*emptypb.Empty, error) {
	_, err := s.psql.AddNewProduct(ctx, p)
	if err != nil {
		log.WithError(err).Error("unable to add new produt")
	}
	return nil, err
}
