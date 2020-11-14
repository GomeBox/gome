package gome

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal"
)

type CreateAdapters func() (adapters.System, error)

//Start starts your game. You must provide the game instance itself and a function to create
//the hardware adapters. You can use a gome adapter package like sdl2 or provide your own adapters
func Start(yourGame interfaces.Game, createAdapters CreateAdapters) error {
	a, err := createAdapters()
	if err != nil {
		return err
	}
	return internal.Run(yourGame, a)
}

//type gome struct {
//	impl        *internal.GomeImpl
//	gameContext Context
//}

//func (gome *gome) Run(game interfaces2.Game) error {
//	callbacks := interfaces.Callbacks{
//		Init: func(gameSystem interfaces.System) error {
//			gome.gameContext = internal.newContext(gameSystem)
//			return game.Initialize(gome.gameContext)
//		},
//		Update: func(timeDelta float32) (keepRunning bool, error error) {
//			return game.Update(timeDelta, gome.gameContext)
//		},
//		Draw: func(timeDelta float32) error {
//			return game.Draw(timeDelta, gome.gameContext)
//		},
//		CreateAdapters: game.CreateAdapters,
//		GetSettings: func() interfaces.Settings {
//			return gome.impl.Settings()
//		},
//	}
//	return gome.impl.Run(&callbacks)
//}
//
//func (gome *gome) Settings() interfaces2.Settings {
//	return gome.impl.Settings()
//}
