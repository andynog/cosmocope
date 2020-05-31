package controller

import (
	"cosmocope/model"
	"encoding/json"
	"fmt"
	"strings"

	"cosmocope/client"
	"github.com/alexeyco/simpletable"
	"github.com/dustin/go-humanize"
	"github.com/schollz/progressbar/v3"
)

func PrintTable(projects []model.Project) {
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

	// Progress Bar
	bar := progressbar.NewOptions(len(projects),
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan] Crawling Github repositories:"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=",
			SaucerHead:    "[green]>",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	for _, p := range projects {
		isCosmosProject := client.LookForModules(p.Url)
		bar.Add(1)
		var description string
		if len(p.Description) > 60 {
			description = p.Description[0:58] + "..."
		} else {
			description = p.Description
		}
		if isCosmosProject {
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

func PrintJSON(projects []model.Project) {
	json, err := json.MarshalIndent(projects, "", " ")
	if err != nil {
		fmt.Println("Failed to print JSON results:", err)
		return
	}
	fmt.Println(string(json))
}
