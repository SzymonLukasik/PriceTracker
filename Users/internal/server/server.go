package server

import (
	"context"

	psql "pricetracker/Users/internal/db"
	pb "pricetracker/pkg/build/pkg/proto"

	prod "pricetracker/pkg/products"

	log "github.com/sirupsen/logrus"
)

type UserServer struct {
	pb.UnimplementedUsersServer
	psql psql.Postgres
	prod pb.ProductsClient
}

type ProductsClient struct {
}

func Start(connString string, prodAddrPort string) pb.UsersServer {
	p := psql.Start(connString)
	return &UserServer{
		psql: p,
		prod: prod.NewProductsClient(prodAddrPort),
	}
}

func (s *UserServer) GetProducts(ctx context.Context, user *pb.User) (*pb.ProductList, error) {
	products, err := s.psql.GetProducts(user)
	if err != nil {
		log.WithField("user", user.Name).WithError(err).Error("unable to get products for user")
		return nil, err
	}
	return products, nil
}

func (s *UserServer) AddProduct(ctx context.Context, up *pb.UserProduct) (*pb.ProductList, error) {
	if err := s.psql.AddProduct(ctx, up); err != nil {
		log.WithError(err).WithField("user-product", *up).Error("unable to add product to user's pool")
		return nil, err
	}
	list, err := s.psql.GetProducts(up.User)
	if err != nil {
		log.WithField("user-product", *up).WithError(err).Error("unable to get user products after adding a new item")
		return nil, err
	}
	_, err = s.prod.AddNewProduct(ctx, up.Product)
	if err != nil {
		log.WithField("product", *up.Product).WithError(err).Error("unable to add product to global pool")
		// FIXME this results in lack of consistency between tables!
		return nil, err
	}
	log.WithField("user-product", *up).Info("successfully added a product to user's pool")
	return list, nil
}
