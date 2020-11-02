package input

import "github.com/GomeBox/gome/internal/input/interfaces"

//System gives access to the input devices
type System interface {
	//Keyboard returns the Keyboard
	Keyboard() Keyboard
}

func NewSystem(internalSystem interfaces.System) System {
	return &system{internalSystem: internalSystem}
}

type system struct {
	internalSystem interfaces.System
}

func (sys *system) Keyboard() Keyboard {
	k := sys.internalSystem.Keyboard()
	return &keyboard{internalKeyboard: k}
}
