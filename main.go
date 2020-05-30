package main

import (
	"cosmoscope/client"
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
    "github.com/dustin/go-humanize"
	"github.com/schollz/progressbar/v3"
)

func main() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "OWNER"},
			{Align: simpletable.AlignCenter, Text: "URL"},
			{Align: simpletable.AlignCenter, Text: "DESCRIPTION"},
			{Align: simpletable.AlignCenter, Text: "STARS"},
			{Align: simpletable.AlignCenter, Text: "FORKS"},
			{Align: simpletable.AlignCenter, Text: "UPDATED"},
		},
	}
	countCosmosProjects := 0
	var searchResults client.GithubSearchResult
	searchResults = client.SearchGithubProjects()

	// Progress Bar
	bar := progressbar.NewOptions(len(searchResults.Items),
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan] Crawling Github repositories with 'cosmos-sdk' topic:"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=",
			SaucerHead:    "[green]>",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	for _, r := range searchResults.Items {
		isCosmosProject := client.LookForModules(r.HTMLURL)
		bar.Add(1)
		var description string
		if len(r.Description) > 70 {
			description = r.Description[0:68] + "..."
		} else {
			description = r.Description
		}
		if isCosmosProject {
			row := []*simpletable.Cell{
				{Text: r.Name},
				{Text: r.Owner.Login},
				{Text: r.HTMLURL},
				{Text: strings.ToValidUTF8(description, "")},
				{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", r.StargazersCount)},
				{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", r.ForksCount)},
				{Text: humanize.Time(r.UpdatedAt)},
			}
			table.Body.Cells = append(table.Body.Cells, row)
			countCosmosProjects++
		}
	}
	bar.Finish()
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
