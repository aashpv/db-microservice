package order

import (
	"FlowerHive/db-microservice/pkg/models"
	"encoding/json"
)

func (s *service) Update(body string) (err error) {
	var order models.Order

	err = json.Unmarshal([]byte(body), &order)
	if err != nil {
		return
	}

	err = s.pgs.Update(order)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
