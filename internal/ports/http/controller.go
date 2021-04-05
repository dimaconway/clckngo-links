package http

import (
	"clckngo/links/internal/application"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"net/http"
	"strings"
)

//go:embed json_schemas/store_link.json
var storeLinkValidationSchema string

type Controller struct{}

func (c Controller) GetLinks(w http.ResponseWriter, r *http.Request) {
	panic("implement me11")
}

func (c Controller) StoreLink(w http.ResponseWriter, r *http.Request) {
	postData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	schemaLoader := gojsonschema.NewStringLoader(storeLinkValidationSchema)
	validationResult, err := gojsonschema.Validate(schemaLoader, gojsonschema.NewBytesLoader(postData))
	if err != nil {
		panic(err)
	}
	if !validationResult.Valid() {
		var errorsForResponse []string
		for _, err := range validationResult.Errors() {
			errorsForResponse = append(errorsForResponse, err.String())
		}
		http.Error(w, strings.Join(errorsForResponse, "\n"), http.StatusBadRequest)
		return
	}
	var link Link
	err = json.Unmarshal(postData, &link)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", link)
	application.ShortURL(link.Url)
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
