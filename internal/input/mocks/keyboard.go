package mocks

import (
	"github.com/GomeBox/gome/internal/input/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type Keyboard struct {
	CallCntUpdate   int
	OnRegisterKey   func(keyType primitives.KeyType) (interfaces.Key, error)
	OnUnregisterKey func(keyType primitives.KeyType)
}

func (k *Keyboard) RegisterKey(keyType primitives.KeyType) (interfaces.Key, error) {
	if k.OnRegisterKey != nil {
		return k.OnRegisterKey(keyType)
	}
	return nil, nil
}

func (k *Keyboard) UnregisterKey(keyType primitives.KeyType) {
	if k.OnUnregisterKey != nil {
		k.OnUnregisterKey(keyType)
	}
}

func (k *Keyboard) Update() error {
	k.CallCntUpdate++
	return nil
}
