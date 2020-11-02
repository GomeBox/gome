package input

import (
	adapters "github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/internal/input/interfaces"
	"github.com/GomeBox/gome/primitives"
)

func newKeyboard(adapter adapters.KeyboardAdapter) *keyboard {
	keyboard := new(keyboard)
	keyboard.adapter = adapter
	keyboard.registeredKeys = make(map[primitives.KeyType]*key)
	return keyboard
}

type keyboard struct {
	adapter        adapters.KeyboardAdapter
	registeredKeys map[primitives.KeyType]*key
}

func (keyboard *keyboard) RegisterKey(keyType primitives.KeyType) (interfaces.Key, error) {
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
