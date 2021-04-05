package main

import (
	portsHttp "clckngo/links/internal/ports/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

func main() {
	var controller portsHttp.Controller

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	handler := portsHttp.HandlerFromMux(controller, router)

	port := os.Getenv("PORT")
	err := http.ListenAndServe(
		"localhost:"+port,
		handler,
	)
	if err != nil {
		return
	}
}
