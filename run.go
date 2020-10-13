package gome

import (
	"errors"
	"github.com/GomeBox/gome/internal/core"
)

type Gome interface {
	Run(game Game, settings Settings) error
	ChangeGameRunner(runner core.GameRunner)
}

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
	update := func() error {
		return game.Update(c)
	}
	draw := func() error {
		return game.Draw(c)
	}
	errChan := gome.gameRunner.Loop(update, draw)
	err = <-errChan
	if err != nil {
		return err
	}
	return nil
}

func (gome *gome) ChangeGameRunner(runner core.GameRunner) {
	gome.gameRunner = runner
}
