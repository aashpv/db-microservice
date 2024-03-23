package product

import (
	"github.com/aashpv/db-microservice/pkg/db/product"
	"github.com/aashpv/db-microservice/pkg/models"
)

type Service interface {
	CreateProduct(body string) (err error)
	UpdateProductById(body string) (err error)
	GetProductById(id int) (product models.Product, err error)
	GetAllProducts() (products []models.Product, err error)
	DeleteProductById(id int) (err error)
}

type service struct {
	pgs product.DataBase
}

func New(postgres product.DataBase) Service {
	return &service{pgs: postgres}
}
