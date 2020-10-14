package graphics

import "github.com/GomeBox/gome/primitives"

//WindowAdapter is to be implemented by the graphics adapter
type WindowAdapter interface {
	//ShowWindow opens a game window
	ShowWindow(windowSettings *WindowSettings) error
}

//WindowSettings contains all the settings needed to show the game window
type WindowSettings struct {
	Fullscreen bool
	Resolution primitives.Dimensions
	Title      string
}
