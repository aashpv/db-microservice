package server

import (
	"context"
	"github.com/aashpv/db-microservice/pkg/server/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Run() error
}

type server struct {
	addr    string
	routers routers.Routers
}

func New(addr string, routers routers.Routers) Server {
	return &server{
		addr:    addr,
		routers: routers,
	}
}

func (s *server) Run() error {
	srv := http.Server{
		Addr:         ":" + s.addr,
		Handler:      s.routers.NewRouter(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("error start server: %v", err)
		}
	}()

	// Ожидание сигнала SIGINT (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT)
	<-stop

	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error shutdown server: %v", err)
		return err
	}

	log.Println("successful shutdown server")
	return nil
}
