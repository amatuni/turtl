package main

// terminal key capturing code borrowed from
// github.com/d4l3k/go-pry
import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	tty "github.com/mattn/go-tty"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ShellCmd)
}

/*
ShellCmd handles launching the turtl shell environment
*/
var ShellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start shell",
	Long:  `Start the turtl shell, a REPL for for the virtual machine`,
	Run: func(cmd *cobra.Command, args []string) {
		err := repl()
		if err != nil {
			log.Fatal(err)
		}
	},
}

/*
ShellContext holds the current shell state
*/
type ShellContext struct {
	id         uint     // session id
	name       string   // session name
	vm         *VM      // the VM instance loaded into this shell
	history    []string // the record of all past input
	line       string   // the current line
	index      int
	count      int
	currentPos int
}

func NewShellContext() ShellContext {
	return ShellContext{
		history:    []string{},
		index:      0,
		count:      0,
		currentPos: 0,
		line:       "",
	}
}

/*
	repl runs the turtl read-eval-print loop
*/
func repl() error {
	t, err := tty.Open()
	if err != nil {
		return err
	}
	defer t.Close()

	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Println(infoGraphic())

	ctx := NewShellContext()

	r := rune(0)
	for {
		prompt := fmt.Sprintf("%s%d%s %s :> ", red("["), ctx.currentPos, red("]"), green("turtl"))
		fmt.Fprintf(os.Stdout, "\r\033[K%s%s \033[0J\033[%dD", prompt, ctx.line, len(ctx.line)-ctx.index+1)
		bPrev := r

		r = 0
		for r == 0 {
			var err error
			r, err = t.ReadRune()
			if err != nil {
				return err
			}
		}
		switch r {
		default:
			if bPrev == 27 && r == 91 {
				continue
			} else if bPrev == 91 {
				switch r {
				case 66: // Down
					ctx.currentPos++
					if len(ctx.history) < ctx.currentPos {
						ctx.currentPos = len(ctx.history)
					}
					if len(ctx.history) == ctx.currentPos {
						ctx.line = ""
					} else {
						ctx.line = ctx.history[ctx.currentPos]
					}
					ctx.index = len(ctx.line)
				case 65: // Up
					ctx.currentPos--
					if ctx.currentPos < 0 {
						ctx.currentPos = 0
					}
					if len(ctx.history) > 0 {
						ctx.line = ctx.history[ctx.currentPos]
					}
					ctx.index = len(ctx.line)
				case 67: // Right
					ctx.index++
					if ctx.index > len(ctx.line) {
						ctx.index = len(ctx.line)
					}
				case 68: // Left
					ctx.index--
					if ctx.index < 0 {
						ctx.index = 0
					}
				}
				continue
			} else if bPrev == 51 && r == 126 { // DELETE
				if len(ctx.line) > 0 && ctx.index < len(ctx.line) {
					ctx.line = ctx.line[:ctx.index] + ctx.line[ctx.index+1:]
				}
				if ctx.index > len(ctx.line) {
					ctx.index = len(ctx.line)
				}
				continue
			}
			ctx.line = ctx.line[:ctx.index] + string(r) + ctx.line[ctx.index:]
			ctx.index++
		case 127, '\b': // Backspace
			if len(ctx.line) > 0 && ctx.index > 0 {
				ctx.line = ctx.line[:ctx.index-1] + ctx.line[ctx.index:]
				ctx.index--
			}
			if ctx.index > len(ctx.line) {
				ctx.index = len(ctx.line)
			}
		case 27:
		case 10, 13: // Enter
			fmt.Fprintln(os.Stdout, "\033[100000C\033[0J")
			ctx.count++
			ctx.currentPos = ctx.count
			ctx.history = append(ctx.history, ctx.line)
			eval(ctx.line, &ctx)
			ctx.line = ""
			ctx.index = 0
		}
	}
}

/*
eval passes shell input through the appropriate pipelines
for execution. REPL commands will be sent to commandValidate()
while turtl code will be sent to interpret(). This function will
return an error if input is malformed.
*/
func eval(input string, ctx *ShellContext) error {
	// Handle repl commands, which start with an @
	if strings.HasPrefix(input, "@") {
		err := commandValidate(input, ctx)
		if err != nil {
			return err
		}
	} else {
		// Handle turtl code
		interpret(input, ctx)
	}
	return nil
}

/*
commandValidate makes sure that interpreter command input
is valid. If it is, it'll parse any supplied arguments and
run that command. This validates and executes REPL commands
(e.g. @help, @quit, etc...), not turtl code itself.
*/
func commandValidate(input string, ctx *ShellContext) error {
	res := strings.Split(input, " ")
	if len(res) > 0 {
		if f, ok := cmdMap[res[0]]; ok {
			err := f(res[1:], ctx)
			if err != nil {
				printShellError(res[0], err)
			}
		} else {
			invalidCommand(res[0])
		}
	}
	return nil
}

/*
CommandMap is a map of shell commands to functions which
execute that command
*/
type CommandMap map[string]shellCommand

/*
cmdMap is the map of console command names to functions
which execute them
*/
var cmdMap = CommandMap{
	"@quit":  cmdQuit,
	"@help":  cmdHelp,
	"@load":  cmdLoad,
	"@run":   cmdRun,
	"@save":  cmdSave,
	"@reset": cmdReset,
	"@name":  cmdName,
}

/*
invalidCommand reports that a supplied command is
not valid
*/
func invalidCommand(cmd string) {
	fmt.Printf("\"%s\" is not a valid command\n", cmd)
}

/*
printShellError is the generic error printing function
for the shell. Used in cases where eval() returns an err.
*/
func printShellError(cmd string, err error) {
	fmt.Println(fmt.Sprintf("there was an error running %s: %s",
		cmd,
		err.Error(),
	))
}

/*
interpret passes input code through the compile/eval pipeline.
It will return an error if compilation fails.
*/
func interpret(input string, ctx *ShellContext) error {
	return nil
}
