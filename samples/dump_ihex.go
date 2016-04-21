package main

import (
	"fmt"

	"github.com/josejc/cpu430"
)

func main() {
	var i uint16

	m := cpu430.NewMemory()
	e := m.LoadIHEX("kk2.hex", 0)
	if e != nil {
		fmt.Println(e)
		return
	}
	// print the i_hex file
	i_hex := m.Dump(256, 64)
	for _, line := range i_hex {
		fmt.Println(line)
	}
	for i = 0; i < 64; i += 2 {
		oc, _ := m.Read(i)
		fmt.Printf("Ad: %x value %x -- %v\n", i, oc, cpu430.Opcode(oc))
	}
}
