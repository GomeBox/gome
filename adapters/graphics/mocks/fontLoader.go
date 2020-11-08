package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type FontLoader struct {
	OnLoad func(fileName string, size int) (graphics.TextCreator, error)
}

func (f *FontLoader) Load(fileName string, size int) (graphics.TextCreator, error) {
	if f.OnLoad != nil {
		return f.OnLoad(fileName, size)
	}
	return nil, nil
}
