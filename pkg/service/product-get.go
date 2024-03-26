package service

import "github.com/aashpv/db-microservice/pkg/models"

func (s *service) GetProductById(id int) (product models.Product, err error) {
	product, err = s.pgs.GetProductById(id)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}

func (s *service) GetAllProducts() (products []models.Product, err error) {
	products, err = s.pgs.GetAllProducts()
	if err != nil { // if err use log for writing in file
		return
	}
	// если ошибок нет, то тоже надо логировать

	return
}
