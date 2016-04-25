package cpu430

import "fmt"

//"errors"

/*
15 	14 	13 	12 	11 	10 	9 	8 	7 	6 	5 	4 	3 	2 	1 	0 	Instrucción

0 	0 	0 	1 	0 	0 	opcode 	    B/W  As 	   register 	Single-operand arithmetic

0 	0 	0 	1 	0 	0 	0 	0 	0 	B/W  As 	   register 	RRC Rotate right through carry
0 	0 	0 	1 	0 	0 	0 	0 	1 	0 	 As 	   register 	SWPB Swap bytes
0 	0 	0 	1 	0 	0 	0 	1 	0 	B/W  As 	   register 	RRA Rotate right arithmetic
0 	0 	0 	1 	0 	0 	0 	1 	1 	0 	 As 	   register 	SXT Sign extend byte to word
0 	0 	0 	1 	0 	0 	1 	0 	0 	B/W  As 	   register 	PUSH Push value onto stack
0 	0 	0 	1 	0 	0 	1 	0 	1 	0 	 As 	   register 	CALL Subroutine call; push PC and move source to PC
0 	0 	0 	1 	0 	0 	1 	1 	0 	0 	0 	0 	0 	0 	0 	0 	RETI Return from interrupt; pop SR then pop PC
*/

const (
	one    = 0xfc00
	opcode = 0x0380
	bw     = 0x0040
	as     = 0x0030
	reg    = 0x000f
)

/*
0 	0 	1 	condition 	       10-bit signed offset 	        Conditional jump; PC = PC + 2×offset

0 	0 	1 	0 	0 	0 	       10-bit signed offset 	        JNE/JNZ Jump if not equal/zero
0 	0 	1 	0 	0 	1 	       10-bit signed offset 	        JEQ/JZ Jump if equal/zero
0 	0 	1 	0 	1 	0 	       10-bit signed offset 	        JNC/JLO Jump if no carry/lower
0 	0 	1 	0 	1 	1 	       10-bit signed offset 	        JC/JHS Jump if carry/higher or same
0 	0 	1 	1 	0 	0 	       10-bit signed offset 	        JN Jump if negative
0 	0 	1 	1 	0 	1 	       10-bit signed offset 	        JGE Jump if greater or equal
0 	0 	1 	1 	1 	0 	       10-bit signed offset 	        JL Jump if less
0 	0 	1 	1 	1 	1 	       10-bit signed offset  	        JMP Jump (unconditionally)
*/

const (
	jmpc = 0xe000
	cond = 0x1c00
	offs = 0x03ff
)

/*
     opcode 	source 	Ad 	B/W 	As 	destination 	Two-operand arithmetic

0 	1 	0 	0 	source 	Ad 	B/W 	As 	destination 	MOV Move source to destination
0 	1 	0 	1 	source 	Ad 	B/W 	As 	destination 	ADD Add source to destination
0 	1 	1 	0 	source 	Ad 	B/W 	As 	destination 	ADDC Add source and carry to destination
0 	1 	1 	1 	source 	Ad 	B/W 	As 	destination 	SUBC Subtract source from destination (with carry)
1 	0 	0 	0 	source 	Ad 	B/W 	As 	destination 	SUB Subtract source from destination
1 	0 	0 	1 	source 	Ad 	B/W 	As 	destination 	CMP Compare (pretend to subtract) source from destination
1 	0 	1 	0 	source 	Ad 	B/W 	As 	destination 	DADD Decimal add source to destination (with carry)
1 	0 	1 	1 	source 	Ad 	B/W 	As 	destination 	BIT Test bits of source AND destination
1 	1 	0 	0 	source 	Ad 	B/W 	As 	destination 	BIC Bit clear (dest &= ~src)
1 	1 	0 	1 	source 	Ad 	B/W 	As 	destination 	BIS Bit set (logical OR)
1 	1 	1 	0 	source 	Ad 	B/W 	As 	destination 	XOR Exclusive or source with destination
1 	1 	1 	1 	source 	Ad 	B/W 	As 	destination 	AND Logical AND source with destination (dest &= src)
*/

const (
	twoop  = 0xf000
	source = 0x0f00
	ad     = 0x0080
	// Next parameters defined in single operation
	//bw     = 0x0040
	//as     = 0x0030
	dest = 0x000f
)

