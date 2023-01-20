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
	return new(emptypb.Empty), err
}

func (s *ProductsServer) AddNewProduct(ctx context.Context, p *pb.Product) (*emptypb.Empty, error) {
	_, err := s.psql.AddNewProduct(ctx, p)
	if err != nil {
		log.WithError(err).Error("unable to add new produt")
		return nil, err
	}
	log.Info("new product added")
	return new(emptypb.Empty), nil
}

func (s *ProductsServer) GetAllProducts(ctx context.Context, e *emptypb.Empty) (*pb.ProductList, error) {
	result, err := s.psql.GetAllProducts(ctx)
	if err != nil {
		log.WithError(err).Error("unable to get all products")
		return nil, err
	}
	log.Info("all products retrieved successfully")
	return result, nil
}
