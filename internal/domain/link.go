package domain

import "time"

type Link struct {
	Id              *string    `json:"id"`
	CreatedDatetime time.Time  `json:"createdDatetime"`
	Url             string     `json:"url"`
	Code            string     `json:"code"`
	Hits            int        `json:"hits"`
	LastHitDatetime *time.Time `json:"lastHitDatetime"`
}

func BuildLink(url string, code string) Link {
	return Link{
		Url:             url,
		Code:            code,
		CreatedDatetime: time.Now(),
	}
}
