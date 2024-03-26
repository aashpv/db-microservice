/*
 * @File: main.go
 * @Description: ...
 * @Author: ashpv (aashpvv@gmail.com)
 */

package main

import (
	"github.com/aashpv/db-microservice/pkg/db"
	"github.com/aashpv/db-microservice/pkg/server"
	"github.com/aashpv/db-microservice/pkg/server/routers"
	"github.com/aashpv/db-microservice/pkg/server/routers/handlers"
	"github.com/aashpv/db-microservice/pkg/service"
)

func main() {
	// Вообще по-хорошему все надо запускать с конфига файла, но это в конце будешь делать

	dataSourceName := "user=postgres password=1234 dbname=flowers host=185.233.82.174 port=5432 sslmode=disable"

	p, err := db.New(dataSourceName)
	if err != nil {
		return
	}
	// log
	src := service.New(p)
	// log
	hrs := handlers.New(src)
	// log
	rts := routers.New(hrs)
	// log
	srv := server.New("8080", rts)
	// log
	err = srv.Run()
	if err != nil {
		return
	}
}
