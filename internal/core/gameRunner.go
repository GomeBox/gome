package core

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
	"time"
)

type UpdateCallBack func(timeDelta float32) error
type DrawCallback func(timeDelta float32) error
type InitializeCallback func() (adapters.System, error)

type GameRunner interface {
	Initialize(initialize InitializeCallback, settings Settings) error
	Loop(update UpdateCallBack, draw DrawCallback) error
	Running() bool
	AdapterSystem() adapters.System
	Quit()
}

func CreateGameRunner() GameRunner {
	return new(gameRunner)
}

type gameRunner struct {
	running       bool
	quit          bool
	adapterSystem adapters.System
}

func (runner *gameRunner) Initialize(initialize InitializeCallback, settings Settings) error {
	var err error
	a, err := initialize()
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
	err = runner.adapterSystem.Graphics().WindowAdapter().ShowWindow(settings.WindowSettings())
	if err != nil {
		return err
	}
	runner.quit = false
	return nil
}

func checkAdapters(adapters adapters.System, runner *gameRunner) error {
	if adapters.Graphics == nil {
		return errors.New("gome.checkAdapters(): Graphics adapter is nil")
	}
	if adapters.Input == nil {
		return errors.New("gome.checkAdapters(): input adapter is nil")
	}
	runner.adapterSystem = adapters
	return nil
}

var elapsedTime float32
var frameStart time.Time

func (runner *gameRunner) Loop(update UpdateCallBack, draw DrawCallback) error {
	runner.running = true
	defer func() { runner.running = false }()
	var err error
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

func (runner *gameRunner) AdapterSystem() adapters.System {
	return runner.adapterSystem
}

func (runner *gameRunner) Quit() {
	runner.quit = true
}

func (runner *gameRunner) Running() bool {
	return runner.running
}
