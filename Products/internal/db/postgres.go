package db

import (
	"context"
	"database/sql"
	"time"

	pb "pricetracker/pkg/build/pkg/proto"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Postgres struct {
	db []*sql.DB
}

func Start(connStrings []string) Postgres {
	dbsCount := len(connStrings)
	connections := make([]*sql.DB, dbsCount)
	for i := 0; i < dbsCount; i++ {
		db, err := sql.Open("postgres", connStrings[i])
		if err != nil {
			log.WithField("conn", connStrings[i]).WithError(err).Fatal("unable to connect to db")
		}
		connections[i] = db
	}
	return Postgres{
		db: connections,
	}
}

func (p *Postgres) GetProductPrices(ctx context.Context, pr *pb.Product) (*pb.ProductPrices, error) {
	rows, err := p.db[getShopShard(pr.Shop)].
		Query(`SELECT update_ts, price FROM Products WHERE shop = $1 AND model = $2 ORDER BY update_ts`, pr.Shop, pr.Name)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"shop":    pr.Shop,
			"product": pr.Name,
		}).Error("unable to get product's prices")
		return nil, err
	}
	defer rows.Close()

	result := pb.ProductPrices{}
	for rows.Next() {
		var price pb.DatePrice
		if err = rows.Scan(&price.Ts, &price.Price); err != nil {
			log.WithError(err).Error("unable to scan price")
			return nil, err
		}
		result.Prices = append(result.Prices, &price)
	}
	log.WithField("product", *pr).Info("successfully scanned product's prices")
	return &result, nil
}

func (p *Postgres) AddNewPrice(ctx context.Context, newP *pb.ProductNewPrice) (*emptypb.Empty, error) {
	_, err := p.db[getShopShard(newP.Product.Shop)].ExecContext(ctx, `INSERT INTO Products VALUES ($1, $2, $3, $4, $5)`,
		newP.Product.Shop, newP.Product.Name, newP.Product.Url, newP.Price.Ts, newP.Price.Price)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"product": *newP.Product,
			"price":   *newP.Price,
		}).Error("failed to insert new price")
		return nil, err
	}
	log.WithFields(log.Fields{
		"product": *newP.Product,
	}).Info("new price for product inserted")
	return &emptypb.Empty{}, nil
}

func (p *Postgres) AddNewProduct(ctx context.Context, pr *pb.Product) (*emptypb.Empty, error) {
	_, err := p.db[getShopShard(pr.Shop)].ExecContext(ctx, `INSERT INTO Products VALUES ($1, $2, $3, $4, $5)`,
		pr.Shop, pr.Name, pr.Url, time.Now(), 0)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"product": *pr,
		}).Error("failed to insert new product")
		return nil, err
	}
	log.WithFields(log.Fields{
		"product": *pr,
	}).Info("new product inserted")
	return new(emptypb.Empty), nil
}

func (p *Postgres) GetAllProducts(ctx context.Context) (*pb.ProductList, error) {
	result := pb.ProductList{}
	for d := range p.db {
		rows, err := p.db[d].Query(`SELECT shop, name, url FROM products GROUP BY shop, name, url ORDER BY 1, 2, 3`)
		if err != nil {
			log.WithError(err).Error("unable to retrieve all products")
			return nil, err
		}
		defer rows.Close()

		res := pb.ProductList{}
		for rows.Next() {
			product := pb.Product{}
			if err = rows.Scan(&product.Shop, &product.Name, &product.Url); err != nil {
				log.WithError(err).Error("unable to scan some fields while retrieving all products")
				return nil, err
			}
			res.ProductsList = append(res.ProductsList, &product)
		}
		result.ProductsList = append(result.ProductsList, res.ProductsList...)
	}
	log.WithField("products retrieved", len(result.ProductsList)).Info("retrieved all products")
	return &result, nil
}

func getShopShard(shop string) int {
	if shop[0] < 'm' {
		return 1
	}
	return 0
}
