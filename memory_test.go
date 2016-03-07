package cpu430

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMemory(t *testing.T) {
	var v, expected uint16
	var address uint20

	address = uint20(rand.Uint32())
	m := NewBasicMemory()
	expected = 65535
	m.Write(address, expected)
	v = m.Read(address)
	if v != expected {
		t.Error("Expected 65535, got ", v)
	}
	m.Write(address+2, 0x4142)
	m.Write(address+4, 0x4344)
	fmt.Println("Address: ", address)
	fmt.Println(m.RawDumpHex(address, 16))
	fmt.Println(m.RawDumpAscii(address, 16))
	fmt.Println(m.Dump(address, 32))
	m.Reset()
	v = m.Read(address)
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}
