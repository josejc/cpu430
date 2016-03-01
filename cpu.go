// Package cpu430 simulates the CPU of msp430
package cpu430

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

type Registers struct {
	R                    [16]uint16
	PC, SP, SR, CG1, CG2 *uint16
}

// Creates a new set of Registers
func NewRegisters() (reg Registers) {
	reg = Registers{}
	reg.Reset()
	return
}

// Resets all registers
func (reg *Registers) Reset() {
	for i := 0; i < 16; i++ {
		reg.R[i] = 0
	}
	reg.PC = &reg.R[0]
	reg.SP = &reg.R[1]
	reg.SR = &reg.R[2]
	reg.CG1 = &reg.R[2]
	reg.CG2 = &reg.R[3]
	// TODO: PC, SP, SR different values?
}

// Prints the values of each register
func (reg *Registers) Print() {
	fmt.Printf("R[0]:%#x, R[1]:%#x, R[2]:%#x, R[3]:%#x\n", reg.R[0], reg.R[1], reg.R[2], reg.R[3])
	fmt.Printf("R[4]:%#x, R[5]:%#x, R[6]:%#x, R[7]:%#x\n", reg.R[4], reg.R[5], reg.R[6], reg.R[7])
	fmt.Printf("R[8]:%#x, R[9]:%#x, R[10]:%#x, R[11]:%#x\n", reg.R[8], reg.R[9], reg.R[10], reg.R[11])
	fmt.Printf("R[12]:%#x, R[13]:%#x, R[14]:%#x, R[15]:%#x\n", reg.R[12], reg.R[13], reg.R[14], reg.R[15])
}

// Represents the msp430 cpu
type cpu struct {
	//Clock		Clock
	Registers Registers
	//Instructions InstructionTable
}
