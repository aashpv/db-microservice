package user

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (p *postgres) GetAllUsers() (users []models.User, err error) {
	err = p.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return
}

func (p *postgres) GetUserById(id int) (user models.User, err error) {
	err = p.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return
	}

	return
}
