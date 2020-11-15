package input

import (
	"errors"
	"fmt"
	"github.com/GomeBox/gome/adapters/input/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyboard_RegisterKey(t *testing.T) {
	errorOnKeyPress := false
	keyPressed := false
	adapter := createKeyboardAdapterMock(&errorOnKeyPress, &keyPressed)
	keyboard := newKeyboard(adapter)
	keyType := primitives.KeyS

	errorOnKeyPress = true
	_, err := keyboard.RegisterKey(keyType)
	assert.Error(t, err, "does not return error of adapter.KeyPressed(keyType)")

	errorOnKeyPress = false
	keyPressed = true
	got, err := keyboard.RegisterKey(keyType)
	want := &key{
		keyType:    keyType,
		isPressed:  keyPressed,
		wasPressed: false,
	}
	assert.Empty(t, err, "returns unexpected error, %v", err)
	assert.Equal(t, want, got, "does not return expected key. Got %v, want %v", got, want)
	_, ok := keyboard.registeredKeys[keyType]
	assert.True(t, ok, "keyType not in registeredKeys")
	got2, err := keyboard.RegisterKey(keyType)
	assert.Empty(t, err, "returns unexpected error, %v", err)
	assert.Same(t, got, got2, "does not return same key reference on second call")
}

func TestKeyboard_UnregisterKey(t *testing.T) {
	errorOnKeyPress := false
	keyPressed := false
	adapter := createKeyboardAdapterMock(&errorOnKeyPress, &keyPressed)
	keyboard := newKeyboard(adapter)
	keyType := primitives.KeyS
	keyType2 := primitives.KeySpace
	_, _ = keyboard.RegisterKey(keyType)
	_, _ = keyboard.RegisterKey(keyType2)
	keyboard.UnregisterKey(keyType)
	_, ok := keyboard.registeredKeys[keyType]
	assert.False(t, ok, "key was not unregistered")
	_, ok = keyboard.registeredKeys[keyType2]
	assert.True(t, ok, "wrong key was unregistered")
}

func TestKeyboard_Update(t *testing.T) {
	errorOnKeyPress := false
	keyPressed := false
	adapter := createKeyboardAdapterMock(&errorOnKeyPress, &keyPressed)
	keyboard := newKeyboard(adapter)
	keyTypes := []primitives.KeyType{
		primitives.KeyS,
		primitives.KeySpace,
		primitives.KeyEsc,
		primitives.KeyRightAlt,
	}
	keys := make([]Key, len(keyTypes))
	for i, kt := range keyTypes {
		k, _ := keyboard.RegisterKey(kt)
		keys[i] = k
	}
	cases := []updateTestData{
		{wantIsPressed: false, wantWasPressed: false},
		{wantIsPressed: true, wantWasPressed: false},
		{wantIsPressed: false, wantWasPressed: true},
	}
	adapter.CallCntKeyPressed = 0
	for i, c := range cases {
		fmt.Println(adapter.CallCntKeyPressed)
		keyPressed = c.wantIsPressed
		_ = keyboard.Update()
		expectedCallCnt := len(keyTypes) * (i + 1)
		assert.Equal(t, expectedCallCnt, adapter.CallCntKeyPressed, "adapter.KeyPressed was not called once per key. Got %v, want %v", adapter.CallCntKeyPressed, expectedCallCnt)
		for _, k := range keys {
			assert.Equal(t, c.wantIsPressed, k.IsPressed(), "IsPressed has unexpected value. got %t, want %t", k.IsPressed(), c.wantIsPressed)
			assert.Equal(t, c.wantWasPressed, k.WasPressed(), "WasPressed has unexpected value. got %t, want %t", k.WasPressed(), c.wantWasPressed)
		}
	}
}

type updateTestData struct {
	wantIsPressed, wantWasPressed bool
}

func createKeyboardAdapterMock(errorOnKeyPress *bool, keyPressed *bool) *mocks.KeyboardAdapter {
	onKeyPressed := func(keyType primitives.KeyType) (bool, error) {
		if *errorOnKeyPress {
			return false, errors.New("test")
		}
		return *keyPressed, nil
	}
	return &mocks.KeyboardAdapter{
		OnKeyPressed:      onKeyPressed,
		CallCntKeyPressed: 0,
	}
}
