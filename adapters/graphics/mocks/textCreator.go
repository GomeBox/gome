package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type TextCreator struct {
	OnCreate func(text string, color primitives.Color) (graphics.TextDrawer, error)
}

func (t *TextCreator) Create(text string, color primitives.Color) (graphics.TextDrawer, error) {
	if t.OnCreate != nil {
		return t.OnCreate(text, color)
	}
	return nil, nil
}
