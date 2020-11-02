package graphics

import (
	"github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/primitives"
)

//Font represents a loaded font of a defined size
type Font interface {
	//CreateText creates a Text instance from the Font of the passed color, containing the passed text
	CreateText(value string, color primitives.Color) (Text, error)
}

type font struct {
	internalFont interfaces.Font
}

func (f *font) CreateText(value string, color primitives.Color) (Text, error) {
	t, err := f.internalFont.CreateText(value, color)
	if err != nil {
		return nil, err
	}
	return &text{internalText: t}, err
}
