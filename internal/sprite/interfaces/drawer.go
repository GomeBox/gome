package interfaces

import "github.com/GomeBox/gome/primitives"

type Drawer interface {
	DrawTo(pos primitives.PointF) error
	Dimensions() primitives.Dimensions
}
