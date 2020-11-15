package input

import (
	adapters "github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/primitives"
)

type Keyboard interface {
	RegisterKey(keyType primitives.KeyType) (Key, error)
	UnregisterKey(keyType primitives.KeyType)
}

func newKeyboard(adapter adapters.KeyboardAdapter) *keyboardImpl {
	keyboard := new(keyboardImpl)
	keyboard.adapter = adapter
	keyboard.registeredKeys = make(map[primitives.KeyType]*key)
	return keyboard
}

type keyboard interface {
	RegisterKey(keyType primitives.KeyType) (Key, error)
	UnregisterKey(keyType primitives.KeyType)
	Update() error
}

type keyboardImpl struct {
	adapter        adapters.KeyboardAdapter
	registeredKeys map[primitives.KeyType]*key
}

func (keyboard *keyboardImpl) RegisterKey(keyType primitives.KeyType) (Key, error) {
	pressed, err := keyboard.adapter.KeyPressed(keyType)
	if err != nil {
		return nil, err
	}
	var keyInstance *key
	keyInstance, ok := keyboard.registeredKeys[keyType]
	if !ok {
		keyInstance = &key{
			keyType:    keyType,
			isPressed:  pressed,
			wasPressed: false,
		}
		keyboard.registeredKeys[keyType] = keyInstance
	}
	return keyInstance, nil
}

func (keyboard *keyboardImpl) UnregisterKey(keyType primitives.KeyType) {
	if _, ok := keyboard.registeredKeys[keyType]; ok {
		delete(keyboard.registeredKeys, keyType)
	}
}

func (keyboard *keyboardImpl) Update() error {
	for keyType, key := range keyboard.registeredKeys {
		pressed, err := keyboard.adapter.KeyPressed(keyType)
		if err != nil {
			return err
		}
		key.wasPressed = key.isPressed
		key.isPressed = pressed
	}
	return nil
}
