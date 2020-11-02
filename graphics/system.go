package graphics

import (
	"github.com/GomeBox/gome/internal/graphics/interfaces"
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

func NewSystem(internalSystem interfaces.System) System {
	return &system{internalSystem: internalSystem}
}

type system struct {
	internalSystem interfaces.System
}

func (sys *system) LoadTexture(fileName string) (Texture, error) {
	tex, err := sys.internalSystem.LoadTexture(fileName)
	if err != nil {
		return nil, err
	}
	return &texture{internalTexture: tex}, nil
}

func (sys *system) LoadFont(fileName string, size int) (Font, error) {
	f, err := sys.internalSystem.LoadFont(fileName, size)
	if err != nil {
		return nil, err
	}
	return &font{internalFont: f}, nil
}

func (sys *system) CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (Texture, error) {
	tex, err := sys.internalSystem.CreateTexture(dimensions, color)
	if err != nil {
		return nil, err
	}
	return &texture{internalTexture: tex}, nil
}

func (sys *system) Window() Window {
	win := sys.internalSystem.Window()
	return &window{internalWindow: win}
}
