package main

import (
	"baseApiServer/service/server"

	"github.com/spf13/cobra"
)

func main() {
	cmd.AddCommand(PlatformApiServer)

	cmd.Execute()
}

var cmd = &cobra.Command{
	Use:   "PlatformApiServer",
	Short: "PlatformApiServer",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var PlatformApiServer = &cobra.Command{
	Use:   "server [file_location option]",
	Short: "Execute PlatformApiServer.",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}
