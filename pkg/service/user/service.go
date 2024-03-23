package user

import (
	"github.com/aashpv/db-microservice/pkg/db/user"
	"github.com/aashpv/db-microservice/pkg/models"
)

type Service interface {
	CreateUser(body string) (err error)
	UpdateUserById(body string) (err error)
	GetUserById(id int) (user models.User, err error)
	GetAllUsers() (users []models.User, err error)
	DeleteUserById(id int) (err error)
}

type service struct {
	pgs user.DataBase
}

func New(postgres user.DataBase) Service {
	return &service{pgs: postgres}
}
