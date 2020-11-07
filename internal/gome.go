package internal

import (
	"errors"
	"github.com/GomeBox/gome/internal/game"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

func NewGome() *GomeImpl {
	g := new(GomeImpl)
	g.gameSettings = game.NewSettings()
	return g
}

type GomeImpl struct {
	gameSettings interfaces.Settings
	running      bool
}

//Run starts a game and starts the game loop. Returns, when the game is ended
func (gome *GomeImpl) Run(callbacks *interfaces.Callbacks) error {
	if gome.running {
		return errors.New("game is already running")
	}
	gome.running = true
	defer func() { gome.running = false }()
	runner, err := game.NewRunner(callbacks)
	if err != nil {
		return err
	}
	return runner.Run()
}

func (gome *GomeImpl) Settings() interfaces.Settings {
	return gome.gameSettings
}
