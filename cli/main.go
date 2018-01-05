package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/andreiamatuni/turtl"
)

func checkRuntimeConditions() error {
	if turtl.IntSize() < 64 {
		return errors.New("[runtime]: CPU must be >= 64 bit")
	}
	return nil
}

func main() {
	// make sure some invariants hold
	err := checkRuntimeConditions()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// root process
	if err = RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
