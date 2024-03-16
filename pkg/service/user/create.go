package user

import (
	"FlowerHive/db-microservice/pkg/models"
	"encoding/json"
)

func (s *service) Create(body string) (err error) {
	var user models.User

	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		return
	}

	err = s.pgs.AddNewUser(user)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
