/*
A simple program to convert an Intel Hex file into binary

Al Williams -- DDJ
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Process a line (string, output file, true for ASCII output)
func procline(s string, w *bufio.Writer, text bool) {
	csum := 0 // checksum
	// skip empty lines
	if len(s) == 0 {
		return
	}
	// find :
	n := strings.Index(s, ":")
	if n == -1 {
		fmt.Fprintln(os.Stderr, "Warning: invalid line found")
		return
	}

	// read 2 byte record length
	llen, err := strconv.ParseInt(s[n+1:n+3], 16, 8)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Length Parse Error on "+s, ":", s[n+1:n+3])
	}
	csum += int(llen)
	// read address, if not consecutive, should warn to stderr but didn't
	add, err := strconv.ParseInt(s[n+5:n+7]+s[n+3:n+5], 16, 16)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Address Parse Error on "+s, ":", s[n+3:n+8])
	}
	// update checksum with two bytes
	csum += int(add & 0xFF)
	csum += int(add >> 8)
	// 00= normal record, 01=EOF, 02=Ext address
	switch s[n+7 : n+9] {
	case "00":
	case "01":
		csum += 1
		return // should we process more?
	case "02":
		csum += 2
		fmt.Fprintln(os.Stderr, "Warning: ExtADDR record ignored")
		return
	default:
		// nevermind checksum if we don't know the type
		fmt.Fprintln(os.Stderr, "Unknown")
	}
	// Process data bytes
	for i := 0; i < int(llen); i++ {
		byt, _ := strconv.ParseInt(s[n+9+i*2:n+9+i*2+2], 16, 16)
		// do output
		if text {
			fmt.Fprintln(w, byt)
		} else {
			tmp := [...]byte{byte(byt)}
			w.Write(tmp[0:1])
		}
		csum += int(byt)
	}
	// Checksum check
	csum = 0x100 - (csum & 0xFF) // computed checksum
	// read checksum
	rcsum, _ := strconv.ParseInt(s[n+9+int(llen)*2:n+9+int(llen)*2+2], 16, 16)
	// better match
	if csum != int(rcsum) {
		fmt.Fprintln(os.Stderr, "Checksum mismatch in ", s)
	}
}

// Process a file by repeatedly calling procline
func dofile(i, o *os.File, text bool) int {
	var buf []byte
	var err error
	rbuf := bufio.NewReader(i)
	wbuf := bufio.NewWriter(o)
	defer wbuf.Flush()
	for err == nil {
		buf, _, err = rbuf.ReadLine()
		if err == nil {
			procline(string(buf), wbuf, text)
		}
	}
	return 0
}

// Main, process arguments, open I/O files, etc.
func main() {
	ascii := false
	flag.BoolVar(&ascii, "a", false, "Set to output ASCII hex bytes instead of binary")
	flag.Parse()
	if len(flag.Args()) != 2 {
		// print error message
		fmt.Fprintln(os.Stderr, "Usage: go run hex2bin.go [-a] infile outfile")
		fmt.Fprintln(os.Stderr, " use - for stdin or stdout")
		fmt.Fprintln(os.Stderr, " -a emits raw ASCII instead of raw binary")
		// exit
		os.Exit(1)
	}
	var infile *os.File
	var outfile *os.File
	var err error
	// use - for standard streams
	if flag.Args()[0] == "-" {
		err = nil
		infile = os.Stdin
	} else {
		infile, err = os.Open(flag.Args()[0])
		if err != nil {
			defer infile.Close()
		}

	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if flag.Args()[1] == "-" {
		err = nil
		outfile = os.Stdout
	} else {
		outfile, err = os.Create(flag.Args()[1])
		if err != nil {
			defer outfile.Close()
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}

	os.Exit(dofile(infile, outfile, ascii))
}
