package main

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "turtl",
	Short: "turtl - a little virtual machine",
	Long:  `turtl - a little virtual machine`,
}
