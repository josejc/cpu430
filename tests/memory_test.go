package cpu430_test

import (
	"fmt"
	"github.com/josejc/cpu430"
	"math/rand"
	"testing"
)

func TestMemory(t *testing.T) {
	var v, expected uint16
	var address uint16

	address = uint16(rand.Uint32())
	address &= 0xfffe // Addres now is even
	m := cpu430.NewMemory()
	expected = 65535
	err := m.WriteW(address, expected)
	if err != nil {
		fmt.Println(err)
	}
	v, _ = m.ReadW(address)
	if v != expected {
		t.Error("Expected 65535, got ", v)
	}
	m.WriteW(address+2, 0x4142)
	m.WriteW(address+4, 0x4344)
	m.WriteW(address+6, 0x4546)
	m.WriteW(address+9, 0x6565)

	fmt.Println("Hex:", m.RawDumpHex(address, 16))
	fmt.Println("ASCII:", m.RawDumpASCII(address, 16))
	fmt.Println("Full:", m.Dump(address, 32))
	m.Reset()
	v, _ = m.ReadW(address)
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
	m.LoadIHEX("../samples/out.hex", 0)
	address = 0
	fmt.Println(m.Dump(address, 64))
}
