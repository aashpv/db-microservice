package order

import (
	"FlowerHive/db-microservice/pkg/models"
	"encoding/json"
)

func (s *service) Create(body string) (err error) {
	var order models.Order

	err = json.Unmarshal([]byte(body), &order)
	if err != nil {
		return
	}

	err = s.pgs.Create(order)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
