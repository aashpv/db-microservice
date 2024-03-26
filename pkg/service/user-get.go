package service

import "github.com/aashpv/db-microservice/pkg/models"

func (s *service) GetUserById(id int) (user models.User, err error) {
	user, err = s.pgs.GetUserById(id)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}

func (s *service) GetAllUsers() (users []models.User, err error) {
	users, err = s.pgs.GetAllUsers()
	if err != nil { // if err use log for writing in file
		return
	}
	// если ошибок нет, то тоже надо логировать

	return
}
