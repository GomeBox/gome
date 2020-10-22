package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type Text interface {
	Draw(position *primitives.Point) error
	DrawF(position *primitives.PointF) error
}
