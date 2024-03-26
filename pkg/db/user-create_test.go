package db

import (
	"fmt"
	"github.com/aashpv/db-microservice/pkg/models"
	"testing"
)

func Test_postgres_CreateUser(t *testing.T) {
	user := models.User{
		ID:       0,
		Username: "",
		Password: "",
		Role:     "",
		Number:   "",
		Email:    "",
	}

	p, err2 := New("addawdada")
	if err2 != nil {
		fmt.Println("Connect error!", err2)
		return
	}

	err := p.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("user created!")
}
