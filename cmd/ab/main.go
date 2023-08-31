package main

import (
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/dep/postgres"
	"net/http"

	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/contoller/server"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := postgres.NewPostgres("pgdb", 5432, "postgres", "postgres", "avito")
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(db)
	l := logrus.New()
	s := &http.Server{Addr: ":8000"}
	c := server.NewServer(s, h, l)

	err = c.Serve()
	if err != nil {
		panic(err)
	}
}
