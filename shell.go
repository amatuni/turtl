package turtl

/*
ShellContext holds the current shell state
*/
type ShellContext struct {
	ID         uint     // session id
	Name       string   // session name
	VM         *VM      // the VM instance loaded into this shell
	History    []string // the record of all past input
	Line       string   // the current line
	Index      int
	Count      int
	CurrentPos int
}

/*
NewShellContext builds a new shell context
*/
func NewShellContext() ShellContext {
	return ShellContext{
		History:    []string{},
		Index:      0,
		Count:      0,
		CurrentPos: 0,
		Line:       "",
	}
}
