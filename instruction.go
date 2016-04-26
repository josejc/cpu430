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

0 	0 	1 	condition 	       10-bit signed offset 	        Conditional jump; PC = PC + 2×offset

0 	0 	1 	0 	0 	0 	       10-bit signed offset 	        JNE/JNZ Jump if not equal/zero
0 	0 	1 	0 	0 	1 	       10-bit signed offset 	        JEQ/JZ Jump if equal/zero
0 	0 	1 	0 	1 	0 	       10-bit signed offset 	        JNC/JLO Jump if no carry/lower
0 	0 	1 	0 	1 	1 	       10-bit signed offset 	        JC/JHS Jump if carry/higher or same
0 	0 	1 	1 	0 	0 	       10-bit signed offset 	        JN Jump if negative
0 	0 	1 	1 	0 	1 	       10-bit signed offset 	        JGE Jump if greater or equal
0 	0 	1 	1 	1 	0 	       10-bit signed offset 	        JL Jump if less
0 	0 	1 	1 	1 	1 	       10-bit signed offset  	        JMP Jump (unconditionally)

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

// Mask for all parameters of instructions
const (
	kind  = 0xe000
	oneOC = 0x0380
	twoOC = 0xf000
	AD    = 0x0080
	BW    = 0x0040
	AS    = 0x0030
	SRC   = 0x0f00
	DST   = 0x000f
	COND  = 0x1c00
	OFFS  = 0x03ff
)

func mask(code, m uint16) uint16 {
	return ((code & m) >> ffs(m))
}

// find first set for obtain the shift
func ffs(m uint16) uint16 {
	var i uint16

	for i = 0; i < 16; i++ {
		if (m & 1) == 1 {
			return i
		}
		m >>= 1
	}
	return i
}

// Opcode return the type of instruction
func Opcode(code uint16) string {

	switch mask(code, kind) {
	case 0:
		return "Single-operand arithmetic:" + single(code)
	case 1:
		return "Conditional jump; PC = PC + 2×offset:" + jmp(code)
	default:
		return "Two-operand arithmetic:" + two(code)
	}
}

func single(code uint16) string {
	var oc, bw, as, r uint16
	var s string

	// Check all bits of single instruction
	if mask(code, COND) != 4 {
		return "Error decoding instruction, single"
	}

	// Check opcode
	oc = mask(code, oneOC)
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
	bw = mask(code, BW)
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
	as = mask(code, AS)
	s += fmt.Sprintf(",as: %d", as)

	// Check r Register
	r = mask(code, DST)
	s += fmt.Sprintf(",reg: %d", r)

	return s
}

func jmp(code uint16) string {
	var c, os uint16
	var s string

	// Check Condition
	c = mask(code, COND)
	s += fmt.Sprintf("condition: %d", c)

	// Check Offset
	os = mask(code, OFFS)
	s += fmt.Sprintf(",offset: %x", os)

	return s
}

func two(code uint16) string {
	var oc, s1, ad, bw, as, s2 uint16
	var s string

	// Check all bits of single instruction
	oc = mask(code, twoOC)
	if oc < 4 {
		return "Error decoding instruction, two"
	}

	// Check opcode
	s1 = mask(code, SRC)
	s += fmt.Sprintf(",s1: %x", s1)

	// Check ad Address Destination
	ad = mask(code, AD)
	s += fmt.Sprintf(",ad: %d", ad)

	// Check B/W
	bw = mask(code, BW)
	if bw == 0 {
		// Ok in all opcodes
		s += ".W"
	} else {
		s += ".B"
	}

	// Check as Address Source
	as = mask(code, AS)
	s += fmt.Sprintf(",as: %d", as)

	// Check Source2
	s2 = mask(code, DST)
	s += fmt.Sprintf(",s2: %x", s2)

	return s
}
