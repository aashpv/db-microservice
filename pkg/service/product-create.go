package service

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (s *service) CreateProduct(product models.Product) (err error) {
	err = s.pgs.CreateProduct(product)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
