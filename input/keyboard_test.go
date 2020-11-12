package input

import (
	"errors"
	"github.com/GomeBox/gome/internal/input/interfaces"
	"github.com/GomeBox/gome/internal/input/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyboard_UnregisterKey(t *testing.T) {
	wantKeyType := primitives.KeyEsc
	var gotKeyType primitives.KeyType
	keyboardMock := &mocks.Keyboard{
		OnUnregisterKey: func(keyType primitives.KeyType) {
			gotKeyType = keyType
		},
	}
	keyboard := keyboard{internalKeyboard: keyboardMock}
	keyboard.UnregisterKey(wantKeyType)
	assert.Equal(t, wantKeyType, gotKeyType)
}

func TestKeyboard_RegisterKey(t *testing.T) {
	wantKeyType := primitives.KeyEsc
	var gotKeyType primitives.KeyType
	retErr := false
	keyMock := new(mocks.Key)
	keyboardMock := &mocks.Keyboard{
		OnRegisterKey: func(keyType primitives.KeyType) (interfaces.Key, error) {
			if retErr {
				return nil, errors.New("test")
			}
			gotKeyType = keyType
			return keyMock, nil
		},
	}
	keyboard := keyboard{internalKeyboard: keyboardMock}
	k, err := keyboard.RegisterKey(wantKeyType)
	assert.NoError(t, err)
	assert.NotNil(t, k)
	assert.Equal(t, wantKeyType, gotKeyType)
	assert.Same(t, keyMock, (k.(*key)).internalKey)
	retErr = true
	_, err = keyboard.RegisterKey(wantKeyType)
	assert.Error(t, err)
}
