package gome

import "github.com/GomeBox/gome/adapters/graphics"

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
