package mocks

import "github.com/GomeBox/gome/primitives"

type Texture struct {
	OnDraw       func(source *primitives.Rectangle, dest primitives.RectangleF) error
	OnDimensions func() primitives.Dimensions
}

func (t *Texture) Unload() error {
	panic("implement me")
}

func (t *Texture) Unloaded() bool {
	panic("implement me")
}

func (t *Texture) Draw(source *primitives.Rectangle, dest primitives.RectangleF) error {
	if t.OnDraw != nil {
		return t.OnDraw(source, dest)
	}
	return nil
}

func (t *Texture) Dimensions() primitives.Dimensions {
	if t.OnDimensions != nil {
		return t.OnDimensions()
	}
	return primitives.Dimensions{}
}
