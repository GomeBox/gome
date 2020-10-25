package sprite

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

type textDrawer struct {
	text graphics.Text
}

func newTextDrawer(text graphics.Text) *textDrawer {
	return &textDrawer{text: text}
}

func (drawer *textDrawer) DrawTo(pos *primitives.PointF) error {
	return drawer.text.DrawF(pos)

}

func (drawer *textDrawer) Dimensions() primitives.Dimensions {
	return drawer.text.Dimensions()
}
