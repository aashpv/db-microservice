package order

import (
	"encoding/json"
	"github.com/aashpv/db-microservice/pkg/models"
)

func (s *service) CreateOrder(body string) (err error) {
	var order models.Order

	err = json.Unmarshal([]byte(body), &order)
	if err != nil {
		return
	}

	err = s.pgs.CreateOrder(order)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
