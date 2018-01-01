package main

import "os"

// Endings of different file types
const (
	sourceSuffix  = ".turtl"  // source files
	binarySuffix  = ".turtlc" // compiled code
	sessionSuffix = ".turtls" // shell session files
)

/*
	load reads in compiled turtl code and returns a ProgramMetadata
*/
func load(path string) ([]byte, error) {
	f, err := os.Open(path)
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
	return buffer, nil
}
