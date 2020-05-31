package model

import "time"

// Project
type Project struct {
	Name        string    `json:"name"`
	Owner       string    `json:"owner"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Stars       int       `json:"stars"`
	Forks       int       `json:"forks"`
	LastUpdated time.Time `json:"updated"`
}
