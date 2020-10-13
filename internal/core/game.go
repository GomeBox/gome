package core

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
)

type Game struct {
	quit          bool
	Error         chan error
	adapterSystem adapters.System
}

type UpdateCallBack func() error
type DrawCallback func() error
type InitializeCallback func() (adapters.System, error)

func (g *Game) Initialize(initialize InitializeCallback, settings Settings) error {
	var err error
	a, err := initialize()
	if err != nil {
		return err
	}
	err = checkAdapters(a, g)
	if err != nil {
		return err
	}
	err = a.Initialize()
	if err != nil {
		return err
	}
	err = g.adapterSystem.Graphics().ShowWindow(settings.WindowSettings())
	if err != nil {
		return err
	}
	g.quit = false
	return nil
}

func checkAdapters(adapters adapters.System, g *Game) error {
	if adapters.Graphics() == nil {
		return errors.New("gome.checkAdapters(): Graphics adapter is nil")
	}
	if adapters.Input() == nil {
		return errors.New("gome.checkAdapters(): input adapter is nil")
	}
	g.adapterSystem = adapters
	return nil
}

func (g *Game) Loop(update UpdateCallBack, draw DrawCallback) chan error {
	errChan := make(chan error, 1)
	var err error
	go func() {
		for !g.quit {
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

func (g *Game) AdapterSystem() adapters.System {
	return g.adapterSystem
}

func (g *Game) Quit() {
	g.quit = true
}
