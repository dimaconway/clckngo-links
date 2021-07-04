package postgresql

import (
	"clckngo/links/internal/domain"
	"database/sql"
	"fmt"
)

const linksTableName = "links"

type LinksRepository struct {
	DB *sql.DB
}

func (lr *LinksRepository) StoreLink(link *domain.Link) error {
	insertSql := fmt.Sprintf(`
INSERT INTO %s (created_datetime, url, code, hits, last_hit_datetime)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`,
		linksTableName,
	)

	var insertedLinkId string

	err := lr.DB.QueryRow(
		insertSql,
		link.CreatedDatetime,
		link.Url,
		link.Code,
		link.Hits,
		link.LastHitDatetime,
	).Scan(&insertedLinkId)
	if err != nil {
		return err
	}

	link.Id = &insertedLinkId

	return nil
}
