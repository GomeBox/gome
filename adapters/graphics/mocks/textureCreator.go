package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type TextureCreator struct {
}

func (t *TextureCreator) Create(dimensions *primitives.Dimensions, color *primitives.Color) (graphics.Texture, error) {
	panic("implement me")
}
