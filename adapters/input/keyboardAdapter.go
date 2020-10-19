package input

import "github.com/GomeBox/gome/primitives"

//Keyboard gives access to the computers keyboard
type KeyboardAdapter interface {
	KeyPressed(keyType primitives.KeyType) (bool, error)
}
