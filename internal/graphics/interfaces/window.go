package interfaces

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/interfaces"
)

type Window interface {
	interfaces.Window
	//Dimensions() (primitives.Dimensions, error)
	Open(settings *graphics.WindowSettings) error
}
