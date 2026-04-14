package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var correctCmd = &cobra.Command{
	Use:   "correct",
	Short: "Suggest a correction for a mistyped command (called by shell hook)",
	// Hidden because users don't call this directly — the shell hook does
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("correct: coming soon")
	},
}

func init() {
	rootCmd.AddCommand(correctCmd)
}
