package product

import "FlowerHive/db-microservice/pkg/models"

func (s *service) GetProduct(id int) (product models.Product, err error) {
	product, err = s.pgs.GetProductById(id)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}

func (s *service) GetAllProduct() (products []models.Product, err error) {
	products, err = s.pgs.GetAllProducts()
	if err != nil { // if err use log for writing in file
		return
	}
	// если ошибок нет, то тоже надо логировать

	return
}
