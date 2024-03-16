package product

import (
	"FlowerHive/db-microservice/pkg/db/product"
	"FlowerHive/db-microservice/pkg/models"
)

type Service interface {
	Create(body string) (err error)
	Update(body string) (err error)
	GetProduct(id int) (product models.Product, err error)
	GetAllProduct() (products []models.Product, err error)
	Delete(id int) (err error)
}

type service struct {
	pgs product.DataBase
}

func New(postgres product.DataBase) Service {
	return &service{pgs: postgres}
}
