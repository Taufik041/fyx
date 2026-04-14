package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Re-scan PATH to pick up newly installed tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("refresh: coming soon")
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
