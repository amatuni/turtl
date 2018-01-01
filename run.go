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
		err := validateRunArgs(args)
		if err != nil {
			log.Fatal(err)
		} else {

		}
	},
}

var ErrInvalidFiletype = errors.New("run: invalid file types")

func validateRunArgs(args []string) error {
	for _, s := range args {
		if !strings.HasSuffix(s, ".turtl") {
			return ErrInvalidFiletype
		}
	}
	return nil
}
