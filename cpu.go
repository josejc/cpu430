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

// Registers 16 of 16bit
type Registers struct {
	R [16]uint16
}

// NewRegisters Creates a new set
func NewRegisters() (reg *Registers) {
	reg = new(Registers)
	reg.Reset()
	return
}

// Reset all registers
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

// Print the values of each register
func (reg *Registers) Print() {
	fmt.Printf(reg.String())
}

// String return a string with the values of registers
func (reg *Registers) String() string {
	l1 := fmt.Sprintf("R[0]:%#x, R[1]:%#x, R[2]:%#x, R[3]:%#x\n", reg.R[0], reg.R[1], reg.R[2], reg.R[3])
	l2 := fmt.Sprintf("R[4]:%#x, R[5]:%#x, R[6]:%#x, R[7]:%#x\n", reg.R[4], reg.R[5], reg.R[6], reg.R[7])
	l3 := fmt.Sprintf("R[8]:%#x, R[9]:%#x, R[10]:%#x, R[11]:%#x\n", reg.R[8], reg.R[9], reg.R[10], reg.R[11])
	l4 := fmt.Sprintf("R[12]:%#x, R[13]:%#x, R[14]:%#x, R[15]:%#x\n", reg.R[12], reg.R[13], reg.R[14], reg.R[15])
	return l1 + l2 + l3 + l4
}

// Dump return a strings with the values
func (reg *Registers) Dump() []string {
	dump := make([]string, 4)
	dump[0] = fmt.Sprintf("R[00]:%#x, R[01]:%#x, R[02]:%#x, R[03]:%#x\n", reg.R[0], reg.R[1], reg.R[2], reg.R[3])
	dump[1] = fmt.Sprintf("R[04]:%#x, R[05]:%#x, R[06]:%#x, R[07]:%#x\n", reg.R[4], reg.R[5], reg.R[6], reg.R[7])
	dump[2] = fmt.Sprintf("R[08]:%#x, R[09]:%#x, R[10]:%#x, R[11]:%#x\n", reg.R[8], reg.R[9], reg.R[10], reg.R[11])
	dump[3] = fmt.Sprintf("R[12]:%#x, R[13]:%#x, R[14]:%#x, R[15]:%#x\n", reg.R[12], reg.R[13], reg.R[14], reg.R[15])
	return dump
}

// CPU Represents the msp430 cpu
type CPU struct {
	reg  Registers
	inst Instruction
	// TODO: interrupts?
}

/*

// SoC: System on a Chip (Computer)
type soc struct {
	cpu CPU
	mem Memory
	// channels for bus implementation
	busaddr, busdata, busctrl channels between memory and cpu
	//	clock Clock
}

// Without channels
func (cpu *CPU) Execute() {
	// fetch
	inst := cpu.i.Decode(soc.mem,cpu.R[PC])
	if inst != invalid {
		fmt.Printf("No such opcode 0x%x\n", opcode)
		os.Exit(1)
	}
	// execute, exec() returns number of cycles
	cycles := inst.execute(soc.cpu, soc.mem)
	// count cycles
	for _ = range cpu.clock.ticker.C {
		cycles--
		if cycles == 0 {
			break
		}
	}
}

// With channels
func (cpu *CPU) Execute() {
	// fetch
	busaddr <- soc.cpu.reg.R[PC]
	busctrl <- ReadW
	inst <- busdata

	inst.Opcode(reg.R[PC])
	if inst.kind == invalid {
		fmt.Printf("No such opcode 0x%x\n", opcode)
		os.Exit(1)
	}

	// execute, exec() returns number of cycles
	// update the registers and execute the insttruction
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
