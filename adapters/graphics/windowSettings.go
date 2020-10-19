package graphics

import "github.com/GomeBox/gome/primitives"

//WindowSettings contains all the settings needed to show the game window
type WindowSettings struct {
	//Fullscreen defines if the game window should use the whole screen.
	//The size of the game window is the Resolution, if Fullscreen is false
	Fullscreen bool
	//Resolution defines the resolution of the game window
	Resolution primitives.Dimensions
	//Title is the game's title that is used as the window title
	Title string
}
