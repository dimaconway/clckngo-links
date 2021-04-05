package application

import (
	"clckngo/links/internal/domain"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
)

func ShortURL(url string) (*domain.Link, error) {
	println(url)
	if !govalidator.IsRequestURI(url) {
		return nil, errors.New(fmt.Sprintf("URL %s is not valid", url))
	}

	return &domain.Link{}, nil
}
