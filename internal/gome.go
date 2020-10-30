package internal

import (
	"errors"
	"github.com/GomeBox/gome/game"
	internalGame "github.com/GomeBox/gome/internal/game"
)

func NewGome() *GomeImpl {
	g := new(GomeImpl)
	g.gameSettings = internalGame.NewSettings()
	return g
}

type GomeImpl struct {
	gameSettings game.Settings
	running      bool
}

//Run starts a game and starts the game loop. Returns, when the game is ended
func (gome *GomeImpl) Run(game game.Interface) error {
	if gome.running {
		return errors.New("game is already running")
	}
	gome.running = true
	defer func() { gome.running = false }()
	gameCallbacks := internalGame.Callbacks{
		OnInitialize: game.Initialize,
		OnUpdate:     game.Update,
		OnDraw:       game.Draw,
	}
	adapterSystem, err := game.CreateAdapters()
	if err != nil {
		return err
	}
	return internalGame.Run(gameCallbacks, adapterSystem, gome.gameSettings)
}

func (gome *GomeImpl) Settings() game.Settings {
	return gome.gameSettings
}
