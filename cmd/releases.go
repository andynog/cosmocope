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
	"flag"
	"fmt"
	"github.com/andynog/cosmocope/v2/controller"
	"github.com/spf13/cobra"
	"os"
)

var repo string

// modulesCmd represents the modules command
var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Find Github repository releases",
	Long: `
This command lists the releases for a Github repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		if len(repo) == 0 {
			fmt.Fprintln(os.Stderr, "Please specify a valid Github repository URL")
			_ = cmd.Usage()
			os.Exit(1)
		}

		releases, err := controller.GetReleases(repo)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if jsonOutput {
			controller.PrintReleasesJSON(releases)
		} else {
			controller.PrintReleasesTable(releases)
		}
	},
}

func init() {
	discoverCmd.AddCommand(releasesCmd)
	releasesCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output results to JSON")
	releasesCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "Specify the Github Repository URL (e.g. https://github.com/cosmos/cosmos-sdk")
}