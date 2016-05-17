package main

import (
	"fmt"

	"github.com/josejc/cpu430"
)

func main() {
	var a uint16

	m := cpu430.NewMemory()
	e := m.LoadIHEX("kk2.hex", 0)
	if e != nil {
		fmt.Println(e)
		return
	}
	// print the i_hex file
	// Address kk2 initial 0x100
	a = 0x0100
	i_hex := m.Dump(a, 64)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	fmt.Println("---")
	for n := 0; n < 5; n++ {
		i := cpu430.NewInstruction(m, a)
		fmt.Printf("--%#x: %v\n", a, i.Dissasm())
		a += 2 * i.Long()
	}
}
