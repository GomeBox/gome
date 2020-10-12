package game

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
)

type gome struct {
	quit          bool
	Error         chan error
	adapterSystem adapters.System
}

func (g *gome) initialize(instance Interface, settings Settings) error {
	var err error
	a, err := instance.Initialize()
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
	err = g.adapterSystem.Graphics().ShowWindow(settings.WindowSettings)
	if err != nil {
		return err
	}
	g.quit = false
	return nil
}

func checkAdapters(adapters adapters.System, g *gome) error {
	if adapters.Graphics() == nil {
		return errors.New("gome.checkAdapters(): Graphics adapter is nil")
	}
	if adapters.Input() == nil {
		return errors.New("gome.checkAdapters(): input adapter is nil")
	}
	g.adapterSystem = adapters
	return nil
}

func (g *gome) loop(instance Interface) chan error {
	context := newContext(g)
	errChan := make(chan error, 1)
	var err error
	go func() {
		for !g.quit {
			err = instance.Update(context)
			if err != nil {
				break
			}
			err = instance.Draw(context)
			if err != nil {
				break
			}
		}
		errChan <- err
	}()
	return errChan
}

func (g *gome) Quit() {
	g.quit = true
}
