package graphics

import (
	adapters "github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

func newFont(creator adapters.TextCreator) graphics.Font {
	font := new(font)
	font.creator = creator
	return font
}

type font struct {
	creator adapters.TextCreator
}

func (font *font) CreateText(text string, color primitives.Color) (graphics.Text, error) {
	drawer, err := font.creator.Create(text, color)
	if err != nil {
		return nil, err
	}
	return newText(drawer), nil
}
