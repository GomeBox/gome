package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type FontLoader struct {
}

func (f *FontLoader) Load(fileName string, size int) (graphics.TextCreator, error) {
	panic("implement me")
}
