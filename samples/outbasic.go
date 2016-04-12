package main

import "github.com/nsf/termbox-go"

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	draw()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			default:
				draw()
			}
		case termbox.EventResize:
			draw()
		}
	}
}
