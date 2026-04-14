package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update your AI provider or API key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config: coming soon")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
