package main

import (
	"fmt"
)

const (
	C      uint16 = 1 << iota // carry
	Z                         // zero
	N                         // negative
	GIE                       // general interrupt enable
	CPUOFF                    // cpu off
	OSCOFF                    // oscillator off
	SCG0                      // system clock generator 0
	SCG1                      // system clock generator 1
	V                         // overfloww
)

const (
	_ = iota
	PC
	SP
	SR
	CG1 = 3
	CG2 = 4
)

type Registers struct {
	R  [4]uint16
	PC *uint16
}

// Creates a new set of Registers
func NewRegisters() (reg *Registers) {
	reg = new(Registers)
	reg.Reset()
	return
}

// Resets all registers
func (reg *Registers) Reset() {
	for i := 0; i < 3; i++ {
		reg.R[i] = 0
	}
	reg.PC = &reg.R[0]
	fmt.Println("---", reg.PC, &reg.R[0])
}

// Prints the values of each register
func (reg *Registers) Print() {
	fmt.Printf("R[0]:%#x, R[1]:%#x, R[2]:%#x, R[3]:%#x\n", reg.R[0], reg.R[1], reg.R[2], reg.R[3])
}

func main() {
	r := NewRegisters()
	fmt.Println(r.PC, &r.R[0])
	fmt.Println(&r.PC)
	fmt.Println(*r.PC, r.R[0])
	r.Print()
	r.R[0] = 10
	*r.PC = uint16(255)
	r.R[3] = 65535
	r.Print()
	fmt.Println(r.PC, &r.R[0])
	fmt.Println(&r.PC)
	fmt.Println(*r.PC, r.R[0])

	fmt.Println(C, Z, N, GIE, CPUOFF, OSCOFF, SCG0, SCG1, V)
	fmt.Println(PC, SP, SR, CG1, CG2)
}
