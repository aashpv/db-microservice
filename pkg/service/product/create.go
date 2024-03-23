package product

import (
	"encoding/json"
	"github.com/aashpv/db-microservice/pkg/models"
)

func (s *service) CreateProduct(body string) (err error) {
	var product models.Product

	err = json.Unmarshal([]byte(body), &product)
	if err != nil {
		return
	}

	err = s.pgs.CreateProduct(product)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
