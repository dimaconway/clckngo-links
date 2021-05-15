package postgresql

import (
	"clckngo/links/internal/domain"
	"database/sql"
)

type LinksRepository struct{
	DB *sql.DB
}

func (lr *LinksRepository) StoreLink(link *domain.Link) error {
	return nil
}
