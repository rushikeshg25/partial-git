package cmd

import (
	"fmt"
	"net/url"
	"partial-git/internal"
	"partial-git/internal/token"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var f flags

var rootCmd = &cobra.Command{
	Use:   "pgit <github-url>",
	Short: "Download files or folders from a repo",
	Args:  cobra.ArbitraryArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateInput(args)
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

func validateInput(args []string) error {
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

	if f.Set != "" {
		if err := token.ValidateToken(f.Set); err != nil {
			return fmt.Errorf("invalid GitHub token: %w", err)
		}
		if len(args) > 0 {
			return fmt.Errorf("cannot provide GitHub URL when setting token")
		}
		return nil
	}

	if f.Auth || f.Check || f.Unset {
		if len(args) > 0 {
			return fmt.Errorf("no arguments expected when using --auth, --check, or --unset")
		}
		return nil
	}

	if len(args) == 0 {
		return fmt.Errorf("GitHub URL is required for download operations")
	}

	for _, arg := range args {
		if err := validateGitHubURL(arg); err != nil {
			return fmt.Errorf("invalid GitHub URL '%s': %w", arg, err)
		}
	}

	return nil
}

func validateGitHubURL(urlStr string) error {
	if urlStr == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}

	if parsedURL.Host != "github.com" && parsedURL.Host != "www.github.com" {
		return fmt.Errorf("URL must be from github.com")
	}

	if parsedURL.Scheme == "" {
		return fmt.Errorf("URL must include scheme (https://)")
	}

	if parsedURL.Scheme != "https" && parsedURL.Scheme != "http" {
		return fmt.Errorf("URL scheme must be http or https")
	}

	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathParts) < 2 {
		return fmt.Errorf("GitHub URL must include owner and repository (e.g., https://github.com/owner/repo)")
	}

	if pathParts[0] == "" || pathParts[1] == "" {
		return fmt.Errorf("GitHub URL must include valid owner and repository names")
	}

	nameRegex := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-_])*[a-zA-Z0-9]$|^[a-zA-Z0-9]$`)
	if !nameRegex.MatchString(pathParts[0]) {
		return fmt.Errorf("invalid GitHub owner name: %s", pathParts[0])
	}
	if !nameRegex.MatchString(pathParts[1]) {
		return fmt.Errorf("invalid GitHub repository name: %s", pathParts[1])
	}

	return nil
}

func Execute() error {
	cmdFlags(rootCmd, &f)
	rootCmd.AddCommand(versionCmd())
	return rootCmd.Execute()
}
