package game

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/game"
)

//Settings contains all settings to run a game
type settings struct {
	windowSettings *graphics.WindowSettings
}

//WindowSettings returns the settings of the game window
func (s settings) WindowSettings() *graphics.WindowSettings {
	return s.windowSettings
}

//NewSettings initializes new Settings with default values
func NewSettings() game.Settings {
	s := new(settings)
	s.windowSettings = new(graphics.WindowSettings)
	s.windowSettings.Resolution.Width = 800
	s.windowSettings.Resolution.Height = 600
	return *s
}
