package user

import (
	"github.com/aashpv/db-microservice/pkg/models"
)

func (p *postgres) UpdateUserById(user models.User) (err error) {
	_, err = p.db.Exec(
		"UPDATE users SET username = $1, password = $2, number = $3, email = $4 WHERE id = $5",
		user.Username,
		user.Password,
		user.Number,
		user.Email,
		user.ID)

	if err != nil {
		return
	}

	return
}
