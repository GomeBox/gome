package graphics

import (
	adapters "github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type Adapters struct {
	TextureLoader   adapters.TextureLoader
	TextureCreator  adapters.TextureCreator
	FontLoader      adapters.FontLoader
	WindowAdapter   adapters.WindowAdapter
	ScreenPresenter adapters.ScreenPresenter
}

func NewSystem(adapters Adapters) interfaces.System {
	sys := new(system)
	sys.textureLoader = adapters.TextureLoader
	sys.fontLoader = adapters.FontLoader
	sys.textureCreator = adapters.TextureCreator
	sys.windowAdapter = adapters.WindowAdapter
	window := newWindow(adapters.WindowAdapter)
	sys.window = window
	sys.screenPresenter = adapters.ScreenPresenter
	return sys
}

type system struct {
	textureCreator  adapters.TextureCreator
	textureLoader   adapters.TextureLoader
	fontLoader      adapters.FontLoader
	windowAdapter   adapters.WindowAdapter
	window          interfaces.Window
	screenPresenter adapters.ScreenPresenter
}

func (sys *system) LoadTexture(filename string) (interfaces.Texture, error) {
	drawer, err := sys.textureLoader.Load(filename)
	if err != nil {
		return nil, err
	}
	return newTexture(drawer), nil
}

func (sys *system) LoadFont(fileName string, size int) (interfaces.Font, error) {
	drawer, err := sys.fontLoader.Load(fileName, size)
	if err != nil {
		return nil, err
	}
	return newFont(drawer), nil
}

func (sys *system) CreateTexture(dimensions *primitives.Dimensions, color *primitives.Color) (interfaces.Texture, error) {
	drawer, err := sys.textureCreator.Create(dimensions, color)
	if err != nil {
		return nil, err
	}
	return newTexture(drawer), nil
}

func (sys *system) Window() interfaces.Window {
	return sys.window
}

func (sys *system) Clear() error {
	return sys.screenPresenter.Clear()
}

func (sys *system) Present() error {
	return sys.screenPresenter.Present()
}
