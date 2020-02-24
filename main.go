package main

import (
	"fmt"
	"jwtauth/api"
	"jwtauth/auth"
	"jwtauth/repository/psql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dipto"
	password = "12345"
	dbname   = "auth"
)

func main() {
	repo := chooseDb("postgres")
	service := auth.NewAuthService(repo)
	handler := api.NewHandler(service)
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/register",handler.RegisterPost)
	r.Post("/login",handler.LoginPost)
	errs := make(chan error, 2)

	go func() {
		fmt.Println("Listening on port 8000")
		errs <- http.ListenAndServe(":8000", r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}

func chooseDb(dbType string) auth.Repository{
	switch dbType {
	case "postgres":
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		repo, err := psql.NewPsqlRepository(psqlInfo)
		if err != nil{
			log.Fatal(err)
		}
		return repo

	}
	return nil
}
