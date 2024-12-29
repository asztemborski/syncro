package cmd

import (
	"github.com/asztemborski/syncro/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run application",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Run(cmd.Context())
	},
}
