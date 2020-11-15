package input

import (
	"github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/primitives"
)

func newKeyboard(adapter input.KeyboardAdapter) keyboard {
	keyboard := new(keyboardImpl)
	keyboard.adapter = adapter
	keyboard.registeredKeys = make(map[primitives.KeyType]*key)
	return keyboard
}

type keyboard interface {
	interfaces.Keyboard
	Update() error
}

type keyboardImpl struct {
	adapter        input.KeyboardAdapter
	registeredKeys map[primitives.KeyType]*key
}

func (keyboard *keyboardImpl) RegisterKey(keyType primitives.KeyType) (interfaces.Key, error) {
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
