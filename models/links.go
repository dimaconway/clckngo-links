// Package models provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package models

import (
	"time"
)

// Link defines model for Link.
type Link struct {
	Code            string    `json:"code"`
	CreatedDatetime time.Time `json:"createdDatetime"`
	Hits            int       `json:"hits"`
	Id              string    `json:"id"`
	LastHitDatetime time.Time `json:"lastHitDatetime"`
	Url             string    `json:"url"`
}
