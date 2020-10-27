package internal

import (
	"errors"
	"github.com/GomeBox/gome/game"
	internalGame "github.com/GomeBox/gome/internal/game"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

func NewGome(runner interfaces.Runner) *GomeImpl {
	g := new(GomeImpl)
	g.gameRunner = runner
	g.gameSettings = internalGame.NewSettings()
	return g
}

type GomeImpl struct {
	gameRunner   interfaces.Runner
	gameSettings game.Settings
}

//Run starts a game and starts the game loop. Returns, when the game is ended
func (gome *GomeImpl) Run(game game.Interface) error {
	if gome.gameRunner.Running() {
		return errors.New("game is already running")
	}
	err := gome.gameRunner.Init(game.CreateAdapters, gome.gameSettings)
	if err != nil {
		return err
	}
	err = gome.gameRunner.Loop(game.Initialize, game.Update, game.Draw)
	if err != nil {
		return err
	}
	return nil
}

func (gome *GomeImpl) Settings() game.Settings {
	return gome.gameSettings
}
