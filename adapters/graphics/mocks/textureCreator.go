package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type TextureCreator struct {
	OnCreate func(dimensions primitives.Dimensions, color primitives.Color) (graphics.Texture, error)
}

func (t *TextureCreator) Create(dimensions primitives.Dimensions, color primitives.Color) (graphics.Texture, error) {
	if t.OnCreate != nil {
		return t.OnCreate(dimensions, color)
	}
	return nil, nil
}
