package service

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (s *service) CreateOrder(order models.Order) (err error) {
	err = s.pgs.CreateOrder(order)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
