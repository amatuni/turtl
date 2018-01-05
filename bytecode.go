package turtl

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var (
	ErrProgramHeaderMalformed = errors.New("[bytecode]: program header is malformed")
)

/*
newBinaryHeader generates the binary header chunk that
prefixes all compiled turtl code. It takes parameter
designating ProgramType (either Execute or Library)
*/
func newBinaryHeader(ptype ProgramType) []byte {
	header := NewProgramHeader(ptype)
	return header.Encode()
}

func encodeProgHeader(header ProgramHeader) []byte {
	buff := bytes.NewBuffer(make([]byte, 0, headerLength))
	binary.Write(buff, binary.LittleEndian, header.sign)
	binary.Write(buff, binary.LittleEndian, header.major)
	binary.Write(buff, binary.LittleEndian, header.minor)
	binary.Write(buff, binary.LittleEndian, header.patch)
	binary.Write(buff, binary.LittleEndian, header.fvers)
	binary.Write(buff, binary.LittleEndian, header.ptype)
	binary.Write(buff, binary.LittleEndian, header.prgid)
	return buff.Bytes()
}

/*
readBinaryHeader parses out the first few bytes of a
Program and returns the ProgramHeader
*/
func readBinaryHeader(p Program) (ProgramHeader, error) {
	header := ProgramHeader{}
	if len(p) < headerLength {
		return header, ErrProgramHeaderMalformed
	}
	copy(
		header.sign[:signLength],
		p[:signLength],
	)
	header.major = p[4]
	header.minor = p[5]
	header.patch = p[6]
	header.fvers = p[7]
	header.ptype = ProgramType(p[8])
	copy(
		header.prgid[:progIDLength],
		p[headerLength-progIDLength:headerLength],
	)

	return header, nil
}
