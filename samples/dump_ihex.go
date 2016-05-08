package main

import (
	"fmt"

	"github.com/josejc/cpu430"
)

func main() {
	var a uint16

	m := cpu430.NewMemory()
	e := m.LoadIHEX("test.hex", 0)
	if e != nil {
		fmt.Println(e)
		return
	}
	// print the i_hex file
	// Beware with the address
	// Address kk2-256d, kk-16d
	// test-48d
	i_hex := m.Dump(48, 64)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	fmt.Println("---")

	fmt.Println("Example single operand")
	fmt.Println("Assembly: rrc.w r5")
	fmt.Println("Address: 0x0100, Instruction code: 0x1005")
	a = 0x0100
	m.WriteW(a, 0x1005)
	i_hex = m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i := m.Decode(a)
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
	i = m.Decode(a)
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
	i = m.Decode(a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")

}
