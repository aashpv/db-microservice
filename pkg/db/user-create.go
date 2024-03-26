package db

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (p *postgres) CreateUser(user models.User) (err error) {
	_, err = p.db.Exec("INSERT INTO users (username, password, role, number, email) VALUES ($1, $2, $3, $4, $5)",
		user.Username,
		user.Password,
		user.Role,
		user.Number,
		user.Email)

	if err != nil {
		return
	}

	return
}
