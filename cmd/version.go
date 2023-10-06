package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "All software has versions. This is IP Tracker's",
	// Long:  `All software has versions. This is IP Tracker's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.0")
	},
}
