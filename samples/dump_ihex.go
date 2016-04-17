package main

import (
	"fmt"
	"github.com/josejc/cpu430"
)

func main() {
	m := cpu430.NewMemory()
	e := m.LoadIHEX("out.hex", 0)
	if e != nil {
		fmt.Println(e)
		return
	}
	// print the i_hex file
	i_hex := m.Dump(0, 64)
	for _, line := range i_hex {
		fmt.Println(line)
	}
}
