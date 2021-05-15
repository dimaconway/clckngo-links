package main

import (
	"clckngo/links/internal/application"
	"clckngo/links/internal/ports/db/postgresql"
	portsHttp "clckngo/links/internal/ports/http"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/http"
	"os"
)

func main() {
	db, err := initDb()
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	linksRepository := postgresql.LinksRepository{DB: db}
	linksService := application.LinksService{LinksRepository: &linksRepository}
	controller := portsHttp.Controller{LinksService: &linksService}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	handler := portsHttp.HandlerFromMux(controller, router)

	port := os.Getenv("LINKS_PORT")
	if port == "" {
		panic("Empty LINKS_PORT in env")
	}

	fmt.Printf("Starting server on a port %s ...\n", port)
	err = http.ListenAndServe(
		"0.0.0.0:"+port,
		handler,
	)
	if err != nil {
		return
	}
}

func initDb() (*sql.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHostname := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	return sql.Open(
		"pgx",
		fmt.Sprintf(
			"%s:%s@%s:%s/%s",
			dbUser,
			dbPassword,
			dbHostname,
			dbPort,
			dbName,
		),
	)
}
