package graphics

import "github.com/GomeBox/gome/primitives"

type TextureDrawer interface {
	Draw(source, dest *primitives.Rectangle) error
	DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error
}
