package controller

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"sort"
	"strings"

	"encoding/json"
	"github.com/alexeyco/simpletable"
	"github.com/andynog/cosmocope/v2/client"
	"github.com/andynog/cosmocope/v2/model"
	"github.com/dustin/go-humanize"
)

func GetProjects(sortBy string) (result []model.Project, err error) {
	topic := "cosmos-sdk"
	var projects []model.Project
	searchResults, err := client.SearchGithub(topic)
	if err != nil {
		return nil, fmt.Errorf("problems fetching projects")
	}

	// Progress bar
	bar := progressbar.NewOptions(len(searchResults.Items),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan] Crawling Github repositories. Looking for Cosmos projects...:"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=",
			SaucerHead:    "[green]>",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	for _, r := range searchResults.Items {
		_ = bar.Add(1)
		cosmosSdk := ""
		// If it's a Golang project check if it uses the Cosmos SDK
		if strings.ToLower(r.Language) == "go" {
			cosmosSdk, err = client.IsCosmosSDK(r.Owner.Login, r.Name, r.DefaultBranch)
		}

		project := model.Project{
			Name:        r.Name,
			Owner:       r.Owner.Login,
			Url:         r.HTMLURL,
			Description: r.Description,
			Language:    r.Language,
			License:     r.License.SpdxID,
			Stars:       r.StargazersCount,
			Forks:       r.ForksCount,
			LastCommit:  r.PushedAt,
			Branch:      r.DefaultBranch,
			CosmosSDK:   cosmosSdk,

		}

		// Logic to remove Azure CosmosDB listings
		if !strings.Contains(strings.ToLower(project.Description), "cosmosdb") {
			projects = append(projects, project)
		}
	}

	_ = bar.Finish()
	switch sortBy {
	case "stars":
		sort.Sort(sort.Reverse(ByStars(projects)))
	case "forks":
		sort.Sort(sort.Reverse(ByForks(projects)))
	default:
		sort.Sort(sort.Reverse(ByLastCommit(projects)))
	}

	return projects, nil
}

// Print Projects in Table format
func PrintProjectsTable(projects []model.Project) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "OWNER"},
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "URL"},
			//{Align: simpletable.AlignCenter, Text: "DESCRIPTION"},
			{Align: simpletable.AlignCenter, Text: "LANGUAGE"},
			{Align: simpletable.AlignCenter, Text: "COSMOS SDK (DEFAULT BRANCH)"},
			{Align: simpletable.AlignCenter, Text: "LICENSE"},
			{Align: simpletable.AlignCenter, Text: "STARS"},
			{Align: simpletable.AlignCenter, Text: "FORKS"},
			{Align: simpletable.AlignCenter, Text: "LAST COMMIT"},
		},
	}
	count := 0

	for _, p := range projects {
		//var description string
		//if len(p.Description) > 28 {
		//	description = p.Description[0:26] + "..."
		//} else {
		//	description = p.Description
		//}
		var sdkBranch string
		if len(p.CosmosSDK) > 0 {
			sdkBranch = p.CosmosSDK + " (" + p.Branch + ")"
		} else {
			sdkBranch = ""
		}
		row := []*simpletable.Cell{
			{Text: p.Owner},
			{Text: p.Name},
			{Text: p.Url},
			//{Text: strings.ToValidUTF8(description, "")},
			{Text: p.Language},
			{Text: sdkBranch},
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
			{Align: simpletable.AlignRight, Text: "Total Repositories Found:"},
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