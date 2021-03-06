/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"cosmocope/controller"
	"fmt"
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
		projects, err := controller.GetProjects()
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