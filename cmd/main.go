package cmd

import (
	"fmt"
	"partial-git/internal"
	"partial-git/internal/token"

	"github.com/spf13/cobra"
)

var f flags

var rootCmd = &cobra.Command{
	Use:   "pgit <github-url>",
	Short: "Download files or folders from a repo",
	Args:  cobra.MaximumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateFlags(args)
	},
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

func validateFlags(args []string) error {
	flagCount := 0
	if f.Auth {
		flagCount++
	}
	if f.Check {
		flagCount++
	}
	if f.Unset {
		flagCount++
	}
	if f.Set != "" {
		flagCount++
	}

	if flagCount > 1 {
		return fmt.Errorf("only one of --set, --auth, --check, or --unset can be used at a time")
	}

	switch {
	case f.Set != "":
		if err := token.ValidateToken(f.Set); err != nil {
			return fmt.Errorf("invalid GitHub token: %w", err)
		}
		if len(args) > 0 {
			return fmt.Errorf("cannot provide GitHub URL when setting token")
		}
		return nil

	case f.Auth || f.Check || f.Unset:
		if len(args) > 0 {
			return fmt.Errorf("no arguments expected when using --auth, --check, or --unset")
		}
		return nil

	default:
		if len(args) == 0 {
			return fmt.Errorf("GitHub URL is required for download operations")
		}
		if len(args) > 1 {
			return fmt.Errorf("only one GitHub URL is allowed, got %d URLs", len(args))
		}
		return nil
	}
}

func Execute() error {
	cmdFlags(rootCmd, &f)
	rootCmd.AddCommand(versionCmd())
	return rootCmd.Execute()
}
