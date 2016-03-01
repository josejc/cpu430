package cpu430

import (
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
	m.Reset()
	v = m.Read(address)
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}
