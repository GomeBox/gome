package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

//Font represents a loaded font of a defined size
type Font interface {
	//CreateText creates a Text instance from the Font of the passed color, containing the passed text
	CreateText(text string, color primitives.Color) (Text, error)
}
