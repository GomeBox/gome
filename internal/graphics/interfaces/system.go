package interfaces

import (
	"github.com/GomeBox/gome/primitives"
)

type System interface {
	LoadTexture(filename string) (Texture, error)
	LoadFont(fileName string, size int) (Font, error)
	CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (Texture, error)
	Window() Window
	Clear() error
	Present() error
}
