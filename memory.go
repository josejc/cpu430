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
func (mem *BasicMemory) RawDumpHex(address uint20, size uint16) string {
	var buffer bytes.Buffer
	var data uint16

	long := address + uint20(size)
	for i := address; i < long; i = i + 2 {
		data = mem.m[i]
		buffer.WriteString(fmt.Sprintf("%04x ", data))
	}
	return buffer.String()
}

func ascii(dx uint8) uint8 {
	if (dx >= 32) && (dx < 127) {
		return dx
	}
	return '.'
}

// ;)
func (mem *BasicMemory) RawDumpAscii(address uint20, size uint16) string {
	var buffer bytes.Buffer
	var data uint16
	var dh, dl uint8

	long := address + uint20(size)
	for i := address; i < long; i = i + 2 {
		data = mem.m[i]
		dh = ascii(uint8((data & 0xff00) >> 8))
		dl = ascii(uint8(data & 0x00ff))
		buffer.WriteString(string(dh))
		buffer.WriteString(string(dl))
	}
	return buffer.String()
}

func (mem *BasicMemory) Dump(address uint20, size uint16) []string {
	const (
		LINE = 0x10 // Size of dump bytes in a line
	)
	buffer := bytes.NewBuffer(nil)

	ad := uint16(address)
	ad &= 0xfff0
	adEnd := uint16(address) + size
	adEnd |= 0x000f
	n_line := ((adEnd + 1) - ad) / LINE
	fmt.Println("address:", address, "ad:", ad, "adEnd:", adEnd, "nline:", n_line)
	dump := make([]string, n_line)
	for i := 0; ad < adEnd; i++ {
		buffer.WriteString(fmt.Sprintf("%04x: ", ad))
		fmt.Println("uint20-ad", uint20(ad))
		fmt.Println(mem.RawDumpHex(uint20(ad), LINE))
		buffer.WriteString(mem.RawDumpHex(uint20(ad), LINE))
		buffer.WriteString(" ")
		buffer.WriteString(mem.RawDumpAscii(uint20(ad), LINE))
		dump[i] = buffer.String()
		buffer = bytes.NewBuffer(nil)
		ad += LINE
	}
	return dump
}

// TODO: function to load memory of file ;)
