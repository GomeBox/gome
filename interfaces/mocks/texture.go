package mocks

import "github.com/GomeBox/gome/primitives"

type Texture struct {
}

func (t Texture) Unload() error {
	panic("implement me")
}

func (t Texture) Unloaded() bool {
	panic("implement me")
}

func (t Texture) Draw(source *primitives.Rectangle, dest primitives.RectangleF) error {
	panic("implement me")
}

func (t Texture) Dimensions() primitives.Dimensions {
	panic("implement me")
}
