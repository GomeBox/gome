package interfaces

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

type System interface {
	LoadTexture(filename string) (graphics.Texture, error)
	LoadFont(fileName string, size int) (graphics.Font, error)
	CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (graphics.Texture, error)
	Window() graphics.Window
}
