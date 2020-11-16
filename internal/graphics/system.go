package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/interfaces"
	gameGraphics "github.com/GomeBox/gome/internal/game/graphics"
	"github.com/GomeBox/gome/primitives"
)

func NewSystem(graphicsAdapters graphics.Adapters) gameGraphics.System {
	sys := new(system)
	sys.textureLoader = graphicsAdapters.TextureLoader()
	sys.fontLoader = graphicsAdapters.FontLoader()
	sys.textureCreator = graphicsAdapters.TextureCreator()
	sys.windowAdapter = graphicsAdapters.WindowAdapter()
	window := newWindow(graphicsAdapters.WindowAdapter())
	sys.window = window
	sys.screenPresenter = graphicsAdapters.ScreenPresenter()
	return sys
}

type system struct {
	textureCreator  graphics.TextureCreator
	textureLoader   graphics.TextureLoader
	fontLoader      graphics.FontLoader
	windowAdapter   graphics.WindowAdapter
	window          *window
	screenPresenter graphics.ScreenPresenter
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

func (sys *system) CreateTexture(dimensions primitives.Dimensions, color primitives.Color) (interfaces.Texture, error) {
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

func (sys *system) OpenWindow(settings gameGraphics.WindowSettings) error {
	return sys.window.Open(settings)
}
