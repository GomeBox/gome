package game

import "github.com/GomeBox/gome/adapters/graphics"

//Settings contains all settings to run a game
type Settings struct {
	windowSettings graphics.WindowSettings
}

//WindowSettings returns the settings of the game window
func (s Settings) WindowSettings() *graphics.WindowSettings {
	return &s.windowSettings
}

//NewSettings initializes new Settings with default values
func NewSettings() Settings {
	s := new(Settings)
	s.windowSettings = *new(graphics.WindowSettings)
	s.windowSettings.Resolution.Width = 800
	s.windowSettings.Resolution.Height = 600
	return *s
}