/*
B/W: Most instructions are available in .B (8-bit byte) and .W (16-bit word) suffixed versions, depending on the value of a B/W bit: the bit is set to 1 for 8-bit and 0 for 16-bit.

MSP430 addressing modes: Ad (Address destination), As (Address source)
As 	Ad 	Register 	Syntax 	Description
00 	0 	n 	Rn 	Register direct. The operand is the contents of Rn.
01 	1 	n 	x(Rn) 	Indexed. The operand is in memory at address Rn+x.
10 	— 	n 	@Rn 	Register indirect. The operand is in memory at the address held in Rn.
11 	— 	n 	@Rn+ 	Indirect autoincrement. As above, then the register is incremented by 1 or 2.

Addressing modes using R0 (PC)
01 	1 	0 (PC) 	ADDR 	Symbolic. Equivalent to x(PC). The operand is in memory at address PC+x.
11 	— 	0 (PC) 	#x 	Immediate. Equivalent to @PC+. The operand is the next word in the instruction stream.

Addressing modes using R2 (SR) and R3 (CG), special-case decoding
01 	1 	2 (SR) 	&ADDR 	Absolute. The operand is in memory at address x.
10 	— 	2 (SR) 	#4 	Constant. The operand is the constant 4.
11 	— 	2 (SR) 	#8 	Constant. The operand is the constant 8.
00 	— 	3 (CG) 	#0 	Constant. The operand is the constant 0.
01 	— 	3 (CG) 	#1 	Constant. The operand is the constant 1. There is no index word.
10 	— 	3 (CG) 	#2 	Constant. The operand is the constant 2.
11 	— 	3 (CG) 	#−1 	Constant. The operand is the constant −1.
*/

// Represents a instruction
//type Opcode interface {
//  Opcode(code uint16) (i instruction) // Return the instruction and their values
//  Execute(i, instruction) (error, cycles)
// TODO:
//}

// Instruction data
/*
type Instruction struct {
	op     int    // 1 Operand, 0 Jmp, 2 Operand
	opcode uint16 //  Operation
	bw     bool   // True -> Byte, False -> Word
	offset uint16 // 10-bit signed offset (jmp Instruction)
	as     uint16 // Address source
	ad     uint16 // Address destination
	reg    uint16 // register

}
*/

// Opcode return the type of instruction
func Opcode(code uint16) string {
	var i uint16

	i = 0xe000
	i &= code
	i >>= 13
	switch i {
	case 0:
		return "Single-operand arithmetic:" + single(code)
	case 1:
		return "Conditional jump; PC = PC + 2×offset:" + jmp(code)
	default:
		return "Two-operand arithmetic:" + two(code)
	}
}

func single(code uint16) string {
	var i, oc, bw, as, r uint16
	var s string

	// Check all bits of single instruction
	i = 0x1c00
	i &= code
	i >>= 10
	if i != 4 {
		return "Error decoding instruction, single"
	}

	// Check opcode
	oc = 0x0380
	oc &= code
	oc >>= 6
	switch oc {
	case 0:
		s = "RRC Rotate right through carry"
	case 1:
		s = "SWPB Swap bytes"
	case 2:
		s = "RRA Rotate right arithmetic"
	case 3:
		s = "SXT Sign extend byte to word"
	case 4:
		s = "PUSH Push value onto stack"
	case 5:
		s = "CALL Subroutine call; push PC and move source to PC"
	case 6:
		s = "ETI Return from interrupt; pop SR then pop PC"
	default:
		return "Error decoding instruction, opcode"
	}

	// Check B/W
	bw = 0x0040
	bw &= code
	bw >>= 5
	if bw == 0 {
		// Ok in all opcodes
		s += ".W"
	} else {
		// Ok in opcodes=0,2 or 4
		if (oc == 0) || (oc == 2) || (oc == 4) {
			s += ".B"
		} else {
			return "Error decoding, B/W"
		}
	}

	// Check as Address Source
	as = 0x0030
	as &= code
	as >>= 4
	s += fmt.Sprintf(",as: %d", as)

	// Check r Register
	r = 0x000f
	r &= code
	s += fmt.Sprintf(",reg: %d", r)

	return s
}

func jmp(code uint16) string {
	var c, os uint16
	var s string

	// Check Condition
	c = 0x1c00
	c &= code
	c >>= 10
	s += fmt.Sprintf("condition: %d", c)

	// Check Offset
	os = 0x03ff
	os &= code
	s += fmt.Sprintf(",offset: %x", os)

	return s
}

func two(code uint16) string {
	var oc, s1, ad, bw, as, s2 uint16
	var s string

	// Check all bits of single instruction
	oc = 0xf000
	oc &= code
	oc >>= 12
	if oc < 4 {
		return "Error decoding instruction, two"
	}

	// Check opcode
	s1 = 0x0f00
	s1 &= code
	s1 >>= 8
	s += fmt.Sprintf(",s1: %x", s1)

	// Check ad Address Destination
	ad = 0x0080
	ad &= code
	ad >>= 7
	s += fmt.Sprintf(",ad: %d", ad)

	// Check B/W
	bw = 0x0040
	bw &= code
	bw >>= 5
	if bw == 0 {
		// Ok in all opcodes
		s += ".W"
	} else {
		s += ".B"
	}

	// Check as Address Source
	as = 0x0030
	as &= code
	as >>= 4
	s += fmt.Sprintf(",as: %d", as)

	// Check Source2
	s2 = 0x000f
	s2 &= code
	s += fmt.Sprintf(",s2: %x", s2)

	return s
}
