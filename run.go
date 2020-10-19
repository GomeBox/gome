package gome

import (
	"errors"
	"github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/internal/core"
)

//Gome is used to run a Interface
type Gome interface {
	//Run starts the Interface. Returns when the game ends
	Run(game game.Interface, settings game.Settings) error
	//ChangeGameRunner is used for testing purposes only
	ChangeGameRunner(runner core.GameRunner)
}

//New returns a new instance of Gome
func New() Gome {
	g := new(gome)
	g.gameRunner = core.NewGameRunner()
	return g
}

type gome struct {
	gameRunner core.GameRunner
}

//Run starts a game and starts the game loop. Returns, when the game is ended
func (gome *gome) Run(game game.Interface, settings game.Settings) error {
	if gome.gameRunner.Running() {
		return errors.New("game is already running")
	}
	err := gome.gameRunner.Init(game.CreateAdapters, settings)
	if err != nil {
		return err
	}
	err = gome.gameRunner.Loop(game.Initialize, game.Update, game.Draw)
	if err != nil {
		return err
	}
	return nil
}

func (gome *gome) ChangeGameRunner(runner core.GameRunner) {
	gome.gameRunner = runner
}
