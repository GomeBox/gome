package mocks

import "github.com/GomeBox/gome/primitives"

type KeyboardAdapter struct {
	OnKeyPressed      func(keyType primitives.KeyType) (bool, error)
	CallCntKeyPressed int
}

func (k *KeyboardAdapter) KeyPressed(keyType primitives.KeyType) (bool, error) {
	k.CallCntKeyPressed++
	if k.OnKeyPressed != nil {
		return k.OnKeyPressed(keyType)
	}
	return false, nil
}
