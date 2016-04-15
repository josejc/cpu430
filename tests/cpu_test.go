package cpu430_test

import (
	"github.com/josejc/cpu430"
	"testing"
)

func TestCpu(t *testing.T) {
	r := cpu430.NewRegisters()
	r.Print()
	r.R[cpu430.PC] = 15
	r.R[cpu430.SP] = 255
	r.R[15] = 65535
	r.Print()
}
