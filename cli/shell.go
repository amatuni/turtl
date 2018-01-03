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

	"github.com/andreiamatuni/turtl"
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

	fmt.Println(turtl.InfoGraphic())
	fmt.Println("Welcome to turtl. Type @help to see available commands.")

	ctx := turtl.NewShellContext()

	r := rune(0)
	for {
		prompt := fmt.Sprintf("%s%d%s %s :> ", red("["), ctx.CurrentPos, red("]"), green("turtl"))
		fmt.Fprintf(os.Stdout, "\r\033[K%s%s \033[0J\033[%dD", prompt, ctx.Line, len(ctx.Line)-ctx.Index+1)
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
					ctx.CurrentPos++
					if len(ctx.History) < ctx.CurrentPos {
						ctx.CurrentPos = len(ctx.History)
					}
					if len(ctx.History) == ctx.CurrentPos {
						ctx.Line = ""
					} else {
						ctx.Line = ctx.History[ctx.CurrentPos]
					}
					ctx.Index = len(ctx.Line)
				case 65: // Up
					ctx.CurrentPos--
					if ctx.CurrentPos < 0 {
						ctx.CurrentPos = 0
					}
					if len(ctx.History) > 0 {
						ctx.Line = ctx.History[ctx.CurrentPos]
					}
					ctx.Index = len(ctx.Line)
				case 67: // Right
					ctx.Index++
					if ctx.Index > len(ctx.Line) {
						ctx.Index = len(ctx.Line)
					}
				case 68: // Left
					ctx.Index--
					if ctx.Index < 0 {
						ctx.Index = 0
					}
				}
				continue
			} else if bPrev == 51 && r == 126 { // DELETE
				if len(ctx.Line) > 0 && ctx.Index < len(ctx.Line) {
					ctx.Line = ctx.Line[:ctx.Index] + ctx.Line[ctx.Index+1:]
				}
				if ctx.Index > len(ctx.Line) {
					ctx.Index = len(ctx.Line)
				}
				continue
			}
			ctx.Line = ctx.Line[:ctx.Index] + string(r) + ctx.Line[ctx.Index:]
			ctx.Index++
		case 127, '\b': // Backspace
			if len(ctx.Line) > 0 && ctx.Index > 0 {
				ctx.Line = ctx.Line[:ctx.Index-1] + ctx.Line[ctx.Index:]
				ctx.Index--
			}
			if ctx.Index > len(ctx.Line) {
				ctx.Index = len(ctx.Line)
			}
		case 27:
		case 10, 13: // Enter
			fmt.Fprintln(os.Stdout, "\033[100000C\033[0J")
			ctx.Count++
			ctx.CurrentPos = ctx.Count
			ctx.History = append(ctx.History, ctx.Line)
			eval(ctx.Line, &ctx)
			ctx.Line = ""
			ctx.Index = 0
		case 4: // Ctrl-D
			fmt.Println()
			return nil
		}
	}
}

/*
eval passes shell input through the appropriate pipelines
for execution. REPL commands will be sent to commandValidate()
while turtl code will be sent to interpret(). This function will
return an error if input is malformed.
*/
func eval(input string, ctx *turtl.ShellContext) error {
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
func commandValidate(input string, ctx *turtl.ShellContext) error {
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
type CommandMap map[string]turtl.ShellCommand

/*
cmdMap is the map of console command names to functions
which execute them
*/
var cmdMap = CommandMap{
	"@quit":  turtl.CmdQuit,
	"@help":  turtl.CmdHelp,
	"@load":  turtl.CmdLoad,
	"@run":   turtl.CmdRun,
	"@save":  turtl.CmdSave,
	"@reset": turtl.CmdReset,
	"@name":  turtl.CmdName,
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
func interpret(input string, ctx *turtl.ShellContext) error {
	return nil
}
