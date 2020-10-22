package input

import (
	adapters "github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/input"
)

type Adapters struct {
	Keyboard adapters.KeyboardAdapter
}

func NewSystem(adapters Adapters) *System {
	system := new(System)
	system.keyboard = newKeyboard(adapters.Keyboard)
	return system
}

type System struct {
	keyboard *keyboard
}

func (sys *System) Keyboard() input.Keyboard {
	return sys.keyboard
}

func (sys *System) Update() error {
	return sys.keyboard.Update()
}
