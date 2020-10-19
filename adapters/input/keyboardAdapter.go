package input

import "github.com/GomeBox/gome/primitives"

//KeyboardAdapter gives access to the computers keyboard
type KeyboardAdapter interface {
	//KeyPressed returns true, if the key is currently pressed
	KeyPressed(keyType primitives.KeyType) (bool, error)
}
