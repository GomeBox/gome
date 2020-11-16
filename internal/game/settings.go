package game

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/graphics"
)

//Settings contains all settings to run a game
type settings struct {
	windowSettings graphics.WindowSettings
}

//WindowSettings returns the settings of the game window
func (s settings) WindowSettings() interfaces.WindowSettings {
	return &s.windowSettings
}

//newSettings initializes new Settings with default values
func newSettings() *settings {
	s := new(settings)
	s.windowSettings = graphics.WindowSettings{}
	s.windowSettings.SetResolution(800, 600)
	return s
}
