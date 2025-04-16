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
	Run: func(cmd *cobra.Command, _ []string) {
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