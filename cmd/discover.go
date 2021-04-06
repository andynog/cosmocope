package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// discoverCmd represents the discover command
var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Find repositories, releases, and modules for projects in the Cosmos ecosystem.",
	Long:  `Find repositories, releases, and modules for projects in the Cosmos ecosystem.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(discoverCmd)
}
