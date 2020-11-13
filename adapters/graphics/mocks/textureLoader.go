package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type TextureLoader struct {
	OnLoad func(fileName string) (graphics.Texture, error)
}

func (t *TextureLoader) Load(fileName string) (graphics.Texture, error) {
	if t.OnLoad != nil {
		return t.OnLoad(fileName)
	}
	return nil, nil
}
