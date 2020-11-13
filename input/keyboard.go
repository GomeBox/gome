package input

import (
	"github.com/GomeBox/gome/internal/input/interfaces"
	"github.com/GomeBox/gome/primitives"
)

//Keyboard contains functions to access the computer's keyboard
type Keyboard interface {
	//RegisterKey registers a key to track it's state and returns a Key instance
	RegisterKey(keyType primitives.KeyType) (Key, error)
	//UnregisterKey unregisters a key. The key-state will not be updated anymore
	UnregisterKey(keyType primitives.KeyType)
}

type keyboard struct {
	internalKeyboard interfaces.Keyboard
}

func (kb *keyboard) RegisterKey(keyType primitives.KeyType) (Key, error) {
	k, err := kb.internalKeyboard.RegisterKey(keyType)
	if err != nil {
		return nil, err
	}
	return &key{internalKey: k}, nil
}

func (kb *keyboard) UnregisterKey(keyType primitives.KeyType) {
	kb.internalKeyboard.UnregisterKey(keyType)
}
