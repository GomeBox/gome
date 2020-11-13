package input

import (
	"github.com/GomeBox/gome/internal/input/interfaces"
	"github.com/GomeBox/gome/primitives"
)

//Key represents a single key of the keyboard
type Key interface {
	//KeyType defines which key is represented
	KeyType() primitives.KeyType
	//IsPressed is true if the key was pressed during the current run of update
	IsPressed() bool
	//WasPressed is true if the key was pressed during the last run of update
	WasPressed() bool
}

type key struct {
	internalKey interfaces.Key
}

func (k *key) KeyType() primitives.KeyType {
	return k.internalKey.KeyType()
}

func (k *key) IsPressed() bool {
	return k.internalKey.IsPressed()
}

func (k *key) WasPressed() bool {
	return k.internalKey.WasPressed()
}
