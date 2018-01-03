package main

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CompileCmd)
}

var CompileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile some code",
	Long:  `Compile some turtl source code files into a bytecode executable`,
	Run: func(cmd *cobra.Command, args []string) {
		err := validateCompileArgs(args)
		if err != nil {
			log.Fatal(err)
		} else {

		}
	},
}

func validateCompileArgs(args []string) error {

	return nil
}
