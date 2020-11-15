package input

import (
	"github.com/GomeBox/gome/adapters/input/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSystem(t *testing.T) {
	adapters := &mocks.Adapters{}
	sys := NewSystem(adapters)
	assert.NotEmpty(t, sys, "NewSystem returned nil")
	system := sys.(*system)
	assert.NotEmpty(t, system.keyboard, "system.keyboardImpl not initialized")
}

func TestSystem_Keyboard(t *testing.T) {
	kb := new(keyboardImpl)
	system := system{keyboard: kb}
	assert.Same(t, kb, system.Keyboard())
}

func TestSystem_Update(t *testing.T) {
	kb := new(KeyboardMock)
	system := system{keyboard: kb}
	_ = system.Update()
	assert.Equal(t, 1, kb.CallCntUpdate)
}
