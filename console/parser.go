package console

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInstruction(s string) (*Instruction, error) {

	sp := strings.Split(s, " ")
	if len(sp) != 2 {
		return nil, fmt.Errorf("invalid instruction format: %s", s)
	}

	opArg, err := strconv.ParseInt(sp[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid op argument format: %s", sp[1])
	}

	i := &Instruction{
		arg: opArg,
	}
	switch sp[0] {
	case "acc":
		i.op = ACC
	case "nop":
		i.op = NOP
	case "jmp":
		i.op = JMP
	default:
		return nil, fmt.Errorf("invalid op: %s", sp[0])
	}

	return i, nil
}


