package mocks

import (
	"github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type Font struct {
	OnCreateText func(text string, color primitives.Color) (interfaces.Text, error)
}

func (f *Font) CreateText(text string, color primitives.Color) (interfaces.Text, error) {
	if f.OnCreateText != nil {
		return f.OnCreateText(text, color)
	}
	return nil, nil
}
