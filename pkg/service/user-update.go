package service

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (s *service) UpdateUserById(user models.User) (err error) {
	err = s.pgs.UpdateUserById(user)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
