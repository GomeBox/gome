package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type Texture interface {
	Draw(source, dest *primitives.Rectangle) error
	DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error
	Dimensions() primitives.Dimensions
}
