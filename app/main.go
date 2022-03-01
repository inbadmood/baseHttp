package main

import (
	"baseHttp/app/api"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "baseHttp",
	Short: "baseHttp",
	Long:  `baseHttp is a project which develop with Clean Architecture in Go (Golang).`,
	Run: func(cmd *cobra.Command, args []string) {
		// default start api server
		api.StartAPIServer()
	},
}

// aPICmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "baseHttp restful api",
	Long:  `baseHttp is a restful api service which develop with Clean Architecture in Go (Golang).`,
	Run: func(cmd *cobra.Command, args []string) {
		api.StartAPIServer()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func main() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
