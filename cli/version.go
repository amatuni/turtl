package main

import (
	"fmt"

	"github.com/andreiamatuni/turtl"
	"github.com/spf13/cobra"
)

func version() string {
	return fmt.Sprintf("%s\n%s", turtl.VersionNum(), turtl.GitHash()[:12])
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nturtl v" + version() + "\n\n")
	},
}
