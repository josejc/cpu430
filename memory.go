package cpu430

import (
	"bytes"
	//	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// TODO: Control of size and memory limit
// Now is limit uint16 = 2^16 -> 64KB

// uint20 represents a 20 bit physical address
//type uint20 uint32

// Represents the RAM memory
//type Memory interface {
//	Reset()                             // Sets all memory locations to zero
//	ReadW(address uint16) (value uint16) // Return the value at memory address
//	WriteW(address uint16, value uint16) // Write the value at memory address
//  ReadB and WriteB address=uint16, value=uint8=Bytes ;)
//}

// Memory represents using a map of uint8's (Bytes).
type Memory struct {
	m map[uint16]uint8
}

// NewMemory Returns a pointer to a new Memory with all memory initialized to zero.
func NewMemory() *Memory {
	return &Memory{
		m: make(map[uint16]uint8),
	}
}

// Reset all memory locations to zero
func (mem *Memory) Reset() {
	n := make(map[uint16]uint8)
	mem.m = n
}

// ReadB read a Byte
// TODO: Check even address and limit of memory size
func (mem *Memory) ReadB(address uint16) (uint8, error) {
	return mem.m[address], nil
}

// ReadW read a Word (2xByte)
// TODO: Check even address and limit of memory size
func (mem *Memory) ReadW(address uint16) (uint16, error) {
	if (address % 2) != 0 {
		return 0, errors.New("Miss aligned memory")
	}
	dl := mem.m[address]
	dh := mem.m[address+1]
	dx := uint16(dh)
	dx = dx << 8
	dx = dx | uint16(dl)
	return dx, nil
}

// WriteB write a Byte
// TODO: Check even address and limit of memory size
func (mem *Memory) WriteB(address uint16, value uint8) error {
	mem.m[address] = value
	return nil
}

// WriteW write a Word (2xByte)
// TODO: Check even address and limit of memory size
func (mem *Memory) WriteW(address uint16, value uint16) error {
	if (address % 2) != 0 {
		return errors.New("Miss aligned memory")
	}
	dh := uint8((value & 0xff00) >> 8)
	dl := uint8(value & 0x00ff)
	mem.m[address] = dl
	mem.m[address+1] = dh
	return nil
}

// RawDumpHex dump in a string the memory
// TODO: Check even address, limit memory
func (mem *Memory) RawDumpHex(address uint16, size uint16) string {
	var buffer bytes.Buffer
	var data uint8

	long := address + size
	// Only dump address even
	// Check even address?
	for i := address; i < long; i += 2 {
		data = mem.m[i]
		buffer.WriteString(fmt.Sprintf("%02x", data))
		data = mem.m[i+1]
		buffer.WriteString(fmt.Sprintf("%02x ", data))
	}
	return buffer.String()
}

func ascii(code uint8) uint8 {
	if (code >= 32) && (code < 127) {
		return code
	}
	return '.'
}

// RawDumpASCII dump in a string in ascci code
func (mem *Memory) RawDumpASCII(address uint16, size uint16) string {
	var buffer bytes.Buffer
	var data uint8

	long := address + size
	for i := address; i < long; i++ {
		data = ascii(mem.m[i])
		buffer.WriteString(string(data))
	}
	return buffer.String()
}

// Dump return the values in memory
func (mem *Memory) Dump(address uint16, size uint16) []string {
	const (
		LINE = 0x10 // Size of dump bytes in a line
	)
	buffer := bytes.NewBuffer(nil)

	address &= 0xfff0
	adEnd := address + size
	adEnd |= 0x000f
	nLine := ((adEnd + 1) - address) / LINE
	dump := make([]string, nLine)
	for i := 0; address < adEnd; i++ {
		buffer.WriteString(fmt.Sprintf("%04x: ", address))
		buffer.WriteString(mem.RawDumpHex(address, LINE))
		buffer.WriteString(" ")
		buffer.WriteString(mem.RawDumpASCII(address, LINE))
		dump[i] = buffer.String()
		buffer = bytes.NewBuffer(nil)
		address += LINE
	}
	return dump
}

// LoadIHEX load a file .hex in memory
// TODO: function to load memory of file ;)
func (mem *Memory) LoadIHEX(filename string, address uint16) error {

	data, err := ioutil.ReadFile(filename)
	s := string(data)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		// Check the format of Intel Hex
		// line[0] = Start code, one character, an ASCII colon ':'
		if line[0] != ':' {
			return errors.New("Intel HEX incorrect format")
		}
		// line[1:3] = Byte count, two hex digits, indicating the number of bytes (hex digit pairs) in the data field
		bc := line[1:3]
		//--DEBUGfmt.Println("Byte count:", bc) //--DEBUG
		nbc, _ := strconv.ParseUint(bc, 16, 16)
		ck := uint8(nbc)
		//--DEBUGfmt.Println("Summatory of values=", ck) //--DEBUG

		nbc = nbc << 1
		// line[3:7] = Address, four hex digits, representing the 16-bit beginning memory address offset of the data
		//--DEBUGad := line[3:7] //--DEBUG
		ah, _ := strconv.ParseUint(line[3:5], 16, 16)
		al, _ := strconv.ParseUint(line[5:7], 16, 16)
		address := uint16(ah)
		address = address << 8
		address = address | uint16(al)
		ck += uint8(ah) + uint8(al)
		//--DEBUGfmt.Println("Summatory of values=", ck) //--DEBUG

		//--DEBUGfmt.Println("Address:", ad, "=", address) //--DEBUG
		// line[7:9] = Record type, two hex digits, 00 to 05, defining the meaning of the data field.
		//   00-Data
		//   01-Enf of file
		//   03..05 Don't implemented :p
		rt, _ := strconv.ParseUint(line[7:9], 16, 16)
		ck += uint8(rt)
		//--DEBUGfmt.Println("Summatory of values=", ck) //--DEBUG

		//brt, _ := hex.DecodeString(rt)
		switch rt {
		case 0:
			//--DEBUGfmt.Println("Record Type:", rt) //--DEBUG
		case 1:
			return nil
		default:
			return errors.New("Record type, don't implemented")
		}
		// line[9:9+n] = Data, a sequence of n bytes of data, represented by 2n hex digits
		//DEBUGdata := line[9 : 9+nbc]    //--DEBUG
		//DEBIGfmt.Println("Data:", data) //--DEBUG
		for i := 9; i < int(9+nbc); i += 2 {
			// TODO Check limits, suppose address and nbytes are even
			data, _ := strconv.ParseUint(line[i:i+2], 16, 16)
			mem.WriteB(address, uint8(data))
			address++
			ck += uint8(data)
			//--DEBUGfmt.Println("Summatory of values=", ck) //--DEBUG
		}
		//--DEBUGfmt.Println("Summatory of values=", ck) //--DEBUG
		// Two's complement: (Bitwise xor FFh) and plus 1
		c2 := ck ^ 0xff
		c2++
		// line[9+n,9+n+2] = Checksum, two hex digits, a computed value that can be used to verify the record has no errors
		check := line[9+nbc : 9+nbc+2]
		ckk, _ := strconv.ParseUint(check, 16, 16)
		//--DEBUGfmt.Println("Checksum=", check) //--DEBUG
		if c2 != uint8(ckk) {
			//--DEBUGfmt.Println("Different Checksum:", check, "!=", c2) //--DEBUG
			return errors.New("Different checksum")
		}
		//   Checksum calculation: A record's checksum byte is the two's complement (negative) of the data checksum,
		//     which is the least significant byte (LSB) of the sum of all decoded byte values in the record preceding the checksum.
		//     It is computed by summing the decoded byte values and extracting the LSB of the sum (i.e., the data checksum),
		//     and then calculating the two's complement of the LSB (e.g., by inverting its bits and adding one)
	}
	return err
	// NOTE: ignoring potential errors from input.Err()
}
