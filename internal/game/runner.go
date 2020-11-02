package game

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

func Run(callbacks *interfaces.Callbacks, settings interfaces.Settings) error {
	adapterSystem, err := callbacks.CreateAdapters()
	if err != nil {
		return err
	}
	gameSystem := createGameSystem(adapterSystem)
	err = initialize(gameSystem, settings, callbacks.Init)
	if err != nil {
		return err
	}
	return run(gameSystem, callbacks, singleThreadedLoop)
}

func createGameSystem(adapterSystem adapters.System) interfaces.System {
	systemsFactory := newSystemsFactory(adapterSystem)
	return newSystem(adapterSystem, systemsFactory)
}

func initialize(gameSystem interfaces.System,
	settings interfaces.Settings,
	initCallback interfaces.InitializeCallback) error {
	err := gameSystem.Initialize()
	if err != nil {
		return err
	}
	err = gameSystem.Graphics().Window().Open(settings.WindowSettings())
	if err != nil {
		return err
	}
	err = initCallback(gameSystem)
	if err != nil {
		return err
	}
	return nil
}

func run(gameSystem interfaces.System,
	callbacks *interfaces.Callbacks,
	loopFunc loop) error {
	loopData := createLoopData(gameSystem, callbacks.Update, callbacks.Draw)
	return loopFunc(loopData)
}

func createLoopData(gameSystem interfaces.System,
	updateCallback interfaces.UpdateCallback,
	drawCallback interfaces.DrawCallback) *loopData {
	return &loopData{
		gameSystem: gameSystem,
		update: func(timeDelta float32) (keepRunning bool, error error) {
			return update(updateCallback, gameSystem, timeDelta)
		},
		draw: func(timeDelta float32) error { return drawCallback(timeDelta) },
	}
}

func update(updateCallback interfaces.UpdateCallback,
	gameSystem interfaces.System,
	timeDelta float32) (keepRunning bool, err error) {
	err = gameSystem.Update()
	if err != nil {
		return false, err
	}
	keepRunning, err = updateCallback(timeDelta)
	if err != nil {
		return false, err
	}
	return keepRunning, nil
}
