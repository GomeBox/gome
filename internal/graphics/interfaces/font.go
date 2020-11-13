package interfaces

import (
	"github.com/GomeBox/gome/primitives"
)

type Font interface {
	CreateText(text string, color primitives.Color) (Text, error)
}
