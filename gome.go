package gome

import (
	"github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/internal"
	internalGame "github.com/GomeBox/gome/internal/game"
)

//Gome is used to run a game
type Gome interface {
	//Run starts the interfaces. Returns when the game ends
	Run(game game.Interface) error
	Settings() game.Settings
}

//New returns a new instance of Gome
func New() Gome {
	g := internal.NewGome(internalGame.NewRunner())
	return g
}
