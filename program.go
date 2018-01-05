package turtl

import (
	"golang.org/x/crypto/sha3"
)

const (
	progIDLength = 8  // program ID length in bytes
	signLength   = 4  // signature length in bytes
	headerLength = 17 // header length in bytes
)

/*
Program is a slice of executable bytecode
*/
type Program []byte

/*
setProgID will hash all of the data that comes after
the header and place the first ${progIDLength}
bytes of that hash at the end of the header.
*/
func (p Program) setProgID() error {
	h := sha3.New256()
	h.Write(p[headerLength:])
	copy(p[headerLength-progIDLength:headerLength],
		h.Sum(nil)[:progIDLength])
	return nil
}

/*
ProgramType defines the type of data held in a slice of
bytecode
*/
type ProgramType uint8

const (
	Execute ProgramType = 0 // contains main() function
	Library ProgramType = 1 // contains only definitions
)

/*
ProgramContainer stores metadata about a Program as
well as a pointer to the raw bytecode
*/
type ProgramContainer struct {
	name         string            // name of the Program
	funcPtrTable map[uint64]uint64 // function pointer table
	prog         *Program          // pointer to the Program
}

/*
ReadHeader bytes from a Program and return the info
as ProgramHeader
*/
func (p *Program) ReadHeader() (ProgramHeader, error) {
	header, err := readBinaryHeader(*p)
	if err != nil {
		return header, err
	}
	return header, nil
}

/*
ProgramHeader contains all the program metadata that's
parsed from the first few bytes of a compiled binary
file. The size of each field here is not arbitrary. Do
not change without making sure alignment updates are
reflected through all other code that will have to deal
with this elsewhere.

The program ID element is a truncated cryptographic hash
of all the data located after the header. Currently it's
the first ${progIDLength} bytes of the SHA3-256 sum this
data.

*/
type ProgramHeader struct {
	sign  [signLength]byte   // header signature
	major byte               // major version
	minor byte               // minor version
	patch byte               // patch version
	fvers byte               // bytecode format version
	ptype ProgramType        // program type
	prgid [progIDLength]byte // program ID
}

/*
NewProgramHeader builds a default configured ProgramHeader.
This is the struct that will be serialized into the first
couple bytes of every compiled .turtlc file.
*/
func NewProgramHeader(ptype ProgramType) ProgramHeader {
	return ProgramHeader{
		sign:  [signLength]byte{1, 2, 3, 4},
		major: major,
		minor: minor,
		patch: patch,
		fvers: 0,
		ptype: ptype,
		prgid: [progIDLength]byte{},
	}
}

/*
Encode the ProgramHeader into a binary representation
*/
func (ph ProgramHeader) Encode() []byte {
	return encodeProgHeader(ph)
}

/*
Signature of the program header
*/
func (ph *ProgramHeader) Signature() [signLength]byte {
	return ph.sign
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

// FormatVersion of bytecode representation
func (ph *ProgramHeader) FormatVersion() uint8 {
	return ph.fvers
}

// ProgramType where executable = 0, and library = 1
func (ph *ProgramHeader) ProgramType() ProgramType {
	return ph.ptype
}

/*
ProgramID is the SHA3-256 hash of all the data coming
after the header
*/
func (ph *ProgramHeader) ProgramID() [progIDLength]byte {
	return ph.prgid
}
