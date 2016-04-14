package cpu430

import (
	"testing"
)

func TestCpu(t *testing.T) {
	r := NewRegisters()
	r.Print()
	r.R[PC] = 15
	r.R[SP] = 255
	r.R[15] = 65535
	r.Print()
}
