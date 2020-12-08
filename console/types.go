package console

type OpCode int
const (
	ACC = iota
	JMP
	NOP
)

type Console struct {
	// accumulator value
	accumulator int64
	// Instruction position
	pos int64
	// Instructions
	instructions []Instruction
}

type Instruction struct {
	op OpCode
	arg int64
}

type Program []Instruction
