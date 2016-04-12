package main

import "github.com/nsf/termbox-go"

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	setString(10, 2, "Initial presentation", termbox.ColorBlue, termbox.ColorBlack)
	termbox.Flush()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			default:
				draw()
				termbox.Flush()
			}
		case termbox.EventResize:
			termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
			setString(10, 7, "Resize the window", termbox.ColorYellow, termbox.ColorBlack)
			termbox.Flush()
		}
	}
}

func setString(x, y int, s string, fg, bg termbox.Attribute) {
	for i, ch := range s {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}
}

func draw() {
	setString(10, 4, "Welcome to CPU430", termbox.ColorRed|termbox.AttrBold, termbox.ColorBlack)
	setString(10, 5, "Press ESC key to exit.", termbox.ColorWhite, termbox.ColorBlack)
}
