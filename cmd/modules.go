package cmd

import (
	"fmt"
	"github.com/andynog/cosmocope/v2/controller"
	"github.com/spf13/cobra"
	"os"
)

// modulesCmd represents the modules command
var modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "Find Cosmos SDK modules",
	Long: `
This command searches for projects on Github that are tagged 
with the 'cosmos-sdk' topic and the tool crawls each repository 
looking for a folder named 'x' in the repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch projects
		projects, err := controller.GetProjects("commit")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Fetch modules for projects
		modules, err := controller.FindModulesInProjects(projects)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if jsonOutput {
			controller.PrintModulesJSON(modules)
		} else {
			controller.PrintModulesTable(modules)
		}
	},
}

func init() {
	discoverCmd.AddCommand(modulesCmd)
	modulesCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output results to JSON")
}