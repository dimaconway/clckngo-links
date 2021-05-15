package application

import (
	"clckngo/links/internal/domain"
	"clckngo/links/internal/ports/db/postgresql"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
)

type LinksService struct {
	LinksRepository *postgresql.LinksRepository
}

func (ls *LinksService) ShortURL(url string) (*domain.Link, error) {
	if !govalidator.IsRequestURL(url) {
		return nil, errors.New(fmt.Sprintf("URL %s is not valid", url))
	}

	code := hashUrl(url)
	link := domain.BuildLink(url, code)
	err := ls.LinksRepository.StoreLink(&link)

	return &link, err
}

func hashUrl(url string) string {
	hash := md5.Sum([]byte(url))
	return fmt.Sprintf("%x", hash)[:5]
}
