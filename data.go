package turtl

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// Endings of different file types
const (
	sourceSuffix  = ".turtl"  // source files
	binarySuffix  = ".turtlc" // compiled code
	sessionSuffix = ".turtls" // shell session files
)

/*
LoadBinary reads in compiled turtl code and returns a Program
*/
func LoadBinary(fpath string) (Program, error) {
	if !strings.HasSuffix(fpath, binarySuffix) {
		return nil, fmt.Errorf("[data]: %s is not a valid turtlc file",
			path.Base(fpath))
	}
	f, err := os.Open(fpath)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	finfo, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fsize := finfo.Size()
	buffer := make([]byte, fsize)
	_, err = f.Read(buffer)
	if err != nil {
		return nil, err
	}

	err = validateBinaryFile(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

/*
saveBinary to a path
*/
func saveBinary(p Program, path string) error {
	if !strings.HasSuffix(path, binarySuffix) {
		path = path + ".turtlc"
	}
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(p)
	return nil
}

func validateBinaryFile(data []byte) error {
	return nil
}

type binaryReadErr struct {
	f string
}

func (err binaryReadErr) Error() string {
	return fmt.Sprintf("[data]: %s is not a valid turtlc file", err.f)
}
