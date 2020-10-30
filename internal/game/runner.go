package game

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

type Callbacks struct {
	OnInitialize interfaces.InitializeCallback
	OnUpdate     interfaces.UpdateCallback
	OnDraw       interfaces.DrawCallback
}

type initializeType func(adapterSystem adapters.System,
	context *context,
	initCallback interfaces.InitializeCallback) error

type runType func(adapterSystem adapters.System,
	settings game.Settings,
	context *context,
	callbacks *Callbacks,
	loopFunc loop) error

func Run(callbacks Callbacks, adapterSystem adapters.System, settings game.Settings) error {
	return initAndRun(&callbacks, adapterSystem, settings, singleThreadedLoop, initialize, run)
}

func initAndRun(
	callbacks *Callbacks,
	adapterSystem adapters.System,
	settings game.Settings,
	loop loop,
	init initializeType,
	run runType) error {
	context := createContext(adapterSystem)
	err := init(adapterSystem, context, callbacks.OnInitialize)
	if err != nil {
		return err
	}
	return run(adapterSystem, settings, context, callbacks, loop)
}

func createContext(adapterSystem adapters.System) *context {
	systemsFactory := newSystemsFactory(adapterSystem)
	system := NewSystem(adapterSystem, systemsFactory)
	return newContext(system)
}

func createLoopData(
	screenPresenter graphics.ScreenPresenter,
	context *context,
	updateCallback interfaces.UpdateCallback,
	drawCallback interfaces.DrawCallback) *loopData {
	return &loopData{
		screenPresenter: screenPresenter,
		update: func(timeDelta float32) (keepRunning bool, error error) {
			return update(updateCallback, context, timeDelta)
		},
		draw: func(timeDelta float32) error { return drawCallback(timeDelta, context) },
	}
}

func initialize(
	adapterSystem adapters.System,
	context *context,
	initCallback interfaces.InitializeCallback) error {
	err := adapterSystem.Initialize()
	if err != nil {
		return err
	}
	err = initCallback(context)
	if err != nil {
		return err
	}
	return nil
}

func run(adapterSystem adapters.System,
	settings game.Settings,
	context *context,
	callbacks *Callbacks,
	loopFunc loop) error {
	err := adapterSystem.Graphics().WindowAdapter().OpenWindow(settings.WindowSettings())
	if err != nil {
		return err
	}
	screenPresenter := adapterSystem.Graphics().ScreenPresenter()
	loopData := createLoopData(screenPresenter, context, callbacks.OnUpdate, callbacks.OnDraw)
	return loopFunc(loopData)
}

func update(updateCallback interfaces.UpdateCallback, context *context, timeDelta float32) (keepRunning bool, err error) {
	err = context.Update()
	if err != nil {
		return false, err
	}
	keepRunning, err = updateCallback(timeDelta, context)
	if err != nil {
		return false, err
	}
	return keepRunning, nil
}
