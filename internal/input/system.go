package input

import (
	adapters "github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/internal/input/interfaces"
)

type Adapters struct {
	Keyboard adapters.KeyboardAdapter
}

func NewSystem(adapters Adapters) interfaces.System {
	system := new(system)
	system.keyboard = newKeyboard(adapters.Keyboard)
	return system
}

type system struct {
	keyboard keyboard
}

func (sys *system) Keyboard() interfaces.Keyboard {
	return sys.keyboard
}

func (sys *system) Update() error {
	return sys.keyboard.Update()
}
