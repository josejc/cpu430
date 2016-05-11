package main

import (
	"fmt"

	"github.com/josejc/cpu430"
)

func main() {
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
}
