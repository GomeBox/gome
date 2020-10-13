package core

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
)

type UpdateCallBack func() error
type DrawCallback func() error
type InitializeCallback func() (adapters.System, error)

type GameRunner interface {
	Initialize(initialize InitializeCallback, settings Settings) error
	Loop(update UpdateCallBack, draw DrawCallback) chan error
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
	err = runner.adapterSystem.Graphics().ShowWindow(settings.WindowSettings())
	if err != nil {
		return err
	}
	runner.quit = false
	return nil
}

func checkAdapters(adapters adapters.System, g *gameRunner) error {
	if adapters.Graphics() == nil {
		return errors.New("gome.checkAdapters(): Graphics adapter is nil")
	}
	if adapters.Input() == nil {
		return errors.New("gome.checkAdapters(): input adapter is nil")
	}
	g.adapterSystem = adapters
	return nil
}

func (runner *gameRunner) Loop(update UpdateCallBack, draw DrawCallback) chan error {
	runner.running = true
	defer func() { runner.running = false }()
	errChan := make(chan error, 1)
	var err error
	go func() {
		for !runner.quit {
			err = runner.adapterSystem.Update()
			if err != nil {
				break
			}
			err = update()
			if err != nil {
				break
			}
			err = draw()
			if err != nil {
				break
			}
		}
		errChan <- err
	}()
	return errChan
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
