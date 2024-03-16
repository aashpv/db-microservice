package user

import "FlowerHive/db-microservice/pkg/models"

func (s *service) GetUser(id int) (user models.User, err error) {
	user, err = s.pgs.GetUserById(id)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}

func (s *service) GetAllUser() (users []models.User, err error) {
	users, err = s.pgs.GetAllUsers()
	if err != nil { // if err use log for writing in file
		return
	}
	// если ошибок нет, то тоже надо логировать

	return
}
