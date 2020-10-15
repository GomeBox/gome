package graphics

import "github.com/GomeBox/gome/primitives"

type TextDrawer interface {
	Draw(position primitives.Point) error
	DrawScaled(destRect primitives.Rectangle) error
}
