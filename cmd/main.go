package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pgit",
	Short: "A partial git implementation",
	Long:  "A partial git implementation with various git commands",
}

func Execute() error {
	rootCmd.AddCommand(versionCmd())
	return rootCmd.Execute()
}

func GetVersion() string {
	return "unknown"
}
