package sprite

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/sprite/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type textDrawer struct {
	text gomeInterfaces.Text
}

func NewTextDrawer(text gomeInterfaces.Text) interfaces.Drawer {
	return &textDrawer{text: text}
}

func (drawer *textDrawer) DrawTo(pos primitives.PointF) error {
	return drawer.text.Draw(pos)
}

func (drawer *textDrawer) Dimensions() primitives.Dimensions {
	return drawer.text.Dimensions()
}
