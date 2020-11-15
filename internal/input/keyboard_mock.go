package input

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type KeyboardMock struct {
	CallCntUpdate   int
	OnRegisterKey   func(keyType primitives.KeyType) (interfaces.Key, error)
	OnUnregisterKey func(keyType primitives.KeyType)
}

func (k *KeyboardMock) RegisterKey(keyType primitives.KeyType) (interfaces.Key, error) {
	if k.OnRegisterKey != nil {
		return k.OnRegisterKey(keyType)
	}
	return nil, nil
}

func (k *KeyboardMock) UnregisterKey(keyType primitives.KeyType) {
	if k.OnUnregisterKey != nil {
		k.OnUnregisterKey(keyType)
	}
}

func (k *KeyboardMock) Update() error {
	k.CallCntUpdate++
	return nil
}
