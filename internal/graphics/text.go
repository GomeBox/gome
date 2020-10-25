package graphics

import (
	adapters "github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

func newText(drawer adapters.TextDrawer) graphics.Text {
	return &text{drawer: drawer}
}

type text struct {
	drawer adapters.TextDrawer
}

func (text *text) Draw(position *primitives.Point) error {
	err := text.drawer.Draw(position)
	if err != nil {
		return err
	}
	return nil
}

func (text *text) DrawF(position *primitives.PointF) error {
	err := text.drawer.DrawF(position)
	if err != nil {
		return err
	}
	return nil
}

func (text *text) Dimensions() primitives.Dimensions {
	return text.Dimensions()
}
