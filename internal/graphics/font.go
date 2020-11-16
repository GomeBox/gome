package graphics

import (
	adapters "github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type font struct {
	creator adapters.TextCreator
}

func newFont(creator adapters.TextCreator) *font {
	font := new(font)
	font.creator = creator
	return font
}

func (font *font) CreateText(text string, color primitives.Color) (interfaces.Text, error) {
	drawer, err := font.creator.Create(text, color)
	if err != nil {
		return nil, err
	}
	return newText(drawer), nil
}
