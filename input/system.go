package input

import "github.com/GomeBox/gome/adapters/input"

type System interface {
	Keyboard() Keyboard
	Update() error
}

type Adapters struct {
	Keyboard input.KeyboardAdapter
}

func NewSystem(adapters Adapters) System {
	system := new(system)
	system.keyboard = newKeyboard(adapters.Keyboard)
	return system
}

type system struct {
	keyboard *keyboard
}

func (sys *system) Keyboard() Keyboard {
	return sys.keyboard
}

func (sys *system) Update() error {
	return sys.keyboard.Update()
}
