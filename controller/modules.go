package controller

import (
	"cosmocope/client"
	"cosmocope/model"
	"encoding/json"
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/schollz/progressbar/v3"
)

func FindModulesInProjects(projects []model.Project) []model.Module {
	var modules []model.Module
	// Progress Bar
	bar := progressbar.NewOptions(len(projects),
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan] Crawling Github Cosmos repositories. Looking for modules...:"),
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
		if isCosmosProject {
			module := model.Module{ Name: p.Name }
			modules = append(modules, module)
		}
	}
	bar.Finish()
	return modules
}

// Print Modules in Table format
func PrintModulesTable(modules []model.Module) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "NAME"},
		},
	}
	count := 0

	for _, p := range modules {
		row := []*simpletable.Cell{
			{Text: p.Name},
		}
		table.Body.Cells = append(table.Body.Cells, row)
		count++
	}
	fmt.Println("\r")

	// Table Footer
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			//{Align: simpletable.AlignRight, Text: "Total Modules"},
			{Align: simpletable.AlignLeft, Span: 1, Text: fmt.Sprintf("Total: %d", count)},
		},
	}

	// Print table
	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

// Print Modules in JSON format
func PrintModulesJSON(modules []model.Module) {
	json, err := json.MarshalIndent(modules, "", " ")
	if err != nil {
		fmt.Println("Failed to print JSON results:", err)
		return
	}
	fmt.Println(string(json))
}