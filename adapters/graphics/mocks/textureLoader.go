package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type TextureLoader struct {
}

func (t *TextureLoader) Load(fileName string) (graphics.Texture, error) {
	panic("implement me")
}
