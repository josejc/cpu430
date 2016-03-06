package cpu430

import (
	"bytes"
	"fmt"
)

// TODO: Control of size and memory limit

const (
	MEMORY_SIZE uint32 = 1048576 // 2^20 (20bit address)
)

// uint20 represents a 20 bit physical address
type uint20 uint32

// Represents the RAM memory
type Memory interface {
	Reset()                             // Sets all memory locations to zero
	Read(address uint20) (value uint16) // Return the value at memory address
	Write(address uint20, value uint16) // Write the value at memory address
	// TODO: Read and Write Bytes ;)
}

// Represents the memory using a map of uint16's.
type BasicMemory struct {
	m map[uint20]uint16
}

// Returns a pointer to a new BasicMemory with all memory initialized
// to zero.
func NewBasicMemory() *BasicMemory {
	return &BasicMemory{
		m: make(map[uint20]uint16),
	}
}

// Resets all memory locations to zero
func (mem *BasicMemory) Reset() {
	n := make(map[uint20]uint16)
	mem.m = n
}

// TODO: Check even address and limit of memory size
func (mem *BasicMemory) Read(address uint20) (value uint16) {
	value = mem.m[address]
	return
}

// TODO: Check even address and limit of memory size
func (mem *BasicMemory) Write(address uint20, value uint16) {
	mem.m[address] = value
	return
}

// TODO: Check even address, limit memory
func (mem *BasicMemory) RawDump(address uint20, size uint16) string {
	var buffer bytes.Buffer
	var data uint16

	long := address + uint20(size)
	for i := address; i < long; i = i + 2 {
		data = mem.m[i]
		buffer.WriteString(fmt.Sprintf("%04x ", data))
	}
	return buffer.String()
}

func (mem *BasicMemory) Dump(address uint20, size uint16) []string {
	const (
		LINE   = 16         // Long of line, 16 BYTES
		MAX    = 65535      // Max. size of dump memory
		N_LINE = MAX / LINE // Number of lines of the dump
	)
	var bufhex, bufasc bytes.Buffer
	dump := make([]string, N_LINE)

	ad := uint16(address)
	ad = ad & 0xfff0
	long := uint16(address) + size
	long = long | 0x000f
	l := (ad + long) / LINE // number of lines
	i := ad
	for j := 0; j < l; j++ {
		data = mem.m[i]
		i = i + 2
	}
	return dump
}

// TODO: function to load memory of file ;)
