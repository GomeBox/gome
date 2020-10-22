package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Font interface {
	CreateText(text string, color primitives.Color) (Text, error)
}

func newFont(creator graphics.TextCreator) Font {
	font := new(font)
	font.creator = creator
	return font
}

type font struct {
	creator graphics.TextCreator
}

func (font *font) CreateText(text string, color primitives.Color) (Text, error) {
	drawer, err := font.creator.Create(text, color)
	if err != nil {
		return nil, err
	}
	return newText(drawer), nil
}
