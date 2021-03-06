package model

// A Cosmos SDK Module
type Module struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
	Url   string `json:"url"`
	Sha  string `json:"sha"`
}
