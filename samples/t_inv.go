package main

import (
	"fmt"

	"github.com/josejc/cpu430"
)

func main() {
	var a uint16

	m := cpu430.NewMemory()

	a = 0x0100
	m.WriteW(a, 0x4701)
	i_hex := m.Dump(a, 0x0f)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	i := cpu430.NewInstruction(m, a)
	fmt.Printf("--%#x: %v\n", a, i.Dissasm())
	fmt.Printf("--%#x\n", i)
	fmt.Println("---")
}
