package controller

import (
	"fmt"
	"sort"
	"strings"

	"cosmocope/client"
	"cosmocope/model"
	"encoding/json"
	"github.com/alexeyco/simpletable"
	"github.com/dustin/go-humanize"
)

/// Logic to sort by Last Commit
type ByLastCommit []model.Project

func (a ByLastCommit) Len() int           { return len(a) }
func (a ByLastCommit) Less(i, j int) bool { return a[i].LastCommit.UnixNano() < a[j].LastCommit.UnixNano() }
func (a ByLastCommit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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
			License:     r.License.Key,
			Stars:       r.StargazersCount,
			Forks:       r.ForksCount,
			LastCommit: r.PushedAt,
		}

		// Logic to remove Azure CosmosDB listings
		if !strings.Contains(strings.ToLower(project.Description), "cosmosdb") {
			projects = append(projects, project)
		}
	}
	sort.Sort(sort.Reverse(ByLastCommit(projects)))
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
			{Align: simpletable.AlignCenter, Text: "LICENSE"},
			{Align: simpletable.AlignCenter, Text: "STARS"},
			{Align: simpletable.AlignCenter, Text: "FORKS"},
			{Align: simpletable.AlignCenter, Text: "LAST COMMIT"},
		},
	}
	count := 0

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
			{Text: p.License},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", p.Stars)},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", p.Forks)},
			{Text: humanize.Time(p.LastCommit)},
		}
		table.Body.Cells = append(table.Body.Cells, row)
		count++
	}
	fmt.Println("\r")

	// Table Footer
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: "Total:"},
			{Align: simpletable.AlignLeft, Span: len(table.Header.Cells) - 1, Text: fmt.Sprintf("%d", count)},
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