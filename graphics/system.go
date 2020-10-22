package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type System interface {
	LoadTexture(fileName string) (Texture, error)
	LoadFont(fileName string, size int) (Font, error)
	CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (Texture, error)
	Window() Window
}
