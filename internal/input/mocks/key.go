package mocks

import "github.com/GomeBox/gome/primitives"

type Key struct {
	CallCntIsPressed  int
	CallCntWasPressed int
	OnKeyType         func() primitives.KeyType
}

func (k *Key) KeyType() primitives.KeyType {
	if k.OnKeyType != nil {
		return k.OnKeyType()
	}
	return primitives.KeyEsc
}

func (k *Key) IsPressed() bool {
	k.CallCntIsPressed++
	return true
}

func (k *Key) WasPressed() bool {
	k.CallCntWasPressed++
	return false
}
