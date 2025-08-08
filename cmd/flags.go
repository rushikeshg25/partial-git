package cmd

import "github.com/spf13/cobra"

type flags struct {
	Set   string
	Auth  bool
	Check bool
	Unset bool
}

func cmdFlags(c *cobra.Command, f *flags) {
	c.Flags().StringVarP(&f.Set, "set", "s", "", "store GitHub Personal Access Token in shell profile")
	c.Flags().BoolVarP(&f.Auth, "auth", "a", false, "show authenticated user information")
	c.Flags().BoolVarP(&f.Check, "check", "c", false, "check token status and availability")
	c.Flags().BoolVarP(&f.Unset, "unset", "u", false, "remove GitHub token from shell profile")
}
