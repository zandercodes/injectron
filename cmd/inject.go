package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(injectCmd)
}

var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Inject code into an Electron App",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify an Electron App to inject into.\n" +
				"You can use the 'help inject' command for more information.")
			return
		}
	},
}
