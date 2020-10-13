package gome

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/internal/core"
)

func Run(game Game, settings Settings) error {
	g := new(core.Game)
	c := newContextWrapper(g)
	update := func() error {
		return game.Update(c)
	}
	draw := func() error {
		return game.Draw(c)
	}
	err := g.Initialize(game.Initialize, settings)
	if err != nil {
		return err
	}
	errChan := g.Loop(update, draw)
	err = <-errChan
	if err != nil {
		return err
	}
	return nil
}

type Settings struct {
	windowSettings graphics.WindowSettings
}

func (s Settings) WindowSettings() *graphics.WindowSettings {
	return &s.windowSettings
}

func NewSettings() Settings {
	s := new(Settings)
	s.windowSettings = *new(graphics.WindowSettings)
	s.windowSettings.Resolution.Width = 800
	s.windowSettings.Resolution.Height = 600
	return *s
}
