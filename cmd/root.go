package cmd

import (
	"fmt"
	"gold/cmd/web"
	"os"

	"github.com/spf13/cobra"
)

// gold
var rootCmd = &cobra.Command{
	Use:   "gold",
	Short: "Explore the Secrets of Gold",
	Long:  `Gold, one of the most precious metals in nature, symbolizes wealth and sanctity.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Gold Exploration Tool!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// TODO : add subcommands
	rootCmd.AddCommand(
		web.WebCmd,
	)
}
