package game

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/internal/game/interfaces"
	"time"
)

func NewRunner() *runner {
	return new(runner)
}

type runner struct {
	running       bool
	quit          bool
	adapterSystem adapters.System
	settings      game.Settings
}

func (runner *runner) Init(createAdapters interfaces.CreateAdapters, settings game.Settings) error {
	var err error
	a, err := createAdapters()
	if err != nil {
		return err
	}
	err = checkAdapters(a, runner)
	if err != nil {
		return err
	}
	err = a.Initialize()
	if err != nil {
		return err
	}
	runner.settings = settings
	runner.quit = false
	return nil
}

func checkAdapters(adapters adapters.System, runner *runner) error {
	if adapters.Graphics() == nil {
		return errors.New("gome.checkAdapters(): Graphics adapter is nil")
	}
	if adapters.Input() == nil {
		return errors.New("gome.checkAdapters(): input adapter is nil")
	}
	runner.adapterSystem = adapters
	return nil
}

var elapsedTime float32
var frameStart time.Time

func (runner *runner) Loop(initialize interfaces.InitializeCallback, update interfaces.UpdateCallback, draw interfaces.DrawCallback) error {
	runner.running = true
	defer func() { runner.running = false }()
	var err error
	context, err := runner.onInitialize(initialize)
	if err != nil {
		return err
	}
	for !runner.quit {
		frameStart = time.Now()
		err = runner.adapterSystem.Graphics().ScreenPresenter().Clear()
		if err != nil {
			break
		}
		err = context.Update()
		if err != nil {
			break
		}
		err = update(elapsedTime, context)
		if err != nil {
			break
		}
		err = draw(elapsedTime, context)
		if err != nil {
			break
		}
		err = runner.adapterSystem.Graphics().ScreenPresenter().Present()
		if err != nil {
			break
		}
		//TODO: Framecount should be customizable
		time.Sleep(time.Millisecond)
		elapsedTime = float32(time.Since(frameStart).Milliseconds()) // / 1000000.0
		if elapsedTime <= 0 {
			elapsedTime = 1.0
		}
	}
	return err
}

func (runner *runner) onInitialize(callback interfaces.InitializeCallback) (*context, error) {
	err := runner.adapterSystem.Graphics().WindowAdapter().OpenWindow(runner.settings.WindowSettings())
	if err != nil {
		return nil, err
	}
	systemsFactory := newSystemsFactory(runner.adapterSystem)
	system := NewSystem(runner.adapterSystem, systemsFactory)
	context := newContext(runner, system)
	err = callback(context)
	if err != nil {
		return nil, err
	}
	return context, nil
}

func (runner *runner) Quit() {
	runner.quit = true
}

func (runner *runner) Running() bool {
	return runner.running
}
