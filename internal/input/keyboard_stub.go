package input

import (
	"github.com/GomeBox/gome/internal/input/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type keyboardStub struct {
	CallCntUpdate int
}

func (k *keyboardStub) RegisterKey(keyType primitives.KeyType) (interfaces.Key, error) {
	panic("implement me")
}

func (k *keyboardStub) UnregisterKey(keyType primitives.KeyType) {
	panic("implement me")
}

func (k *keyboardStub) Update() error {
	k.CallCntUpdate++
	return nil
}
