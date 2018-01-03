package turtl

import (
	"bytes"
	"encoding/binary"
)

const (
	headerLength = 10
)

/*
NewBinaryHeader generates the binary header chunk that
prefixes all compiled turtl code
*/
func NewBinaryHeader() []byte {
	/*
		TODO: preallocate exact size of the buffer before
		beginning to fill it
	*/
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, major)
	binary.Write(buff, binary.LittleEndian, minor)
	binary.Write(buff, binary.LittleEndian, patch)

	return buff.Bytes()
}

/*
ReadBinaryHeader parses out the first few bytes of a
Program and returns the ProgramHeader
*/
func ReadBinaryHeader(p Program) ProgramHeader {
	header := ProgramHeader{}
	buff := bytes.NewBuffer(p[:headerLength])
	binary.Read(buff, binary.LittleEndian, header.major)
	return header
}
