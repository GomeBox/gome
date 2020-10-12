package game

import "github.com/GomeBox/gome/adapters/graphics"

func Run(instance Interface, settings Settings) error {
	g := new(game)
	err := g.initialize(instance, settings)
	if err != nil {
		return err
	}
	errChan := g.loop(instance)
	err = <-errChan
	if err != nil {
		return err
	}
	return nil
}

type Settings struct {
	WindowSettings graphics.WindowSettings
}

func NewSettings() Settings {
	s := new(Settings)
	s.WindowSettings = *new(graphics.WindowSettings)
	s.WindowSettings.Resolution.Width = 800
	s.WindowSettings.Resolution.Height = 600
	return *s
}
