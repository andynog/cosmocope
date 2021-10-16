package model

import "time"

// Release
type Release struct {
	Name        string     `json:"name"`
	TagName     string     `json:"tag_name"`
	Url         string     `json:"url"`
	Draft       bool       `json:"draft"`
	PreRelease  bool       `json:"pre_release"`
	Description string     `json:"description"`
	PublishedAt time.Time  `json:"published"`
	Dependency  Dependency `json:"dependencies"`
}
