package mocks

import "github.com/GomeBox/gome/primitives"

type Texture struct {
	OnDraw       func(source, dest *primitives.Rectangle) error
	OnDrawF      func(source *primitives.Rectangle, dest *primitives.RectangleF) error
	OnDimensions func() primitives.Dimensions
}

func (t Texture) Draw(source, dest *primitives.Rectangle) error {
	if t.OnDraw != nil {
		return t.OnDraw(source, dest)
	}
	return nil
}

func (t Texture) DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error {
	if t.OnDrawF != nil {
		return t.OnDrawF(source, dest)
	}
	return nil
}

func (t Texture) Dimensions() primitives.Dimensions {
	if t.OnDimensions != nil {
		return t.OnDimensions()
	}
	return primitives.Dimensions{}
}
