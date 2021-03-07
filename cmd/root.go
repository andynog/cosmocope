package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	if version {
	//		fmt.Printf("cosmocope:\r\tversion: 0.0.1\r\t%sgolang:" + runtime.Version())
	//	} else {
	//		cmd.Help()
	//		os.Exit(0)
	//	}
	//},
	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetVersionTemplate(`
{{.Use}} - version: {{.Version}}
`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Disable configuration for now
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cosmocope.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Displays current tool version")
}

// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		// Search config in home directory with name ".cosmocope" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".cosmocope")
//	}
//
//	viper.AutomaticEnv() // read in environment variables that match
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}
