package sprite

import (
	interfaces2 "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/sprite/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type textDrawer struct {
	text interfaces2.Text
}

func NewTextDrawer(text interfaces2.Text) interfaces.Drawer {
	return &textDrawer{text: text}
}

func (drawer *textDrawer) DrawTo(pos *primitives.PointF) error {
	return drawer.text.DrawF(pos)
}

func (drawer *textDrawer) Dimensions() primitives.Dimensions {
	return drawer.text.Dimensions()
}
