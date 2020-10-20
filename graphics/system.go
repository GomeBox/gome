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
}

func NewSystem(adapters Adapters) System {
	sys := new(system)
	sys.textureLoader = adapters.TextureLoader
	sys.fontLoader = adapters.FontLoader
	sys.textureCreator = adapters.TextureCreator
	return sys
}

type system struct {
	textureCreator graphics.TextureCreator
	textureLoader  graphics.TextureLoader
	fontLoader     graphics.FontLoader
	windowAdapter  graphics.WindowAdapter
}

func (sys *system) LoadTexture(filename string) (Texture, error) {
	drawer, err := sys.textureLoader.Load(filename)
	if err != nil {
		return nil, err
	}
	return NewTexture(drawer), nil
}

func (sys *system) LoadFont(fileName string, size int) (Font, error) {
	drawer, err := sys.fontLoader.Load(fileName, size)
	if err != nil {
		return nil, err
	}
	return NewFont(drawer), nil
}

func (sys *system) CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (Texture, error) {
	drawer, err := sys.textureCreator.Create(dimensions, color)
	if err != nil {
		return nil, err
	}
	return NewTexture(drawer), nil
}
