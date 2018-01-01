package main

import (
	"fmt"
	"unsafe"

	"github.com/spf13/cobra"
)

const (
	MAJOR = 0
	MINOR = 1
	PATCH = 0
)

var versionNumStr = fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, PATCH)
var intSize = unsafe.Sizeof(uint(0)) * 8 // number of bits in an uint

// GitHash is the git HEAD commit hash the binary
// was built with
var GitHash string

func version() string {
	return fmt.Sprintf("%s\n%s", versionNumStr, GitHash[:12])
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nturtl v" + version() + "\n\n")
	},
}

func infoGraphic() string {
	return fmt.Sprintf(
		`

    .=.  ____  .=.
    \ .-'    '-. /       turtl  v%s  - %d bit
    /.'\_/\_/'.\.-p.     %s
--=|: -<_><_>- :|   >
    \'./ \/ \.'/'-b'
    / '-.____.-' \
    '='        '='

Welcome to turtl. Type @help to see available commands.
`, versionNumStr, intSize, GitHash[:13])

}
