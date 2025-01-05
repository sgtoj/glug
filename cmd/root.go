package cmd

import (
	"github.com/spf13/cobra"
)

var version = "0.0.0"

var rootCmd = &cobra.Command{
	Use:   "glug",
	Short: "Tool to download other tools.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose logging")
	rootCmd.AddCommand(getCmd)
}
