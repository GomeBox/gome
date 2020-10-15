package graphics

import "github.com/GomeBox/gome/primitives"

type TextCreator interface {
	Create(text string, color primitives.Color) (TextDrawer, error)
}
