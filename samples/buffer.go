package glib

import (
	//"fmt"
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
)

func printAt(b Buffer, x, y int, s string, fg, bg termbox.Attribute) {
	for _, r := range s {
		b.SetCell(x, y, r, fg, bg)
		x++
	}
}

type Buffer interface {
	SetCell(x, y int, ch rune, fg, bg termbox.Attribute)
	GetCell(x, y int) (rune, termbox.Attribute, termbox.Attribute)
	Size() (int, int)
	CellBuffer() []termbox.Cell
}

type TermboxBufferType struct {
}

var TermboxBuffer TermboxBufferType

func (b TermboxBufferType) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
}

func (b TermboxBufferType) GetCell(x, y int) (rune, termbox.Attribute, termbox.Attribute) {
	cells := termbox.CellBuffer()
	w, _ := termbox.Size()
	cell := &cells[y*w+x]
	return cell.Ch, cell.Fg, cell.Bg
}

func (b TermboxBufferType) Size() (int, int) {
	return termbox.Size()
}

func (b TermboxBufferType) CellBuffer() []termbox.Cell {
	return termbox.CellBuffer()
}

type MemBuffer struct {
	w, h  int
	cells []termbox.Cell
}

func MakeMemBuffer(w, h int) *MemBuffer {
	return &MemBuffer{w, h, make([]termbox.Cell, w*h)}
}

func (b *MemBuffer) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	if x < 0 || x >= b.w {
		return
	}
	if y < 0 || y >= b.h {
		return
	}
	b.cells[y*b.w+x] = termbox.Cell{ch, fg, bg}
}

func (b *MemBuffer) GetCell(x, y int) (rune, termbox.Attribute, termbox.Attribute) {
	cell := &b.cells[y*b.w+x]
	return cell.Ch, cell.Fg, cell.Bg
}

func (b *MemBuffer) Size() (int, int) {
	return b.w, b.h
}

func (b *MemBuffer) CellBuffer() []termbox.Cell {
	return b.cells
}

const (
	CHAR    = 1
	FG      = 2
	BG      = 4
	ALL     = CHAR | FG | BG
	ATTRIBS = FG | BG
)

func GetBufferRow(src Buffer, row int) []termbox.Cell {
	w, _ := src.Size()
	return src.CellBuffer()[row*w : row*w+w]
}

func FillArea(dst Buffer, area rect.Rectangle, ch rune, fg, bg termbox.Attribute, filltype int) {
	area, ok := rect.Intersection(rect.WH(dst.Size()), area)

	if ok {
		for x := area.Min.X; x < area.Max.X; x++ {
			for y := area.Min.Y; y < area.Max.Y; y++ {
				if filltype == ALL {
					dst.SetCell(x, y, ch, fg, bg)
				} else {
					nch, nfg, nbg := dst.GetCell(x, y)
					if filltype&CHAR != 0 {
						nch = ch
					}
					if filltype&FG != 0 {
						nfg = fg
					}
					if filltype&BG != 0 {
						nbg = bg
					}
					dst.SetCell(x, y, nch, nfg, nbg)
				}
			}
		}
	}
}

func Fill(dst Buffer, ch rune, fg, bg termbox.Attribute, filltype int) {
	FillArea(dst, rect.WH(dst.Size()), ch, fg, bg, filltype)
}

func BlitBuffer(src, dst Buffer, xpos, ypos int) {
	srcSz := rect.WH(src.Size())
	srcSz.Translate(rect.Vec{xpos, ypos})

	dstSz := rect.WH(dst.Size())

	if intersect, ok := rect.Intersection(srcSz, dstSz); ok {
		for y := intersect.Min.Y; y < intersect.Max.Y; y++ {
			for x := intersect.Min.X; x < intersect.Max.X; x++ {
				nch, nfg, nbg := src.GetCell(x-xpos, y-ypos)
				dst.SetCell(x, y, nch, nfg, nbg)
			}
		}
	}
}
