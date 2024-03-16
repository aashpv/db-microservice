package user

import (
	"FlowerHive/db-microservice/pkg/models"
	"encoding/json"
)

func (s *service) Update(body string) (err error) {
	var user models.User

	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		return
	}

	err = s.pgs.Update(user)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
