package user

import (
	"FlowerHive/db-microservice/pkg/db/user"
	"FlowerHive/db-microservice/pkg/models"
)

type Service interface {
	Create(body string) (err error)
	Update(body string) (err error)
	GetUser(id int) (user models.User, err error)
	GetAllUser() (users []models.User, err error)
	Delete(id int) (err error)
}

type service struct {
	pgs user.DataBase
}

func New(postgres user.DataBase) Service {
	return &service{pgs: postgres}
}
