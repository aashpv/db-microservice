package product

import (
	"github.com/aashpv/db-microservice/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataBase interface {
	CreateProduct(product models.Product) (err error)
	GetAllProducts() (products []models.Product, err error)
	GetProductById(id int) (product models.Product, err error)
	UpdateProductById(product models.Product) (err error)
	DeleteProductById(idProduct int) (err error)
}

type postgres struct {
	db *sqlx.DB
}

func New(dataSourceName string) (p DataBase, err error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &postgres{db: db}, nil
}
