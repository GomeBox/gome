package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSystem(t *testing.T) {
	adapters := Adapters{}
	sys := NewSystem(adapters)
	assert.NotEmpty(t, sys, "NewSystem returned nil")
	system := sys.(*system)
	assert.NotEmpty(t, system.keyboard, "system.keyboardImpl not initialized")
}

func TestSystem_Keyboard(t *testing.T) {
	kb := new(keyboardImpl)
	system := system{keyboard: kb}
	assert.Same(t, kb, system.keyboard)
}

func TestSystem_Update(t *testing.T) {
	kb := new(keyboardStub)
	system := system{keyboard: kb}
	_ = system.Update()
	assert.Equal(t, 1, kb.CallCntUpdate)
}
