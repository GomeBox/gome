package mocks

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Graphics struct {
	OnOpenWindow func(settings graphics.WindowSettings) error
}

func (g *Graphics) OpenWindow(settings graphics.WindowSettings) error {
	if g.OnOpenWindow != nil {
		return g.OnOpenWindow(settings)
	}
	return nil
}

func (g *Graphics) Clear() error {
	panic("implement me")
}

func (g *Graphics) Present() error {
	panic("implement me")
}

func (g *Graphics) LoadTexture(fileName string) (interfaces.Texture, error) {
	panic("implement me")
}

func (g *Graphics) LoadFont(fileName string, size int) (interfaces.Font, error) {
	panic("implement me")
}

func (g *Graphics) CreateTexture(dimensions primitives.Dimensions, color primitives.Color) (interfaces.Texture, error) {
	panic("implement me")
}

func (g *Graphics) Window() interfaces.Window {
	panic("implement me")
}
