package turtl

import (
	"testing"
	"unsafe"
)

func TestNewBinaryHeader(t *testing.T) {
	header := newBinaryHeader(Execute)
	lenX := len(header)
	if lenX != headerLength {
		t.Errorf("length of header: %d != %d", lenX, headerLength)
	}
	sizeOfElem := unsafe.Sizeof(header[0])
	if sizeOfElem != 1 {
		t.Errorf("size of header elements: %d != %d", sizeOfElem, 1)
	}
}

func TestReadBinaryHeader(t *testing.T) {
	header := newBinaryHeader(Execute)
	result, err := readBinaryHeader(header)
	if err != nil {
		t.Error(err)
	}
	if result.Major() != Major() {
		t.Errorf("parsedHeader.Major %d != %d", result.Major(), Major())
	}
	if result.Minor() != Minor() {
		t.Errorf("parsedHeader.Minor %d != %d", result.Minor(), Minor())
	}
	if result.Patch() != Patch() {
		t.Errorf("parsedHeader.Patch %d != %d", result.Patch(), Patch())
	}
	if result.FormatVersion() != 0 {
		t.Errorf("parsedHeader.FormatVersion %d != %d", result.FormatVersion(), 0)
	}
}
