package gome

import (
	"github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/internal"
	"github.com/GomeBox/gome/internal/core"
)

//Gome is used to run a game
type Gome interface {
	//Run starts the Interface. Returns when the game ends
	Run(game game.Interface) error
	Settings() *game.Settings
}

//New returns a new instance of Gome
func New() Gome {
	g := new(internal.GomeImpl)
	g.GameRunner = core.NewGameRunner()
	g.GameSettings = game.NewSettings()
	return g
}
