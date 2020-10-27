package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

//System contains functions to access the graphics system
type System interface {
	//LoadTexture loads a texture from a file and returns that Texture
	LoadTexture(fileName string) (Texture, error)
	//LoadFont loads a font from a file, assigns it a size and returns a Font
	LoadFont(fileName string, size int) (Font, error)
	//CreateTexture creates a rectangular, unicolored Texture from a file and returns it
	CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (Texture, error)
	//Window returns a Window-instance, which holds information about the game's window
	Window() Window
}
