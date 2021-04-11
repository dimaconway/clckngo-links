package application

import (
	"clckngo/links/internal/domain"
	"clckngo/links/internal/ports"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
)

func ShortURL(url string) (*domain.Link, error) {
	if !govalidator.IsRequestURL(url) {
		return nil, errors.New(fmt.Sprintf("URL %s is not valid", url))
	}

	code := ports.HashUrl(url)

	link := domain.BuildLink(url, code)

	return &link, nil
}
