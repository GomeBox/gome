package interfaces

import "github.com/GomeBox/gome/input"

type System interface {
	Keyboard() input.Keyboard
	Update() error
}
