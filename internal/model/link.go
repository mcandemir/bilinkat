package model

import "time"

type Link struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
