package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.0-dev"
var commit = "0000000"
var date = "0000-00-00T00:00:00Z"
var builtBy = "unknown"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Injectron",
	Long:  `All software has versions. This is Injectron's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Injectron v%s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Date: %s\n", date)
		fmt.Printf("Built by: %s\n", builtBy)
	},
}
