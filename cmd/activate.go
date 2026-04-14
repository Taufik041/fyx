package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Install the shell hook to enable command correction",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("activate: coming soon")
	},
}

func init() {
	rootCmd.AddCommand(activateCmd)
}
