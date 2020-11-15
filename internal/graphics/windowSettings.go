package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type WindowSettings struct {
	Fullscreen bool
	Resolution primitives.Dimensions
	Title      string
}
