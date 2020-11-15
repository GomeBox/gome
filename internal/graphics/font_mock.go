package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type FontMock struct {
	OnCreateText func(text string, color primitives.Color) (Text, error)
}

func (f *FontMock) CreateText(text string, color primitives.Color) (Text, error) {
	if f.OnCreateText != nil {
		return f.OnCreateText(text, color)
	}
	return nil, nil
}
