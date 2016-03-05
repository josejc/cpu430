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

const (
	PC = iota
	SP
	SR
	CG
)

type Registers struct {
	R [16]uint16
}

// Creates a new set of Registers
func NewRegisters() (reg *Registers) {
	reg = new(Registers)
	reg.Reset()
	return
}

// Resets all registers
func (reg *Registers) Reset() {
	for i := 4; i < 16; i++ {
		reg.R[i] = 0
	}
	reg.R[PC] = 0
	reg.R[SP] = 0
	reg.R[SR] = 0
	reg.R[CG] = 0
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
type CPU struct {
	reg      Registers
	src, dst uint16
	inst     uint16
	//instructions InstructionTable
}

/*

// SoC: System on a Chip (Computer)
type soc struct {
	cpu CPU
	mem Memory
	//	clock Clock
}

func (cpu *CPU) Execute() {
	// fetch
	opcode := OpCode(soc.memory.read(soc.cpu.reg.R[PC]))
	inst, ok := cpu.instructions[opcode]

	if !ok {
		fmt.Printf("No such opcode 0x%x\n", opcode)
		os.Exit(1)
	}

	// execute, exec() returns number of cycles
	cycles := inst.exec(cpu)

	// count cycles
	for _ = range cpu.clock.ticker.C {
		cycles--

		if cycles == 0 {
			break
		}
	}
}
*/
