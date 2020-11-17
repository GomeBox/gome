package mocks

import "github.com/GomeBox/gome/primitives"

type Text struct {
	OnDraw       func(position primitives.PointF) error
	OnDimensions func() primitives.Dimensions
}

func (t Text) Draw(position primitives.PointF) error {
	if t.OnDraw != nil {
		return t.OnDraw(position)
	}
	return nil
}

func (t Text) Dimensions() primitives.Dimensions {
	if t.OnDimensions != nil {
		return t.OnDimensions()
	}
	return primitives.Dimensions{}
}
