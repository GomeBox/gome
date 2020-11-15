package input

import (
	"github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/interfaces"
	gameInput "github.com/GomeBox/gome/internal/game/input"
)

func NewSystem(inputAdapters input.Adapters) gameInput.System {
	system := new(system)
	system.keyboard = newKeyboard(inputAdapters.Keyboard())
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
