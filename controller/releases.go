package controller

import (
	"encoding/json"
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/andynog/cosmocope/v2/client"
	"github.com/andynog/cosmocope/v2/model"
	"github.com/dustin/go-humanize"
	"github.com/schollz/progressbar/v3"
)

// GetReleases fetches releases from a Github repository
func GetReleases(url string) (result []model.Release, err error) {
	var releases []model.Release
	owner, name, err := ParseGithubURL(url)
	if err != nil {
		return nil, fmt.Errorf("problems parsing Github URL")
	}
	releaseResult, err := client.GetReleasesFromGithub(owner, name)
	if err != nil {
		return nil, fmt.Errorf("problems fetching releases: %w", err)
	}

	// Progress bar
	bar := progressbar.NewOptions(len(releaseResult),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(fmt.Sprintf("[cyan] Crawling [green] %s [cyan] Github repository. Looking for releases...:", url)),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=",
			SaucerHead:    "[green]>",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	for _, r := range releaseResult {
		_ = bar.Add(1)

		release := model.Release{
			Name:        r.Name,
			TagName:     r.TagName,
			URL:         r.HTMLURL,
			Draft:       r.Draft,
			Description: r.Body,
			PublishedAt: r.PublishedAt,
		}

		releases = append(releases, release)
	}

	_ = bar.Finish()

	return releases, nil
}

// PrintReleasesTable prints Releases in Table format
func PrintReleasesTable(projects []model.Release) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "TAG NAME"},
			{Align: simpletable.AlignCenter, Text: "URL"},
			{Align: simpletable.AlignCenter, Text: "PUBLISHED"},
		},
	}
	count := 0

	for _, p := range projects {

		// Ignore draft releases
		if !p.Draft {
			row := []*simpletable.Cell{
				{Text: p.Name},
				{Text: p.TagName},
				{Text: p.URL},
				{Text: humanize.Time(p.PublishedAt)},
			}
			table.Body.Cells = append(table.Body.Cells, row)
			count++
		}
	}
	fmt.Println("\r")

	// Table Footer
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: "Total Releases Found:"},
			{Align: simpletable.AlignLeft, Span: len(table.Header.Cells) - 1, Text: fmt.Sprintf("%d", count)},
		},
	}

	// Print table
	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

// PrintReleasesJSON prints Releases in JSON format
func PrintReleasesJSON(releases []model.Release) {
	json, err := json.MarshalIndent(releases, "", " ")
	if err != nil {
		fmt.Println("Failed to print JSON results:", err)
		return
	}
	fmt.Println(string(json))
}