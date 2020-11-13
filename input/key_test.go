package input

import (
	"github.com/GomeBox/gome/internal/input/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKey_IsPressed(t *testing.T) {
	keyMock := new(mocks.Key)
	k := key{internalKey: keyMock}
	_ = k.IsPressed()
	assert.Equal(t, 1, keyMock.CallCntIsPressed)
}

func TestKey_WasPressed(t *testing.T) {
	keyMock := new(mocks.Key)
	k := key{internalKey: keyMock}
	_ = k.WasPressed()
	assert.Equal(t, 1, keyMock.CallCntWasPressed)
}

func TestKey_KeyType(t *testing.T) {
	keyType := primitives.KeyE
	onKeyType := func() primitives.KeyType {
		return keyType
	}
	keyMock := &mocks.Key{
		OnKeyType: onKeyType,
	}
	k := key{internalKey: keyMock}
	gotKeyType := k.KeyType()
	assert.Equal(t, keyType, gotKeyType)
}
