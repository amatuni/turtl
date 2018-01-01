package main

/*
Program is a slice of executable bytecode
*/
type Program []uint

/*
ProgramMetadata stores metadata about a Program as
well as a pointer to the raw bytecode
*/
type ProgramMetadata struct {
	name         string          // name of the Program
	funcPtrTable map[uint]string // function pointer table
	prog         *Program        // pointer to the Program
}
