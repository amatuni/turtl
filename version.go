package turtl

import (
	"fmt"
	"unsafe"
)

const (
	major uint8 = 0
	minor uint8 = 1
	patch uint8 = 0
)

var versionNumStr = fmt.Sprintf("%d.%d.%d", major, minor, patch)
var intSize = unsafe.Sizeof(uint(0)) * 8 // number of bits in an uint

// gitHash is the git HEAD commit hash the binary
// was built with
var gitHash string

/*
VersionNum returns the string representation of the
version number: Major.Minor.Patch
*/
func VersionNum() string {
	return versionNumStr
}

/*
GitHash returns the git commit hash for the version
of turtl that was compiled
*/
func GitHash() string {
	return gitHash
}

/*
IntSize returns the size of an unsigned int on this
platform
*/
func IntSize() uintptr {
	return intSize
}

/*
InfoGraphic returns a welcome message with some system
info and an ascii turtle.
*/
func InfoGraphic() string {
	return fmt.Sprintf(
		`

    .=.  ____  .=.
    \ .-'    '-. /       turtl  v%s  - %d bit
    /.'\_/\_/'.\.-p.     %s
--=|: -<_><_>- :|   >
    \'./ \/ \.'/'-b'
    / '-.____.-' \
    '='        '='

`, versionNumStr, intSize, gitHash[:13])
}
