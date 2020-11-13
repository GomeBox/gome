package gome

import (
	"github.com/GomeBox/gome/internal"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

//Gome is used to run a game
type Gome interface {
	//Run starts the interfaces. Returns when the game ends
	Run(game Interface) error
	Settings() Settings
}

//New returns a new instance of Gome
func New() Gome {
	g := internal.NewGome()
	return &gome{
		impl: g,
	}
}

type gome struct {
	impl        *internal.GomeImpl
	gameContext Context
}

func (gome *gome) Run(game Interface) error {
	callbacks := interfaces.Callbacks{
		Init: func(gameSystem interfaces.System) error {
			gome.gameContext = newContext(gameSystem)
			return game.Initialize(gome.gameContext)
		},
		Update: func(timeDelta float32) (keepRunning bool, error error) {
			return game.Update(timeDelta, gome.gameContext)
		},
		Draw: func(timeDelta float32) error {
			return game.Draw(timeDelta, gome.gameContext)
		},
		CreateAdapters: game.CreateAdapters,
		GetSettings: func() interfaces.Settings {
			return gome.impl.Settings()
		},
	}
	return gome.impl.Run(&callbacks)
}

func (gome *gome) Settings() Settings {
	return gome.impl.Settings()
}
