package cpu430

import ()

// TODO: Control of size and memory limit

const (
	MEMORY_SIZE uint32 = 1048576 // 2^20 (20bit address)
)

// Represents the RAM memory
type Memory interface {
	Reset()                             // Sets all memory locations to zero
	Read(address uint32) (value uint16) // Return the value at memory address
	Write(address uint16, value uint16) // Write the value at memory address
	// TODO: Read and Write Bytes ;)
}

// Represents the memory using a map of uint16's.
type BasicMemory struct {
	m map[uint32]uint16
}

// Returns a pointer to a new BasicMemory with all memory initialized
// to zero.
func NewBasicMemory() *BasicMemory {
	return &BasicMemory{
		m: make(map[uint32]uint16),
	}
}

// Resets all memory locations to zero
func (mem *BasicMemory) Reset() {
	n := make(map[uint32]uint16)
	mem.m = n
}

// TODO: Check even address and limit of memory size
func (mem *BasicMemory) Read(address uint32) (value uint16) {
	value = mem.m[address]
	return
}

// TODO: Check even address and limit of memory size
func (mem *BasicMemory) Write(address uint32, value uint16) {
	mem.m[address] = value
	return
}

// TODO: function to load memory of file ;)
