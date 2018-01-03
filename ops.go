package turtl

type Opcode uint8

const (
	HALT Opcode = iota
	ADD
	SUB
	MUL
	DIV
	LOAD
	LOADI
	MOV
	INC
	DEC
	CMP
	AND
	OR
	XOR
	NOT
	JMP
	JEQ
	JNE
	JGT
	JGE
	JLT
	JLE
)

func (vm *VM) mov() {
	vm.r2 = vm.r1
}

func (vm *VM) add() {
	vm.r1 = vm.r2 + vm.r3
}

func (vm *VM) sub() {
	vm.r1 = vm.r2 - vm.r3
}
