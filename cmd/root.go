package cmd

import (
	"fmt"
	"github.com/myugen/go-kit-example/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "App entry command",
	RunE: func(_ *cobra.Command, _ []string) error {
		fmt.Printf("Server.Port: %s", config.Config.Server.Port)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringP("port", "p", "8080", "Port to run Application server on")

	cobra.OnInitialize(config.Init(config.ConfigFlag{
		Field: "server.port",
		Flag:  rootCmd.Flags().Lookup("port"),
	}))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
