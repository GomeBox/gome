package input

import "github.com/GomeBox/gome/primitives"

type KeyMock struct {
	CallCntIsPressed  int
	CallCntWasPressed int
	OnKeyType         func() primitives.KeyType
}

func (k *KeyMock) KeyType() primitives.KeyType {
	if k.OnKeyType != nil {
		return k.OnKeyType()
	}
	return primitives.KeyEsc
}

func (k *KeyMock) IsPressed() bool {
	k.CallCntIsPressed++
	return true
}

func (k *KeyMock) WasPressed() bool {
	k.CallCntWasPressed++
	return false
}
