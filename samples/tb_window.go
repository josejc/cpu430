package main

import
//"fmt"

(
	"github.com/hoffoo/termboxutil"
	termbox "github.com/nsf/termbox-go"
	//"strings"
)

func main() {
	// create us a new screen
	screen := termboxutil.Screen{}

	// make us a window, not displayed yet
	searchWindow := screen.NewWindow(
		termbox.ColorWhite,   // window foreground
		termbox.ColorDefault, // window background
		termbox.ColorGreen,   // selected item foreground
		termbox.ColorBlack)   // selected item background

	// this is the main window, since created last it is already focused
	// if we created out of order we can call screen.Focus(&mainWindow)
	mainWindow := screen.NewWindow(termbox.ColorWhite, termbox.ColorDefault, termbox.ColorGreen, termbox.ColorBlack)
	mainWindow.Scrollable(true)

	mainWindow.Draw(TODO) // draw mainWindow's contents
	// etc...

	// when we are done flush the output to the screen
	termbox.Flush()

	// now create us a handler func for the main window
	mainWindow.CatchEvent = func(event termbox.Event) {
		// simple scrolling
		if event.Ch == 'j' || event.Key == termbox.KeyArrowDown {
			mainWindow.NextRow()
		} else if event.Ch == 'k' || event.Key == termbox.KeyArrowUp {
			mainWindow.PrevRow()
		} else if event.Ch == 'i' {

			// this is the data that the new screen will show
			// calling mainWindow.CurrentRow() here to find out where on the screen we are
			searchResult := makeSomeOutputForSearchWindow(mainWindow.CurrentRow().Text)

			// fill the new data and display this window
			searchWindow.Draw(searchResult)

			// focus so we get input to our searchWindow handler func CatchEvent
			screen.Focus(&searchWindow)

			// show output
			termbox.Flush()
			return
		}

		// setup any updates (for example PrevRow or Scrolling)
		mainWindow.Redraw()
		// show output
		termbox.Flush()
	}

	// handler for the search window
	searchWindow.CatchEvent = func(event termbox.Event) {
		// if we quit we dont want to see search results anymore
		if event.Ch == 'q' || event.Key == termbox.KeyEsc {

			// draw the mainWindow again
			mainWindow.Redraw()

			// set is as focused
			screen.Focus(&mainWindow)

			// show output
			termbox.Flush()

			return
		}

		// otherwise just redraw with whatever other updates we want here
		searchWindow.Redraw()
		termbox.Flush()
	}

	// now start looping for input - this will block, passing any termbox
	// event to the currently focused window
	screen.Loop()
}
