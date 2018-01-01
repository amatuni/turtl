package main

/*
VM is the main virtual machine struct for 64 bit platforms.

We avoid placing pointers in the VM struct so that it can be
allocated on the stack by the Go runtime. Access to turtl
runtime resources like heaps and bytecode are accessed
through 8 bit integer ID's that serve as pointers into global
maps which store these resources.

Support for 32 bit platforms is provided by the VM32 type.
*/
type VM struct {
	ip     uint  // instruction pointer
	fp     uint  // frame pointer
	sp     uint  // stack pointer
	r1     int64 // register 1
	r2     int64 // register 2
	r3     int64 // register 3
	progID uint8 // ID of program associated to this VM
	heapID uint8 // ID of heap associated to this VM
}

/*
NewVM construct a new virtual machine instance
*/
func NewVM(progID uint8) *VM {
	return &VM{
		progID: progID,
	}
}

func (vm *VM) loadProgram(path string) error {
	return nil
}

/*
tick runs the VM forward one clock cycle
*/
func (vm *VM) tick() {

}

/*
VM32 is the main virtual machine struct for 32 bit platforms
*/
type VM32 struct {
	ip     uint  // instruction pointer
	fp     uint  // frame pointer
	sp     uint  // stack pointer
	r1     int32 // register 1
	r2     int32 // register 2
	r3     int32 // register 3
	progID uint8 // ID of program associated to this VM
	heapID uint8 // ID of heap associated to this VM
}

func NewVM32() *VM {
	return &VM{}
}
