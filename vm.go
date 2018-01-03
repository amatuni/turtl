package turtl

/*
VM is the main virtual machine struct for 64 bit platforms.

We avoid placing pointers in the VM struct so that it can be
allocated on the stack by the Go runtime. Access to turtl
runtime resources like heaps and bytecode are accessed
through 8 bit integer ID's that serve as pointers into global
maps which store these resources.

*/
type VM struct {
	ip     uint  // instruction pointer
	fp     uint  // frame pointer
	sp     uint  // stack pointer
	r1     int64 // register 1
	r2     int64 // register 2
	r3     int64 // register 3
	currOp uint8 // currently loaded opcode
	progID uint8 // ID of program associated to this VM
	heapID uint8 // ID of heap associated to this VM
}

/*
NewVM construct a new virtual machine instance
*/
func NewVM() *VM {
	return &VM{}
}

/*
Run the VM, at the end return resulting value as
pointer to turtl Object
*/
func (vm *VM) Run() (*Object, error) {
	return &Object{}, nil
}

/*
LoadProgram into the virtual machine
*/
func (vm *VM) LoadProgram(prog []byte) error {
	return nil
}

/*
LoadProgramFile into the virtual machine
*/
func (vm *VM) LoadProgramFile(path string) error {
	return nil
}

/*
tick runs the VM forward one clock cycle
*/
func (vm *VM) tick() {
	switch vm {

	}
}

/*
IP returns the instruction pointer
*/
func (vm *VM) IP() uint {
	return vm.ip
}

/*
FP returns the frame pointer
*/
func (vm *VM) FP() uint {
	return vm.fp
}

/*
SP returns the stack pointer
*/
func (vm *VM) SP() uint {
	return vm.sp
}
