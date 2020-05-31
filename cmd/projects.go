package cmd

import (
	"cosmocope/controller"
	"fmt"

	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Find Cosmos SDK projects",
	Long: `This command searches for projects on Github that are tagged with the 'cosmos-sdk' topic`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Discovering projects...")
		projects := controller.GetProjects()
		controller.PrintTable(projects)
	},
}

func init() {
	discoverCmd.AddCommand(projectsCmd)
	projectsCmd.PersistentFlags().String("json", "", "Output results to JSON")
}
