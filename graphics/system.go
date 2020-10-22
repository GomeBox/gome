package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Adapters struct {
	TextureLoader  graphics.TextureLoader
	TextureCreator graphics.TextureCreator
	FontLoader     graphics.FontLoader
	WindowAdapter  graphics.WindowAdapter
}

type System interface {
	LoadTexture(fileName string) (Texture, error)
	LoadFont(fileName string, size int) (Font, error)
	CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (Texture, error)
	Window() Window
}

func NewSystem(adapters Adapters) (System, error) {
	sys := new(system)
	sys.textureLoader = adapters.TextureLoader
	sys.fontLoader = adapters.FontLoader
	sys.textureCreator = adapters.TextureCreator
	sys.windowAdapter = adapters.WindowAdapter
	window, err := newWindow(adapters.WindowAdapter)
	if err != nil {
		return nil, err
	}
	sys.window = window
	return sys, nil
}

type system struct {
	textureCreator graphics.TextureCreator
	textureLoader  graphics.TextureLoader
	fontLoader     graphics.FontLoader
	windowAdapter  graphics.WindowAdapter
	window         Window
}

func (sys *system) LoadTexture(filename string) (Texture, error) {
	drawer, err := sys.textureLoader.Load(filename)
	if err != nil {
		return nil, err
	}
	return newTexture(drawer), nil
}

func (sys *system) LoadFont(fileName string, size int) (Font, error) {
	drawer, err := sys.fontLoader.Load(fileName, size)
	if err != nil {
		return nil, err
	}
	return newFont(drawer), nil
}

func (sys *system) CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (Texture, error) {
	drawer, err := sys.textureCreator.Create(dimensions, color)
	if err != nil {
		return nil, err
	}
	return newTexture(drawer), nil
}

func (sys *system) Window() Window {
	return sys.window
}
