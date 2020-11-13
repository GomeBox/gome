package graphics

import (
	"github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/primitives"
)

//Window contains information about the game's window
type Window interface {
	//Dimensions returns the size of the Window
	Dimensions() (primitives.Dimensions, error)
}

type window struct {
	internalWindow interfaces.Window
}

func (w *window) Dimensions() (primitives.Dimensions, error) {
	return w.internalWindow.Dimensions()
}
