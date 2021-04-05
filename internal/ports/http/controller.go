package http

import (
	"clckngo/links/internal/application"
	"encoding/json"
	"net/http"
)

type Controller struct{}

func (c Controller) StoreLink(w http.ResponseWriter, r *http.Request) {
	var requestBodyObject StoreLinkRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBodyObject)
	if err != nil {
		panic(err)
	}
	application.ShortURL(requestBodyObject.Url)
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
