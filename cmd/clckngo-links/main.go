package main

import (
	portsHttp "clckngo/links/internal/ports/http"
	"fmt"
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

	port := os.Getenv("LINKS_PORT")

	fmt.Printf("Starting server on a port %s ...\n", port)
	err := http.ListenAndServe(
		"0.0.0.0:"+port,
		handler,
	)
	if err != nil {
		return
	}
}
