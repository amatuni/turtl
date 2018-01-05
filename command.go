package turtl

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

/*
ShellCommand is a type of function which executes a shell
command
*/
type ShellCommand func([]string, *ShellContext) error

func CmdQuit(args []string, ctx *ShellContext) error {
	fmt.Printf("\nbye bye :)\n\n")
	os.Exit(0)
	return nil
}

func CmdHelp(args []string, ctx *ShellContext) error {
	_, err := fmt.Println(shellHelpStr)
	if err != nil {
		return err
	}
	return nil
}

func CmdLoad(args []string, ctx *ShellContext) error {
	fmt.Println("hello from load")
	return nil
}

func CmdRun(args []string, ctx *ShellContext) error {
	fmt.Println("hello from run")
	return nil
}

func CmdSave(args []string, ctx *ShellContext) error {
	buff := bytes.Buffer{}
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(ctx)

	var name string
	if len(args) > 0 {

	} else {
		if ctx.Name == "" {
			name = string(ctx.ID)
		} else {
			name = ctx.Name
		}
	}

	fpath := fmt.Sprintf("%s.%s", name, sessionSuffix)
	f, err := os.Create(fpath)
	defer f.Close()
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

func CmdReset(args []string, ctx *ShellContext) error {
	ctx.History = make([]string, 0, 0)
	ctx.Line = ""
	ctx.Count = 0
	ctx.CurrentPos = 0
	ctx.Index = 0
	return nil
}

func CmdName(args []string, ctx *ShellContext) error {
	if len(args) > 0 {
		ctx.Name = args[0]
		fmt.Printf("session name set to: %s\n", ctx.Name)
		return nil
	}
	fmt.Printf("session name: %s\n", ctx.Name)
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


Shell commands should be prefixed with 
an @, for example: @help

`
