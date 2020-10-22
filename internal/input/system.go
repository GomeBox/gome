package input

import (
	adapters "github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/input"
)

type Adapters struct {
	Keyboard adapters.KeyboardAdapter
}

func NewSystem(adapters Adapters) input.System {
	system := new(system)
	system.keyboard = newKeyboard(adapters.Keyboard)
	return system
}

type system struct {
	keyboard *keyboard
}

func (sys *system) Keyboard() input.Keyboard {
	return sys.keyboard
}

func (sys *system) Update() error {
	return sys.keyboard.Update()
}
