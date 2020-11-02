package interfaces

import (
	"github.com/GomeBox/gome/adapters"
)

type Callbacks struct {
	Init           InitializeCallback
	Update         UpdateCallback
	Draw           DrawCallback
	CreateAdapters CreateAdapters
	GetSettings    GetSettings
}

type InitializeCallback func(gameSystem System) error
type UpdateCallback func(timeDelta float32) (keepRunning bool, error error)
type DrawCallback func(timeDelta float32) error
type CreateAdapters func() (adapters.System, error)
type GetSettings func() Settings
