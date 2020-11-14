package interfaces

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type System interface {
	LoadTexture(filename string) (interfaces.Texture, error)
	LoadFont(fileName string, size int) (interfaces.Font, error)
	CreateTexture(dimensions primitives.Dimensions, color primitives.Color) (interfaces.Texture, error)
	Window() Window
	Clear() error
	Present() error
}
