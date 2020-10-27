package mocks

import (
	"github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

type Runner struct {
	OnQuit func()
}

func (r Runner) Init(createAdapters interfaces.CreateAdapters, settings game.Settings) error {
	return nil
}

func (r Runner) Loop(initialize interfaces.InitializeCallback, update interfaces.UpdateCallback, draw interfaces.DrawCallback) error {
	return nil
}

func (r Runner) Running() bool {
	return false
}

func (r Runner) Quit() {
	if r.OnQuit != nil {
		r.OnQuit()
	}
}
