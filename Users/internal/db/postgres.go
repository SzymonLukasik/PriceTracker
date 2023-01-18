package db

import (
	"context"
	"database/sql"

	pb "pricetracker/pkg/build/pkg/proto"
	dbtx "pricetracker/pkg/db"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Postgres struct {
	db *sql.DB
}

func Start(connString string) Postgres {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.WithField("conn", connString).WithError(err).Fatal("unable to connect to db")
	}
	return Postgres{
		db: db,
	}
}

func (p *Postgres) GetProducts(user *pb.User) (*pb.ProductList, error) {
	rows, err := p.db.Query(`SELECT shop, model, url FROM Users WHERE username=$1`, user.Name)
	if err != nil {
		log.WithError(err).WithField("user", user).Error("unable to retrieve products")
		return nil, err
	}
	defer rows.Close()

	result := pb.ProductList{}
	for rows.Next() {
		var product pb.Product
		if err = rows.Scan(&product.Shop, &product.Name, &product.Url); err != nil {
			log.WithError(err).Error("unable to scan product")
			return nil, err
		}
		result.ProductsList = append(result.ProductsList, &product)
	}
	log.WithFields(log.Fields{
		"user":            user.Name,
		"prods retrieved": len(result.ProductsList),
	}).Info("successfully retrieved user's products")
	return &result, nil
}

func (p *Postgres) AddProduct(ctx context.Context, up *pb.UserProduct) error {
	f := func() error {
		_, err := p.db.ExecContext(ctx, `INSERT INTO Users VALUES ($1, $2, $3, $4)`,
			up.User.Name, up.Product.Shop, up.Product.Name, up.Product.Url)
		if err != nil {
			log.WithError(err).Error("unable to add new product")
			return err
		}
		log.WithFields(log.Fields{
			"user":    up.User.Name,
			"product": up.Product,
		}).Info("products successfully added")
		return nil
	}
	return dbtx.ExecuteTransaction(ctx, p.db, f)
}
