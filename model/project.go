package model

import "time"

// Project
type Project struct {
	Name        string    `json:"name"`
	Owner       string    `json:"owner"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	License     string    `json:"license"`
	Stars       int       `json:"stars"`
	Forks       int       `json:"forks"`
	LastCommit  time.Time `json:"lastcommit"`
	Modules     []Module  `json:"modules,omitempty"`
	Branch      string    `json:"branch"`
	CosmosSDK   string     `json:"cosmos_sdk,omitempty"`
}
