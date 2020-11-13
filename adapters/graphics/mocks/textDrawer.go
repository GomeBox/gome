package mocks

import (
	"github.com/GomeBox/gome/primitives"
)

type TextDrawer struct {
	OnDraw       func(position *primitives.Point) error
	OnDrawF      func(position *primitives.PointF) error
	OnDimensions func() primitives.Dimensions
}

func (t *TextDrawer) Draw(position *primitives.Point) error {
	if t.OnDraw != nil {
		return t.OnDraw(position)
	}
	return nil
}

func (t *TextDrawer) DrawScaled(destRect *primitives.Rectangle) error {
	panic("implement me")
}

func (t *TextDrawer) DrawF(position *primitives.PointF) error {
	if t.OnDrawF != nil {
		return t.OnDrawF(position)
	}
	return nil
}

func (t *TextDrawer) DrawScaledF(destRect *primitives.RectangleF) error {
	panic("implement me")
}

func (t *TextDrawer) Dimensions() primitives.Dimensions {
	if t.OnDimensions != nil {
		return t.OnDimensions()
	}
	return primitives.Dimensions{}
}
