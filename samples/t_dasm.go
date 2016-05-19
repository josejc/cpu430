package main

import (
	"fmt"
	"github.com/josejc/cpu430"
	"strings"
)

func main() {
	var a uint16

	m := cpu430.NewMemory()
	e := m.LoadIHEX("out.hex", 0)
	if e != nil {
		fmt.Println(e)
		return
	}
	// print the i_hex file
	// Address out.hex initial 0x0000
	a = 0x0000
	i_hex := m.Dump(a, 64)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	fmt.Println("---")
	for n := 0; n < 10; n++ {
		i := cpu430.NewInstruction(m, a)
		// For fill the values of instruction call Opcode o Disassm ;)
		i.Opcode()
		if i.Long() != 0 {
			b := strings.Repeat(" ", (20 - len(i.Hex())))
			fmt.Printf("%04x:%v %v %v\n", a, i.Hex(), b, i.Dissasm())
			a += 2 * i.Long()
		}
	}
}
