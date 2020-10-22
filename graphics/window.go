package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type Window interface {
	Dimensions() (primitives.Dimensions, error)
}
