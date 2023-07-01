package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nets v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
