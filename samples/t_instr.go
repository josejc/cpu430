package main

import (
	"fmt"

	"github.com/josejc/cpu430"
)

func main() {
	var a uint16

	m := cpu430.NewMemory()

	fmt.Println("Example single operand")
	fmt.Println("Assembly: rrc.w r5")
	fmt.Println("Address: 0x0100, Instruction code: 0x1005")
	a = 0x0100
	m.WriteW(a, 0x1005)
	i_hex := m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i := cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")

	fmt.Println("Example jmp")
	fmt.Println("Assembly: jc main")
	fmt.Println("Instruction code: 0x2fe4")
	// a = 0x0100
	m.WriteW(a, 0x2fe4)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i = cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")

	fmt.Println("Example double operand")
	fmt.Println("Assembly: mov.w r5,r4")
	fmt.Println("Instruction code: 0x4504")
	// a = 0x0100
	m.WriteW(a, 0x4504)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i = cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")

	fmt.Println("Example--")
	fmt.Println("Assembly: mov.w @r5,r6")
	fmt.Println("Instruction code: 0x4526")
	// a = 0x0100
	m.WriteW(a, 0x4526)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i = cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")

	fmt.Println("Assembly: mov.w @r5+,r6")
	fmt.Println("Instruction code: 0x4536")
	// a = 0x0100
	m.WriteW(a, 0x4536)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i = cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")

	fmt.Println("SPECIAL CASE")
	fmt.Println("Assembly: mov.w @PC+, r6")
	fmt.Println("Assembly: mov.w #100,r6")
	fmt.Println("Instruction code: 0x4536, 0x0100")
	// a = 0x0100
	m.WriteW(a, 0x4036)
	m.WriteW(a+2, 0x0100)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i = cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---DON'T WORK")

	/* The same case...
	fmt.Println("Assembly: MOV.W #0x123a,R15")
	fmt.Println("Instruction code: 0x4536, 0x123a")
	// a = 0x0100
	m.WriteW(a, 0x403f)
	m.WriteW(a+2, 0x123a)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i = cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")
	*/

	fmt.Println("Assembly: MOV.B &P1IN,R15")
	fmt.Println("Instruction code: 0x425F, 0x0020")
	// a = 0x0100
	m.WriteW(a, 0x425f)
	m.WriteW(a+2, 0x0020)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i = cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")

	fmt.Println("Assembly: MOV.B #0x00f3,&P1OUT")
	fmt.Println("Instruction code: 0x40F2, 0x00F3, 0x0021")
	// a = 0x0100
	m.WriteW(a, 0x40f2)
	m.WriteW(a+2, 0x00f3)
	m.WriteW(a+4, 0x0021)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i = cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")
}
