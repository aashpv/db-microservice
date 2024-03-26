package service

import (
	"encoding/json"
	"github.com/aashpv/db-microservice/pkg/models"
)

func (s *service) UpdateUserById(body string) (err error) {
	var user models.User

	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		return
	}

	err = s.pgs.UpdateUserById(user)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
