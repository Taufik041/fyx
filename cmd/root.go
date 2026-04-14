package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command — running `fyx` with no subcommand hits this
var rootCmd = &cobra.Command{
	Use:   "fyx",
	Short: "A smart CLI companion that fixes commands and browses tools",
	Long: `fyx does two things:
  - Fixes mistyped commands using AI
  - Lets you browse any tool's subcommands interactively`,

	// If someone just runs "fyx" with no subcommand, show help
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute is called by main.go — this is what starts everything
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
