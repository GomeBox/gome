package sprite

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

type TextDrawer struct {
	text graphics.Text
}

func NewTextDrawer(text graphics.Text) *TextDrawer {
	return &TextDrawer{text: text}
}

func (drawer *TextDrawer) DrawTo(pos *primitives.PointF) error {
	return drawer.text.DrawF(pos)

}

func (drawer *TextDrawer) Dimensions() primitives.Dimensions {
	return drawer.text.Dimensions()
}
