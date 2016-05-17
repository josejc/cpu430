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
10 	— 	n 	@Rn 	Register indirect. The operand is in memory at the address held in Rn. // NOT for destination
11 	— 	n 	@Rn+ 	Indirect autoincrement. As above, then the register is incremented by 1 or 2. // NOT for destination

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

// Instruction all the values for disassm and execute a Instruction
type Instruction struct {
	kind, oneoc, twooc, ad, bw, as, src, dst, cond, offs uint16
	asX, adX                                             uint16    // as,ad=01 -> X is stored in the next world and now store in this variable
	l                                                    uint16    // long of instruction, MAX=3
	hex                                                  [3]uint16 // Content in address memory (3 words)
	asm                                                  string    // mnemonic asm instruction and operands
}

// NewInstruction Returns a pointer to a new Instruction with all values initialized to zero.
func NewInstruction(m *Memory, adr uint16) *Instruction {
	i := &Instruction{}
	i.hex[0], _ = m.ReadW(adr)
	i.hex[1], _ = m.ReadW(adr + 2)
	i.hex[2], _ = m.ReadW(adr + 4)
	return i
}

// Mnemonics asm, slice of slice strings ;)
var mnemonic = [][]string{
	{"rrc", "swpb", "rra", "sxt", "push", "call", "reti"},
	{"jnz", "jz", "jnc", "jc", "jn", "jge", "jl", "jmp"},
	{"mov", "add", "addc", "subc", "sub", "cmp", "dadd", "bit", "bic", "bis", "xor", "and"},
}

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

// Dissasm return the string of disassembly instruction
func (i *Instruction) Dissasm() string {
	if i.l == 0 {
		i.Opcode()
	}
	return i.asm
}

// Long return the number of words of instruction
func (i *Instruction) Long() uint16 {
	return i.l
}

// Opcode return the type of instruction
func (i *Instruction) Opcode() {
	i.kind = mask(i.hex[0], kind)
	switch i.kind {
	case 0:
		i.single()
	case 1:
		i.jmp()
	default:
		i.two()
	}
}

func (i *Instruction) single() {
	code := i.hex[0]
	// Check all bits of single instruction
	if mask(code, COND) != 4 {
		i.l = 0 // Long = 0 => Invalid instruction
		i.asm = "Single: Error invalid instruction"
		return
	}

	// Check opcode
	i.oneoc = mask(code, oneOC)
	if i.oneoc > 6 {
		i.l = 0
		i.asm = "Single: Error decoding instruction"
		return
	}
	i.l = 1
	i.asm += mnemonic[0][i.oneoc]

	// Check B/W
	i.bw = mask(code, BW)
	if i.bw == 0 {
		// Ok in all opcodes
		i.asm += ".w "
	} else {
		// Ok in opcodes=0,2 or 4
		if (i.oneoc == 0) || (i.oneoc == 2) || (i.oneoc == 4) {
			i.asm += ".b "
		} else {
			i.l = 0
			i.asm = "Single: Error decoding, B/W"
			return
		}
	}

	// Check as Address Source
	i.as = mask(code, AS)
	// Check r Register
	i.dst = mask(code, DST)

	dst := true
	switch i.dst {
	case 0: // PC
		switch i.as {
		case 1: // Symbolic. Equivalent to x(PC). The operand is in memory at address PC+x.
			// disassm nothing to do, in execution we need PC+X
		case 3: // Immediate. Equivalent to @PC+. The operand is the next word in the instruction stream.
			dst = false
			i.dst = i.hex[i.l]
			i.l++
			i.asm += fmt.Sprintf("#%x", i.dst)

		}
	case 2: // SR
		if i.as != 0 {
			dst = false
			switch i.as {
			case 1: // Absolute. The operand is in memory at address x.
				i.dst = i.hex[i.l]
				i.l++
				i.asm += fmt.Sprintf("&%x", i.dst)
			case 2: // Constant. The operand is the constant 4
				i.dst = 4
				i.asm += fmt.Sprintf("#%d", i.dst)
			case 3: // Constant. The operand is the constant 8
				i.dst = 8
				i.asm += fmt.Sprintf("#%d", i.dst)
			}
		}
	case 3: // CG
		dst = false
		switch i.as {
		case 0: // Constant. The operand is the constant 0
			i.dst = 0
		case 1: // Constant. The operand is the constant 1. There is no index word
			i.dst = 1
		case 2: // Constant. The operand is the constant 2
			i.dst = 2
		case 3: // Constant. The operand is the constant −1
			i.dst = 0xffff
		}
		i.asm += fmt.Sprintf("#%d", i.dst)
	}

	if dst {
		// Disassm with as and dst
		switch i.as {
		case 0: // Register direct, Rn
			i.asm += "r"
		case 1: // Indexed, X(Rn)
			i.l++ // Instruction long 2 word, second word is X
			i.asm += fmt.Sprintf("%x(r", i.hex[1])
		default: //Register indirect, @Rn or Indirect autoincrement, @Rn+
			i.asm += "@r"
		}
		i.asm += fmt.Sprintf("%d", i.dst)
		switch i.as {
		case 1: // Indexed, X(Rn)
			i.asm += ")"
		case 3: //Indirect autoincrement, @Rn+
			i.asm += "+"
		}
	}
}

