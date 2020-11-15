package game

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/graphics"
)

//Settings contains all settings to run a game
type settings struct {
	windowSettings *graphics.windowSettings
}

//WindowSettings returns the settings of the game window
func (s settings) WindowSettings() interfaces.WindowSettings {
	return s.windowSettings
}

//NewSettings initializes new Settings with default values
func NewSettings() interfaces.Settings {
	s := new(settings)
	s.windowSettings = new(graphics.windowSettings)
	s.windowSettings.Resolution(800, 600)
	return *s
}
