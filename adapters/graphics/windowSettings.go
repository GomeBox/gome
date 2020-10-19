package graphics

import "github.com/GomeBox/gome/primitives"

//WindowSettings contains all the settings needed to show the game window
type WindowSettings struct {
	Fullscreen bool
	Resolution primitives.Dimensions
	Title      string
}
