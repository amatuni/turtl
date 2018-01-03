package turtl_test

import (
	"fmt"
	"log"

	"github.com/andreiamatuni/turtl"
)

func Example() {

	compiledCode := turtl.Program{ // not real code here, just dummy bytecode
		0x12,
		0x34,
		0x56,
		0x78,
	}

	// construct a new turtl VM
	vm := turtl.NewVM()

	// load some code into it
	if err := vm.LoadProgram(compiledCode); err != nil {
		log.Fatal(err)
	}

	// run and print the result
	if result, err := vm.Run(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(result)
	}
}
