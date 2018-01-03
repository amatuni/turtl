package turtl

/*
Program is a slice of executable bytecode
*/
type Program []byte

/*
ProgramMetadata stores metadata about a Program as
well as a pointer to the raw bytecode
*/
type ProgramMetadata struct {
	name         string        // name of the Program
	funcPtrTable map[byte]uint // function pointer table
	prog         *Program      // pointer to the Program
}

/*
ReadHeader bytes from a Program and return the info
as ProgramHeader
*/
func (p *Program) ReadHeader() (ProgramHeader, error) {
	header := ProgramHeader{}
	return header, nil
}

/*
ProgramHeader contains all the program metadata that's
parsed from the first few bytes of a compiled binary
file
*/
type ProgramHeader struct {
	// turtl version
	major uint8
	minor uint8
	patch uint8
}

// Major version
func (ph *ProgramHeader) Major() uint8 {
	return ph.major
}

// Minor version
func (ph *ProgramHeader) Minor() uint8 {
	return ph.minor
}

// Patch version
func (ph *ProgramHeader) Patch() uint8 {
	return ph.patch
}
