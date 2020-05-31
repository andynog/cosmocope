package controller

import (
	"cosmocope/client"
)

func GetProjects() client.GithubSearchResult {
	topic := "cosmos-sdk"
	var searchResults client.GithubSearchResult
	searchResults = client.SearchGithub(topic)
	return searchResults
}
