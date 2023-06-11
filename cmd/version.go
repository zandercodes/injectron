package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.0-dev"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Injectron",
	Long:  `All software has versions. This is Injectron's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Injectron v%s", version)
	},
}
