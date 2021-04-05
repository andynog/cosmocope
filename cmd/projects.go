package cmd

import (
	"flag"
	"fmt"
	"github.com/andynog/cosmocope/v2/controller"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var jsonOutput bool
var sortBy string

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Find Cosmos SDK projects",
	Long:  `
This command searches for projects on Github that are tagged with the 'cosmos-sdk' topic`,
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		if len(sortBy) > 0 {
			if err := validateSort(); err != nil {
				fmt.Fprintln(os.Stderr, "Error: ", err)
				os.Exit(1)
			}
		}
		projects, err := controller.GetProjects(sortBy)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if jsonOutput {
			controller.PrintProjectsJSON(projects)
		} else {
			controller.PrintProjectsTable(projects)
		}
	},
}

func init() {
	discoverCmd.AddCommand(projectsCmd)
	projectsCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output results to JSON")
	projectsCmd.PersistentFlags().StringVarP(&sortBy, "sort", "s", "commit", "Sort by the specified field. Valid values are 'commit', 'stars', 'forks'")
}

func validateSort() error {
	validSort := []string{"commit", "stars", "forks"}
	for _, sortFlagValue := range validSort {
		if sortBy == sortFlagValue {
			// sort is valid
			return nil
		}
	}
	return fmt.Errorf("value '%s' is invalid for flag 'sort', valid values are: %v", sortBy, strings.Join(validSort[:], ","))
}
