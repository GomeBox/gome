package gome

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal"
	"github.com/GomeBox/gome/internal/game"
)

type CreateAdapters func() adapters.System

//Start starts your game. You must provide the game instance itself and a function to create
//the hardware adapters. You can use a gome adapter package like sdl2 or provide your own adapters
func Start(yourGame interfaces.Game, createAdapters CreateAdapters) error {
	adapterSystem := createAdapters()
	return internal.Start(
		yourGame,
		adapterSystem,
		internal.NewSystemsFactory(adapterSystem),
		game.SingleThreadedLoop)
}
