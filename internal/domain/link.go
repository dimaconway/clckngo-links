package domain

import "time"

type Link struct {
	Id              string     `json:"id"`
	CreatedDatetime time.Time  `json:"createdDatetime"`
	Url             string     `json:"url"`
	Code            string     `json:"code"`
	Hits            int        `json:"hits"`
	LastHitDatetime *time.Time `json:"lastHitDatetime"`
}
