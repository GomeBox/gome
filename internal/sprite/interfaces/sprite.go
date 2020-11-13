package interfaces

import "github.com/GomeBox/gome/primitives"

type Sprite interface {
	Dimensions() primitives.Dimensions
	Position() primitives.PointF
	SetPosition(x, y float32)
	Draw() error
}
