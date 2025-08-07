package cmd

import (
	"partial-git/internal"

	"github.com/spf13/cobra"
)

var f flags

var rootCmd = &cobra.Command{
	Use:   "pgit <github-url>",
	Short: "Download files or folders from a repo",
	Args:  cobra.ArbitraryArgs, // Allow any number of arguments
	Run: func(cmd *cobra.Command, args []string) {
		internalFlags := internal.Flags{
			Set:   f.Set,
			Auth:  f.Auth,
			Check: f.Check,
			Unset: f.Unset,
		}
		internal.Run(internalFlags, args)
	},
}

func Execute() error {
	cmdFlags(rootCmd, &f)
	rootCmd.AddCommand(versionCmd())
	return rootCmd.Execute()
}
