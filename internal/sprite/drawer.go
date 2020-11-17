package sprite

import "github.com/GomeBox/gome/primitives"

type drawer interface {
	DrawTo(pos primitives.PointF) error
	Dimensions() primitives.Dimensions
}
