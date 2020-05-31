package controller

import (
	"fmt"
	"strings"

	"cosmocope/client"
	"cosmocope/model"
	"encoding/json"
	"github.com/alexeyco/simpletable"
	"github.com/dustin/go-humanize"
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

		// Logic to remove Azure CosmosDB listings
		if !strings.Contains(strings.ToLower(project.Description), "cosmosdb") {
			projects = append(projects, project)
		}
	}
	return projects
}

// Print Projects in Table format
func PrintProjectsTable(projects []model.Project) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "OWNER"},
			{Align: simpletable.AlignCenter, Text: "URL"},
			{Align: simpletable.AlignCenter, Text: "DESCRIPTION"},
			{Align: simpletable.AlignCenter, Text: "LANGUAGE"},
			{Align: simpletable.AlignCenter, Text: "STARS"},
			{Align: simpletable.AlignCenter, Text: "FORKS"},
			{Align: simpletable.AlignCenter, Text: "UPDATED"},
		},
	}
	countCosmosProjects := 0

	for _, p := range projects {
		var description string
		if len(p.Description) > 44 {
			description = p.Description[0:42] + "..."
		} else {
			description = p.Description
		}
		row := []*simpletable.Cell{
			{Text: p.Name},
			{Text: p.Owner},
			{Text: p.Url},
			{Text: strings.ToValidUTF8(description, "")},
			{Text: p.Language},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", p.Stars)},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", p.Forks)},
			{Text: humanize.Time(p.LastUpdated)},
		}
		table.Body.Cells = append(table.Body.Cells, row)
		countCosmosProjects++
	}
	fmt.Println("\r")

	// Table Footer
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: "Total Projects"},
			{Align: simpletable.AlignLeft, Span: len(table.Header.Cells) - 1, Text: fmt.Sprintf("%d", countCosmosProjects)},
		},
	}

	// Print table
	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

// Print Projects in JSON format
func PrintProjectsJSON(projects []model.Project) {
	json, err := json.MarshalIndent(projects, "", " ")
	if err != nil {
		fmt.Println("Failed to print JSON results:", err)
		return
	}
	fmt.Println(string(json))
}