func (i *Instruction) jmp() {
	code := i.hex[0]
	// Check Condition
	i.cond = mask(code, COND)
	i.asm = mnemonic[1][i.cond]
	i.l = 1 // All jmp instruction are 1 word long

	// Check Offset
	i.offs = mask(code, OFFS)
	i.offs <<= 1 // 2x
	// TODO (offs+2) += PC
	i.asm += fmt.Sprintf(" @%x", i.offs)
}

func (i *Instruction) two() {
	code := i.hex[0]
	// Check all bits of single instruction
	i.twooc = mask(code, twoOC)
	if i.twooc < 4 {
		i.l = 0
		i.asm = "Two Op:Error decoding instruction"
		return
	}
	i.l = 1
	i.twooc -= 4
	i.asm = mnemonic[2][i.twooc]

	// Check B/W OK in all instructions of two op
	i.bw = mask(code, BW)
	if i.bw == 0 {
		i.asm += ".w "
	} else {
		i.asm += ".b "
	}

	// Check as Address Source
	i.as = mask(code, AS)
	// Check source
	i.src = mask(code, SRC)
	// Check ad Address Destination
	i.ad = mask(code, AD)
	// Check Source2
	i.dst = mask(code, DST)

	src := true
	srcstring := ""
	switch i.src {
	case 0:
		switch i.as {
		case 1: // Symbolic. Equivalent to x(PC). The operand is in memory at address PC+x.
			// disassm nothing to do, in execution we need PC+X
		case 3: // Immediate. Equivalent to @PC+. The operand is the next word in the instruction stream.
			src = false
			i.src = i.hex[i.l]
			i.l++
			srcstring = fmt.Sprintf("#%x", i.src)
		}
	case 2: // SR
		if i.as != 0 {
			src = false
			switch i.as {
			case 1: // Absolute. The operand is in memory at address x.
				i.src = i.hex[i.l]
				i.l++
				srcstring = fmt.Sprintf("&%x", i.src)
			case 2: // Constant. The operand is the constant 4
				i.src = 4
				srcstring = fmt.Sprintf("#%d", i.dst)
			case 3: // Constant. The operand is the constant 8
				i.src = 8
				srcstring = fmt.Sprintf("#%d", i.dst)
			}
		}
	case 3: // CG
		src = false
		switch i.as {
		case 0: // Constant. The operand is the constant 0
			i.src = 0
		case 1: // Constant. The operand is the constant 1. There is no index word
			i.src = 1
		case 2: // Constant. The operand is the constant 2
			i.src = 2
		case 3: // Constant. The operand is the constant −1
			i.src = 0xffff
		}
		srcstring = fmt.Sprintf("#%d", i.src)
	}

	dst := true
	dststring := ""
	switch i.dst {
	case 0:
		if i.ad == 1 { // Symbolic. Equivalent to x(PC). The operand is in memory at address PC+x.
			// disassm nothing to do, in execution we need PC+X
		}
	case 2:
		if i.ad == 1 { // Absolute. The operand is in memory at address x.
			dst = false
			i.dst = i.hex[i.l]
			i.l++
			dststring = fmt.Sprintf("&%x", i.dst)
		}
	}

	if src {
		switch i.as {
		case 0: // Register direct, Rn
			srcstring += "r"
		case 1: // Indexed, X(Rn)
			srcstring += fmt.Sprintf("%x(r", i.hex[i.l])
			i.l++ // Instruction long 2 word
		default: //Register indirect, @Rn or Indirect autoincrement, @Rn+
			srcstring += "@r"
		}
		srcstring += fmt.Sprintf("%d", i.src)
		switch i.as {
		case 1: // Indexed, X(Rn)
			srcstring += ")"
		case 3: //Indirect autoincrement, @Rn+
			srcstring += "+"
		}
	}
	i.asm += srcstring + ","
	if dst {
		if i.ad == 0 { // Register direct, Rn
			dststring += "r"
		} else { // Indexed, X(Rn)
			dststring += fmt.Sprintf("%x(r", i.hex[i.l])
			i.l++ // Instruction long 2 word or 3
		}
		dststring += fmt.Sprintf("%d", i.dst)
		if i.ad != 0 {
			dststring += ")"
		}
	}
	i.asm += dststring
}
