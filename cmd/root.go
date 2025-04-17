package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Const for version
const (
	Version = "0.1.0"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cosmocope",
	Short: "A tool to find Cosmos (cosmos.network) related projects and technologies",
	Long: `
This tool allows you to crawl Github in order to discover 
Cosmos (cosmos.network) related projects. Currently, the 
only source crawled by the tool is Github. It leverages 
Github's search API to discover projects and only public 
repositories can be crawled.`,
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetVersionTemplate(`{{.Use}} - version: {{.Version}}`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
