package cmd

import (
	"cosmocope/controller"
	"fmt"

	"github.com/spf13/cobra"
)

var jsonOutput bool

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Find Cosmos projects",
	Long:  `This command searches for projects on Github that are tagged with the 'cosmos-sdk' topic`,
	Run: func(cmd *cobra.Command, args []string) {
		projects := controller.GetProjects()
		if jsonOutput {
			controller.PrintJSON(projects)
		} else {
			fmt.Println("Discovering projects...")
			controller.PrintTable(projects)
		}
	},
}

func init() {
	discoverCmd.AddCommand(projectsCmd)
	projectsCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output results to JSON")
}
