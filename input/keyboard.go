package input

import (
	"github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/primitives"
)

type Keyboard interface {
	RegisterKey(keyType primitives.KeyType) (Key, error)
	UnregisterKey(keyType primitives.KeyType)
}

func newKeyboard(adapter input.KeyboardAdapter) *keyboard {
	keyboard := new(keyboard)
	keyboard.adapter = adapter
	keyboard.registeredKeys = make(map[primitives.KeyType]*key)
	return keyboard
}

type keyboard struct {
	adapter        input.KeyboardAdapter
	registeredKeys map[primitives.KeyType]*key
}

func (keyboard *keyboard) RegisterKey(keyType primitives.KeyType) (Key, error) {
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

func (keyboard *keyboard) UnregisterKey(keyType primitives.KeyType) {
	if _, ok := keyboard.registeredKeys[keyType]; ok {
		delete(keyboard.registeredKeys, keyType)
	}
}

func (keyboard *keyboard) Update() error {
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
