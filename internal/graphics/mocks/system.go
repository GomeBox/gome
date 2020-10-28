package mocks

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

type System struct {
}

func (s System) LoadTexture(filename string) (graphics.Texture, error) {
	panic("implement me")
}

func (s System) LoadFont(fileName string, size int) (graphics.Font, error) {
	panic("implement me")
}

func (s System) CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (graphics.Texture, error) {
	panic("implement me")
}

func (s System) Window() graphics.Window {
	panic("implement me")
}
