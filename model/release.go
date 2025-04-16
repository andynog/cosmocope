package model

import "time"

// Release represents a GitHub release
type Release struct {
	Name        string    `json:"name"`
	TagName     string    `json:"tag_name"`
	URL         string    `json:"url"`
	Draft       bool      `json:"draft"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published"`
}
