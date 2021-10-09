package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl repositories from a pre-defined list of Cosmos SDK projects.",
	Long:  `Crawl repositories from a pre-defined list of Cosmos SDK projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "Output results to JSON")
	crawlCmd.PersistentFlags().StringVarP(&sortBy, "repository list", "l", "", "List of URLs of GitHub Cosmos SDK repositories to crawl")
}
