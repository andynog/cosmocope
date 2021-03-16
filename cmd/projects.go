package cmd

import (
	"fmt"
	"github.com/andynog/cosmocope/controller"
	"github.com/spf13/cobra"
	"os"
)

var jsonOutput bool

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Find Cosmos SDK projects",
	Long:  `
This command searches for projects on Github that are tagged with the 'cosmos-sdk' topic`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := controller.GetProjects()
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
}
