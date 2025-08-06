package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "unknown"

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of pgit",
		Long:  "All software has versions. This is pgit's",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("pgit v%s\n", Version)
		},
	}
}
