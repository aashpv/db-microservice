package service

import "github.com/aashpv/db-microservice/pkg/models"

func (s *service) GetOrderById(id int) (order models.Order, err error) {
	order, err = s.pgs.GetOrderById(id)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}

func (s *service) GetAllOrders() (orders []models.Order, err error) {
	orders, err = s.pgs.GetAllOrders()
	if err != nil { // if err use log for writing in file
		return
	}
	// если ошибок нет, то тоже надо логировать

	return
}
