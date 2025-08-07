package cmd

import "github.com/spf13/cobra"

type flags struct {
	Set   string
	Auth  bool
	Check bool
	Unset bool
}

func cmdFlags(c *cobra.Command, f *flags) {
	c.Flags().StringVarP(&f.Set, "set", "s", "", "set github token into os environment variable (e.g., gitty -s=your_github_token)")
	c.Flags().BoolVarP(&f.Auth, "auth", "a", false, "print authenticated username")
	c.Flags().BoolVarP(&f.Check, "check", "c", false, "check client status and remaining rate limit")
	c.Flags().BoolVarP(&f.Unset, "unset", "u", false, "unset github token from os environment variable")
}
