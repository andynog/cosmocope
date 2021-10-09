package cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
)

var list string

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl repositories from a pre-defined list of Cosmos SDK projects.",
	Long:  `Crawl repositories from a pre-defined list of Cosmos SDK projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		flag.Parse()
		if list == "" {
			fmt.Println("Please specify the repository list filename")
		} else {
			fmt.Printf("Crawling the repositories in %s\n", list)
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "output results to JSON")
	crawlCmd.PersistentFlags().StringVarP(&list, "list", "l", "", "path of the file with the list of Github repositories to crawl")
}
