package http

import (
	"clckngo/links/internal/application"
	"encoding/json"
	"net/http"
)

func respondJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		panic(err)
	}
}

type Controller struct{}

func (c Controller) StoreLink(w http.ResponseWriter, r *http.Request) {
	var requestBodyObject StoreLinkRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBodyObject)
	if err != nil {
		panic(err)
	}
	link, err := application.ShortURL(requestBodyObject.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondJSON(w, link)
}

type StoreLinkRequestBody struct {
	Url string
}

func (c Controller) GetLinks(w http.ResponseWriter, r *http.Request) {
	panic("implement me11")
}

func (c Controller) DeleteLink(w http.ResponseWriter, r *http.Request, code string) {
	panic("implement me")
}

func (c Controller) GetLink(w http.ResponseWriter, r *http.Request, code string) {
	panic("implement me")
}

func (c Controller) IncrementHitsCounter(w http.ResponseWriter, r *http.Request, code string) {
	panic("implement me")
}

func (c Controller) RedirectByLink(w http.ResponseWriter, r *http.Request, code string) {
	panic("implement me")
}
