package http

import (
	"net/http"
)

type Server struct{}

func (s Server) GetLinks(w http.ResponseWriter, r *http.Request) {
	panic("implement me11")
}

func (s Server) StoreLink(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s Server) DeleteLink(w http.ResponseWriter, r *http.Request, code string) {
	panic("implement me")
}

func (s Server) GetLink(w http.ResponseWriter, r *http.Request, code string) {
	panic("implement me")
}

func (s Server) IncrementHitsCounter(w http.ResponseWriter, r *http.Request, code string) {
	panic("implement me")
}

func (s Server) RedirectByLink(w http.ResponseWriter, r *http.Request, code string) {
	panic("implement me")
}
