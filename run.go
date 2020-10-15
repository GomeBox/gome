package gome

import (
	"errors"
	"github.com/GomeBox/gome/internal/core"
)

//Gome is used to run a Game
type Gome interface {
	//Run starts the Game. Returns when the game ends
	Run(game Game, settings Settings) error
	//ChangeGameRunner is used for testing purposes only
	ChangeGameRunner(runner core.GameRunner)
}

//New returns a new instance of Gome
func New() Gome {
	g := new(gome)
	g.gameRunner = core.CreateGameRunner()
	return g
}

type gome struct {
	gameRunner core.GameRunner
}

func (gome *gome) Run(game Game, settings Settings) error {
	if gome.gameRunner.Running() {
		return errors.New("game is already running")
	}
	err := gome.gameRunner.Initialize(game.Initialize, settings)
	if err != nil {
		return err
	}
	c := newContextWrapper(gome.gameRunner)
	update := func(timeDelta float32) error {
		return game.Update(timeDelta, c)
	}
	draw := func(timeDelta float32) error {
		return game.Draw(timeDelta, c)
	}
	err = gome.gameRunner.Loop(update, draw)
	//err = <-errChan
	if err != nil {
		return err
	}
	return nil
}

func (gome *gome) ChangeGameRunner(runner core.GameRunner) {
	gome.gameRunner = runner
}
