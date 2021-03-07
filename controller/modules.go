package controller

import (
	"cosmocope/client"
	"cosmocope/model"
	"encoding/json"
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/schollz/progressbar/v3"
)


func FindModulesInProjects(projects []model.Project) (result []model.Module, err error) {
	var modules []model.Module
	// Progress Bar
	bar := progressbar.NewOptions(len(projects),
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan] Crawling Github repositories. Looking for Cosmos modules...:"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=",
			SaucerHead:    "[green]>",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	for _, p := range projects {
		_ = bar.Add(1)
		hasModulesFolder := client.LookForModules(p.Url)
		if hasModulesFolder {
			result, err := client.GetContentFromGithub(p.Owner, p.Name)
			if err != nil {
				return nil, fmt.Errorf("error fetching modules: %s", err)
			}
			if len(result) > 0 {
				for _, m := range result {
					if m.Type == "dir" {
						module := model.Module{Name: m.Name, Owner: p.Owner, Repo: p.Name, Url: m.HTMLURL}
						modules = append(modules, module)
					}
				}
			}
		}
	}
	_ = bar.Finish()
	return modules, nil
}

// Print Modules in Table format
func PrintModulesTable(modules []model.Module) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "MODULE"},
			{Align: simpletable.AlignCenter, Text: "REPOSITORY"},
			{Align: simpletable.AlignCenter, Text: "URL"},
		},
	}
	count := 0

	for _, m := range modules {
		row := []*simpletable.Cell{
			{Text: m.Name},
			{Text: m.Owner + "/" + m.Repo},
			{Text: m.Url},
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

// Print Modules in JSON format
func PrintModulesJSON(modules []model.Module) {
	json, err := json.MarshalIndent(modules, "", " ")
	if err != nil {
		fmt.Println("Failed to print JSON results:", err)
		return
	}
	fmt.Println(string(json))
}
