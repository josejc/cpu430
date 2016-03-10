package cpu430

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

// TODO: Control of size and memory limit
// Now is limit uint16 = 2^16 -> 64KB

// uint20 represents a 20 bit physical address
//type uint20 uint32

// Represents the RAM memory
//type Memory interface {
//	Reset()                             // Sets all memory locations to zero
//	Read(address uint16) (value uint16) // Return the value at memory address
//	Write(address uint16, value uint16) // Write the value at memory address
// TODO: Read and Write Bytes ;)
//}

// Represents the memory using a map of uint16's.
type Memory struct {
	m map[uint16]uint16
}

// Returns a pointer to a new Memory with all memory initialized
// to zero.
func NewMemory() *Memory {
	return &Memory{
		m: make(map[uint16]uint16),
	}
}

// Resets all memory locations to zero
func (mem *Memory) Reset() {
	n := make(map[uint16]uint16)
	mem.m = n
}

// TODO: Check even address and limit of memory size
func (mem *Memory) Read(address uint16) (uint16, error) {
	if (address % 2) != 0 {
		return 0, errors.New("Miss aligned memory")
	}
	return mem.m[address], nil
}

// TODO: Check even address and limit of memory size
func (mem *Memory) Write(address uint16, value uint16) error {
	if (address % 2) != 0 {
		return errors.New("Miss aligned memory")
	}
	mem.m[address] = value
	return nil
}

// TODO: Check even address, limit memory
func (mem *Memory) RawDumpHex(address uint16, size uint16) string {
	var buffer bytes.Buffer
	var data uint16

	long := address + size
	for i := address; i < long; i = i + 2 {
		data = mem.m[i]
		buffer.WriteString(fmt.Sprintf("%04x ", data))
	}
	return buffer.String()
}

func ascii(code uint8) uint8 {
	if (code >= 32) && (code < 127) {
		return code
	}
	return '.'
}

// ;)
func (mem *Memory) RawDumpAscii(address uint16, size uint16) string {
	var buffer bytes.Buffer
	var data uint16
	var dh, dl uint8

	long := address + size
	for i := address; i < long; i = i + 2 {
		data = mem.m[i]
		dh = ascii(uint8((data & 0xff00) >> 8))
		dl = ascii(uint8(data & 0x00ff))
		buffer.WriteString(string(dh))
		buffer.WriteString(string(dl))
	}
	return buffer.String()
}

func (mem *Memory) Dump(address uint16, size uint16) []string {
	const (
		LINE = 0x10 // Size of dump bytes in a line
	)
	buffer := bytes.NewBuffer(nil)

	address &= 0xfff0
	adEnd := address + size
	adEnd |= 0x000f
	n_line := ((adEnd + 1) - address) / LINE
	dump := make([]string, n_line)
	for i := 0; address < adEnd; i++ {
		buffer.WriteString(fmt.Sprintf("%04x: ", address))
		buffer.WriteString(mem.RawDumpHex(address, LINE))
		buffer.WriteString(" ")
		buffer.WriteString(mem.RawDumpAscii(address, LINE))
		dump[i] = buffer.String()
		buffer = bytes.NewBuffer(nil)
		address += LINE
	}
	return dump
}

// TODO: function to load memory of file ;)
func loadIHEX(filePtr *string, address uint16) error {
	var r *strings.Reader
	var b byte
	var x, y, oldy int

	f, err := os.Open(*filePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		return errors.New("Can't open file")
	} else {
		input := bufio.NewScanner(f)
	header:
		// Read patterns and positions
		for input.Scan() {
			r = strings.NewReader(input.Text())
			b, _ = r.ReadByte()
			if b == '#' {
				b, _ = r.ReadByte()
				if b == 'P' {
					s := strings.Split(input.Text(), " ")
					x, _ = strconv.Atoi(s[1])
					y, _ = strconv.Atoi(s[2])
					x += (M / 2)
					y += (N / 2)
					oldy = y
				} else {
					fmt.Fprintf(os.Stderr, "ERROR: Expected Position or blocks not config parameters\n")
					return false
				}
			} else {
				p.x = x
				for cells := int(r.Size()); cells > 0; cells-- {
					p.y = y
					switch b {
					case '.':
						{
							//m[p] = 0
						}
					case '*':
						{
							m[p] = 1
						}
					default:
						{
							fmt.Fprintf(os.Stderr, "ERROR: Character not valid, only '.' or '*'\n")
							return false
						}
					}
					b, _ = r.ReadByte()
					y++
				}
			}
			x++
			y = oldy
		}
	}
	f.Close()
	return nil
	// NOTE: ignoring potential errors from input.Err()
}
