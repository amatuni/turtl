package turtl

import (
	"bytes"
	"testing"
)

var correctHash = []byte{
	218, 138, 223,
	26, 251, 76,
	51, 4}

func TestSetProgID(t *testing.T) {
	header := newBinaryHeader(Execute)
	prog := Program(append(header, []byte{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9}..., // dummy program data to hash
	))
	prog.setProgID()
	idRegion := prog[headerLength-progIDLength : headerLength]
	if bytes.Compare(
		idRegion,
		correctHash) != 0 {
		t.Errorf("program id %v != %v",
			idRegion,
			correctHash,
		)
	}
}
