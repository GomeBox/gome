package interfaces

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Window interface {
	Dimensions() (primitives.Dimensions, error)
	Open(settings *graphics.WindowSettings) error
}
