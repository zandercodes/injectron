package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

var rootCmd = &cobra.Command{
	Use:   "injectron",
	Short: "Injectron another way to inject Electron Apps",
	Long:  `Injectron is a tool to inject Electron Apps with custom code.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
