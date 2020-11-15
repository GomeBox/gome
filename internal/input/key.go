package input

import "github.com/GomeBox/gome/primitives"

type Key interface {
	KeyType() primitives.KeyType
	IsPressed() bool
	WasPressed() bool
}

type key struct {
	keyType    primitives.KeyType
	isPressed  bool
	wasPressed bool
}

func (key *key) KeyType() primitives.KeyType {
	return key.keyType
}

func (key *key) IsPressed() bool {
	return key.isPressed
}

func (key *key) WasPressed() bool {
	return key.wasPressed
}
