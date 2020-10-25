package sprite

import "github.com/GomeBox/gome/primitives"

type Interface interface {
	Dimensions() primitives.Dimensions
	Position() primitives.PointF
	SetPosition(x, y float32) error
	Draw() error
}
