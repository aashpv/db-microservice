package user

import (
	"github.com/aashpv/db-microservice/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataBase interface {
	CreateUser(user models.User) (err error)
	GetAllUsers() (users []models.User, err error)
	GetUserById(id int) (user models.User, err error)
	UpdateUserById(user models.User) (err error)
	DeleteUserById(idUser int) (err error)
}

type postgres struct {
	db *sqlx.DB
}

func New(dataSourceName string) (p DataBase, err error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &postgres{db: db}, nil
}
