package cpu430

import (
	"fmt"
	"testing"
)

func TestCpu(t *testing.T) {
	r := NewRegisters()
	r.Print()
	r.R[0] = 10
	r.R[1] = 11
	*r.PC = uint16(255)
	*r.SP = 65535
	r.R[15] = 65535
	r.Print()
	fmt.Println(*r.PC, *r.SP)
}
