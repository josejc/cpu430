package main

import (
	//"fmt"
	"github.com/nsf/termbox-go"
	//"strings"
)

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func draw_all() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	tbprint(0, 0, termbox.ColorMagenta, coldef, "Press 'esc' to quit")
	//tbprint i_hex file
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	draw_all()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			default:
				{
					tbprint(1, 1, termbox.ColorBlack, termbox.ColorDefault, "The key is not 'esc'")
					termbox.Flush()
				}
			}
		case termbox.EventResize:
			draw_all()
		}
	}

}
