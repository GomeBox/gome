package graphics

import "github.com/GomeBox/gome/primitives"

type TextMock struct {
	OnDraw       func(position *primitives.Point) error
	OnDrawF      func(position *primitives.PointF) error
	OnDimensions func() primitives.Dimensions
}

func (t TextMock) Draw(position *primitives.Point) error {
	if t.OnDraw != nil {
		return t.OnDraw(position)
	}
	return nil
}

func (t TextMock) DrawF(position *primitives.PointF) error {
	if t.OnDrawF != nil {
		return t.OnDrawF(position)
	}
	return nil
}

func (t TextMock) Dimensions() primitives.Dimensions {
	if t.OnDimensions != nil {
		return t.OnDimensions()
	}
	return primitives.Dimensions{}
}
