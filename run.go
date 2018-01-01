package main

import (
	"errors"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(RunCmd)
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run some code",
	Long:  `Run the specified code`,
	Run: func(cmd *cobra.Command, args []string) {
		valid_err := validate_run_args(args)
		if valid_err != nil {
			log.Fatal(valid_err)
		} else {

		}
	},
}

var invalid_filetype_err = errors.New("run: invalid file types")

func validate_run_args(args []string) error {
	for _, s := range args {
		if !strings.HasSuffix(s, ".turtl") {
			return invalid_filetype_err
		}
	}
	return nil
}
