package service

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (s *service) UpdateOrderById(order models.Order) (err error) {
	err = s.pgs.UpdateOrderById(order)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
