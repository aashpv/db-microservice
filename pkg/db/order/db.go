package order

import (
	"FlowerHive/db-microservice/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataBase interface {
	Create(order models.Order) (err error)
	GetAllOrders() (orders []models.Order, err error)
	GetOrderById(id int) (order models.Order, err error)
	Update(order models.Order) (err error)
	Delete(idOrder int) (err error)
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
