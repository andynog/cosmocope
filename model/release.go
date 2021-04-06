package model

import "time"

// Release
type Release struct {
	Name        string    `json:"name"`
	TagName     string    `json:"tag_name"`
	Url         string    `json:"url"`
	Draft       bool      `json:"draft"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published"`
}
