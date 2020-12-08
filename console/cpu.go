package console

import "fmt"

func New(program []string) (*Console, error) {
	// build the computer
	instructions := make([]Instruction, 0)
	for _, line := range program {
		i, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, *i)
	}
	c := &Console{
		accumulator:  0,
		pos:          0,
		instructions: instructions,
	}
	return c, nil
}

func (c *Console) Process() error {
	if c.pos < 0 || c.pos >= int64(len(c.instructions)) {
		return fmt.Errorf("invalid instruction pointer %d", c.pos)
	}

	i := c.instructions[c.pos]
	switch i.op {
	case ACC:
		c.accumulator += i.arg
		c.pos += 1
	case JMP:
		c.pos += i.arg
	case NOP:
		c.pos += 1
	default:
		return fmt.Errorf("invalid instruction found: %v", i.op)
	}
	return nil
}

func (c *Console) Run() {

	seen := make(map[int]bool, 0)
	for true {
		if seen[int(c.pos)] || int(c.pos) >= len(c.instructions) {
			break
		}
		seen[int(c.pos)] = true
		c.Process()
	}
}

func (c *Console) HasFinished() bool {
	return int(c.pos) >= len(c.instructions)
}

func (c *Console) FixBrokenCode() {

	origInstructions := make([]Instruction, len(c.instructions))
	copy(origInstructions, c.instructions)

	for i := 0; i < len(c.instructions); i++ {
		replaceInst := c.instructions[i].op
		switch replaceInst {
		case JMP:
			c.instructions[i].op = NOP
		case NOP:
			c.instructions[i].op = JMP
		default:
			continue
		}
		c.Run()
		if c.HasFinished() {
			return
		}
		// reset state of the computer
		c.instructions[i].op = replaceInst
		c.pos = 0
		c.accumulator = 0
	}
}

func (c *Console) Accumulator() int64 {
	return c.accumulator
}
