package graphics

import "github.com/GomeBox/gome/primitives"

//WindowAdapter is to be implemented by the graphics adapter
type WindowAdapter interface {
	//OpenWindow opens a game window
	OpenWindow(windowSettings *WindowSettings) error
	//IsOpen returns true if the window was opened
	IsOpen() bool
	//Size returns the window's dimensions
	Size() primitives.Dimensions
}
