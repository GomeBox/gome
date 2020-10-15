package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Text interface {
	Draw(position *primitives.Point) error
	DrawF(position *primitives.PointF) error
}

func newText(drawer graphics.TextDrawer) Text {
	return &text{drawer: drawer}
}

type text struct {
	drawer graphics.TextDrawer
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
