package cpu430

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMemory(t *testing.T) {
	var v, expected uint16
	var address uint16

	address = uint16(rand.Uint32())
	address &= 0xfffe // Addres now is even
	m := NewBasicMemory()
	expected = 65535
	err := m.Write(address, expected)
	if err != nil {
		fmt.Println(err)
	}
	v, _ = m.Read(address)
	if v != expected {
		t.Error("Expected 65535, got ", v)
	}
	m.Write(address+2, 0x4142)
	m.Write(address+4, 0x4344)
	m.Write(address+6, 0x4546)
	m.Write(address+9, 0x6565)
	fmt.Println(m.RawDumpHex(address, 16))
	fmt.Println(m.RawDumpAscii(address, 16))
	fmt.Println(m.Dump(address, 32))
	m.Reset()
	v, _ = m.Read(address)
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}
