package interfaces

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/game"
)

type InitializeCallback func(context game.Context) error
type UpdateCallback func(timeDelta float32, context game.Context) (keepRunning bool, error error)
type DrawCallback func(timeDelta float32, context game.Context) error
type CreateAdapters func() (adapters.System, error)
type GetSettings func() game.Settings

type Runner interface {
	Init(createAdapters CreateAdapters, settings game.Settings) error
	Loop(initialize InitializeCallback, update UpdateCallback, draw DrawCallback) error
	Running() bool
	Quit()
}
