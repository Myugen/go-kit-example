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
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Config => %+v", config.Config)
	},
}

func init() {
	rootCmd.Flags().StringP("environment", "e", "development", "Environment where Application is in")
	rootCmd.Flags().StringP("port", "p", "8080", "Port to run Application server on")

	cobra.OnInitialize(config.Init(
		config.ConfigFlag{
			Field: "server.port",
			Flag:  rootCmd.Flags().Lookup("port"),
		},
		config.ConfigFlag{
			Field: "environment",
			Flag:  rootCmd.Flags().Lookup("environment"),
		}))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
