package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse [tool]",
	Short: "Interactively browse a tool's subcommands",
	// DisableFlagParsing lets us pass the tool name freely
	DisableFlagParsing: true,
	Args:               cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("browse: would show commands for '%s' — coming soon\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(browseCmd)
}
