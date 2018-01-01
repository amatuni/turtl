package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type shellCommand func([]string, *ShellContext) error

func cmdQuit(args []string, ctx *ShellContext) error {
	fmt.Printf("\nbye bye :)\n\n")
	os.Exit(0)
	return nil
}

func cmdHelp(args []string, ctx *ShellContext) error {
	_, err := fmt.Println(shellHelpStr)
	if err != nil {
		return err
	}
	return nil
}

func cmdLoad(args []string, ctx *ShellContext) error {
	fmt.Println("hello from load")
	return nil
}

func cmdRun(args []string, ctx *ShellContext) error {
	fmt.Println("hello from run")
	return nil
}

func cmdSave(args []string, ctx *ShellContext) error {
	buff := bytes.Buffer{}
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(ctx)

	var name string
	if ctx.name == "" {
		name = string(ctx.id)
	} else {
		name = ctx.name
	}

	fpath := fmt.Sprintf("%s.%s", name, sessionSuffix)
	f, err := os.Create(fpath)
	if err != nil {
		return err
	}

	_, err = f.Write(buff.Bytes())
	if err != nil {
		return err
	}

	fmt.Printf("session saved to: %s\n", fpath)
	return nil
}

func cmdReset(args []string, ctx *ShellContext) error {
	ctx.history = make([]string, 0, 0)
	ctx.line = ""
	ctx.count = 0
	ctx.currentPos = 0
	ctx.index = 0
	return nil
}

func cmdName(args []string, ctx *ShellContext) error {
	if len(args) > 0 {
		ctx.name = args[0]
	}
	fmt.Printf("session name set to: %s\n", ctx.name)
	return nil
}

const shellHelpStr = `
Available Commands:

help [command]    show usage information
quit              quit the shell
load <program>    load a program
run  <program>    run a program
save [path]       save the current shell session to a file
reset             reset the shell to the initial state
name [name]       set the session name


Interpreter commands should be prefixed with 
an @, for example: @help

`
