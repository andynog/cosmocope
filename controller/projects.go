package controller

import (
	"cosmocope/client"
	"cosmocope/model"
)

func GetProjects() []model.Project {
	topic := "cosmos-sdk"
	var projects []model.Project
	var searchResults client.GithubSearchResult
	searchResults = client.SearchGithub(topic)
	for _, r := range searchResults.Items {
		project := model.Project{
			Name:        r.Name,
			Owner:       r.Owner.Login,
			Url:         r.HTMLURL,
			Description: r.Description,
			Language:    r.Language,
			Stars:       r.StargazersCount,
			Forks:       r.ForksCount,
			LastUpdated: r.UpdatedAt,
		}
		projects = append(projects, project)
	}
	return projects
}
