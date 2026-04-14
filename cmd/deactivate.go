package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Remove the shell hook to disable command correction",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deactivate: coming soon")
	},
}

func init() {
	rootCmd.AddCommand(deactivateCmd)
}
