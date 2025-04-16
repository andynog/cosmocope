// Package model provides the data structures for the application
package model

// Module represents a Cosmos SDK Module
type Module struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
	URL   string `json:"url"`
	Sha   string `json:"sha"`
}
