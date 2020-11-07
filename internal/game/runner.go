package game

import (
	"github.com/GomeBox/gome/internal/game/interfaces"
)

type runner struct {
	initCallback   interfaces.InitializeCallback
	getSettings    interfaces.GetSettings
	updateCallback interfaces.UpdateCallback
	drawCallback   interfaces.DrawCallback
	gameSystem     interfaces.System
	loop           loop
}

func NewRunner(callbacks *interfaces.Callbacks) (interfaces.Runner, error) {
	adapterSystem, err := callbacks.CreateAdapters()
	if err != nil {
		return nil, err
	}
	systemsFactory := newSystemsFactory(adapterSystem)
	gameSystem := newSystem(adapterSystem, systemsFactory)
	runner := runner{
		initCallback:   callbacks.Init,
		getSettings:    callbacks.GetSettings,
		updateCallback: callbacks.Update,
		drawCallback:   callbacks.Draw,
		gameSystem:     gameSystem,
		loop:           singleThreadedLoop,
	}
	return &runner, nil
}

func (runner *runner) Run() error {
	settings := runner.getSettings()
	err := initialize(runner.gameSystem, settings, runner.initCallback)
	if err != nil {
		return err
	}
	return run(runner.gameSystem, runner.updateCallback, runner.drawCallback, runner.loop)
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
	update interfaces.UpdateCallback,
	draw interfaces.DrawCallback,
	loop loop) error {
	loopData := createLoopData(gameSystem, update, draw)
	return loop(loopData)
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
