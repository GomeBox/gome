package graphics

import (
	adapters "github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Adapters struct {
	TextureLoader  adapters.TextureLoader
	TextureCreator adapters.TextureCreator
	FontLoader     adapters.FontLoader
	WindowAdapter  adapters.WindowAdapter
}

func NewSystem(adapters Adapters) *System {
	sys := new(System)
	sys.textureLoader = adapters.TextureLoader
	sys.fontLoader = adapters.FontLoader
	sys.textureCreator = adapters.TextureCreator
	sys.windowAdapter = adapters.WindowAdapter
	window := newWindow(adapters.WindowAdapter)
	sys.window = window
	return sys
}

type System struct {
	textureCreator adapters.TextureCreator
	textureLoader  adapters.TextureLoader
	fontLoader     adapters.FontLoader
	windowAdapter  adapters.WindowAdapter
	window         graphics.Window
}

func (sys *System) LoadTexture(filename string) (graphics.Texture, error) {
	drawer, err := sys.textureLoader.Load(filename)
	if err != nil {
		return nil, err
	}
	return newTexture(drawer), nil
}

func (sys *System) LoadFont(fileName string, size int) (graphics.Font, error) {
	drawer, err := sys.fontLoader.Load(fileName, size)
	if err != nil {
		return nil, err
	}
	return newFont(drawer), nil
}

func (sys *System) CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (graphics.Texture, error) {
	drawer, err := sys.textureCreator.Create(dimensions, color)
	if err != nil {
		return nil, err
	}
	return newTexture(drawer), nil
}

func (sys *System) Window() graphics.Window {
	return sys.window
}
