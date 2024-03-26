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

	p, err := db.New("addr")
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
