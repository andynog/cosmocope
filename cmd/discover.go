package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// discoverCmd represents the discover command
var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Find projects, tools, libraries and modules for the Cosmos ecosystem",
	Long:  `Find projects, tools, libraries and modules for the Cosmos ecosystem`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(discoverCmd)
}
