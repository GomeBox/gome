package core

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
	"time"
)

type InitializeCallback func() error
type UpdateCallback func(timeDelta float32) error
type DrawCallback func(timeDelta float32) error
type CreateAdapters func() (adapters.System, error)

type GameRunner interface {
	Init(createAdapters CreateAdapters, settings Settings) error
	Loop(initialize InitializeCallback, update UpdateCallback, draw DrawCallback) error
	Running() bool
	AdapterSystem() adapters.System
	Quit()
}

func NewGameRunner() GameRunner {
	return new(gameRunner)
}

type gameRunner struct {
	running       bool
	quit          bool
	adapterSystem adapters.System
	settings      Settings
}

func (runner *gameRunner) Init(createAdapters CreateAdapters, settings Settings) error {
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

func checkAdapters(adapters adapters.System, runner *gameRunner) error {
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

func (runner *gameRunner) Loop(initialize InitializeCallback, update UpdateCallback, draw DrawCallback) error {
	runner.running = true
	defer func() { runner.running = false }()
	var err error
	err = runner.onInitialize(initialize)
	if err != nil {
		return err
	}
	for !runner.quit {
		frameStart = time.Now()
		err = runner.adapterSystem.Graphics().ScreenPresenter().Clear()
		if err != nil {
			break
		}
		err = runner.adapterSystem.Update()
		if err != nil {
			break
		}
		err = update(elapsedTime)
		if err != nil {
			break
		}
		err = draw(elapsedTime)
		if err != nil {
			break
		}
		err = runner.adapterSystem.Graphics().ScreenPresenter().Present()
		if err != nil {
			break
		}
		elapsedTime = float32(time.Since(frameStart).Nanoseconds()) / 1000000.0
		if elapsedTime == 0.0 {
			elapsedTime = 1.0
		}
	}
	return err
}

func (runner *gameRunner) onInitialize(callback InitializeCallback) error {
	err := runner.adapterSystem.Graphics().WindowAdapter().ShowWindow(runner.settings.WindowSettings())
	if err != nil {
		return err
	}
	err = callback()
	if err != nil {
		return err
	}
	return nil
}

func (runner *gameRunner) AdapterSystem() adapters.System {
	return runner.adapterSystem
}

func (runner *gameRunner) Quit() {
	runner.quit = true
}

func (runner *gameRunner) Running() bool {
	return runner.running
}
