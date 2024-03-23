package server

import (
	"github.com/aashpv/db-microservice/pkg/db/product"
	"github.com/aashpv/db-microservice/pkg/server/handlers"
	product2 "github.com/aashpv/db-microservice/pkg/service/product"
	"log"
)

func InitServer() {
	// Создаем экземпляр базы данных
	dataSourceName := "postgres://postgres:postgres@localhost:5432/flowers?sslmode=disable"
	db, err := product.New(dataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Создаем экземпляр службы
	_ = product2.New(db)
	if err != nil {
		log.Fatal("Failed to connect to the service: ", err)
	}

	// Передаем экземпляр службы в обработчики и настраиваем маршруты
	server := handlers.SetupRouter()

	// Запускаем сервер
	if err := server.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
